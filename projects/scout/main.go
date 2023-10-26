package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptrace"
	"os"
	"path/filepath"
	"sync"
	"time"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type Response struct {
	Status     string  `json:"status"`
	StatusCode int     `json:"statusCode"`
	Body       Message `json:"body"`
}

type Message struct {
	Message string `json:"message"`
	Time    string `json:"time"`
}

var port = 8080

// handlers handles the / endpoint for remote scout instances
func handlers() {
	// http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		dt := time.Now()
		data := Response{
			Status:     "OK",
			StatusCode: 200,
			Body: Message{
				Message: "connected to scout",
				Time:    dt.String(),
			},
		}
		json.NewEncoder(w).Encode(data)
		log.Printf("host: %s", r.Host)
		log.Printf("requestURI: %s", r.RequestURI)
		log.Printf("contentLength: %d", r.ContentLength)
		log.Printf("method: %s", r.Method)
		log.Printf("protocol: %s", r.Proto)
		log.Printf("remoteAddr: %s", r.RemoteAddr)

		log.Println("request headers:")
		for key, values := range r.Header {
			for _, value := range values {
				log.Printf("\t%s: %s", key, value)
			}
		}
		log.Println("---------------------------------------------------------")

	})
}

func getClientsetInternal() *kubernetes.Clientset {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clientset
}

func getClientsetExternal() *kubernetes.Clientset {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clientset
}

func getPods(clientset *kubernetes.Clientset) *v1.PodList {
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{
		LabelSelector: "app=scout",
	})
	if err != nil {
		fmt.Printf("Error fetching pods: %s\n", err)
		os.Exit(1)
	}
	return pods
}

func makeRequest(pod v1.Pod) {
	log.Printf("probing: %s.%s.%s @ node: %s", pod.Name,
		pod.Namespace,
		pod.Status.PodIP,
		pod.Spec.NodeName,
	)
	req, _ := http.NewRequest("GET", "http://localhost:8081", nil)
	var c1, c2 time.Time
	trace := &httptrace.ClientTrace{
		ConnectStart: func(_, _ string) {
			if c1.IsZero() {
				// connecting to IP
				c1 = time.Now()
			}
			log.Printf("\tConnectStart: %s", c1)
		},
		ConnectDone: func(net, addr string, err error) {
			if err != nil {
				log.Fatalf("unable to connect to host %v: %v", addr, err)
			}
			c2 = time.Now()
			log.Printf("\tConnectDone: %s", c2)
			log.Printf("\tlatency: %s", c2.Sub(c1))
		},
		// GotConn: func(connInfo httptrace.GotConnInfo) {
		// 	log.Printf("\tconnection reused: %v\n", connInfo.Reused)
		// 	log.Printf("\tconnection idle: %v\n", connInfo.WasIdle)
		// },
	}
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	// takes a request, returns response
	// resp, err := http.DefaultTransport.RoundTrip(req)
	// create new transport, based on default transport
	// clone, modify, create new client using modified transport
	// use new client in get request
	newTransport := http.DefaultTransport.(*http.Transport).Clone()
	newTransport.DisableKeepAlives = true
	httpClient := &http.Client{Transport: newTransport}
	resp, err := httpClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("\tresponse: %s", string(body))
}
func probe() {
	pods := getPods(getClientsetExternal())
	log.Println("scout pods:")
	for _, pod := range pods.Items {
		log.Printf("\tnode: %s", pod.Spec.NodeName)
		log.Printf("\t\t%s.%s.%s", pod.Name,
			pod.Namespace,
			pod.Status.PodIP,
		)
	}
	for true {
		for _, pod := range pods.Items {
			makeRequest(pod)
		}
		time.Sleep(5 * time.Second)
	}
}
func main() {
	go handlers()
	go probe()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		log.Printf("server is listening at http://localhost:%d", port)
		http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	}()
	wg.Wait()
}
