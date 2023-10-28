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

// handlers handles the endpoint for remote scout instances
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

		// log.Printf("host: %s", r.Host)
		// log.Printf("requestURI: %s", r.RequestURI)
		// log.Printf("contentLength: %d", r.ContentLength)
		// log.Printf("method: %s", r.Method)
		// log.Printf("protocol: %s", r.Proto)
		// log.Printf("remoteAddr: %s", r.RemoteAddr)

		// log.Println("request headers:")
		// for key, values := range r.Header {
		// 	for _, value := range values {
		// 		log.Printf("\t%s: %s", key, value)
		// 	}
		// }
		// log.Println("---------------------------------------------------------")
	})
}

//getClientsetInternal creates a clientset config when inside the cluster
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

//getClientsetInternal creates a clientset config when outside the cluster
func getClientsetExternal() *kubernetes.Clientset {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"),
			"(optional) absolute path to the kubeconfig file")
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

	os.Setenv("HOSTNAME", "localhost")
	return clientset
}

//getPods returns a list of pods that match the scout label
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

//make request executes the get request to arg pod and prints output to stdout
func makeRequest(pod v1.Pod) {
	log.Printf("probing: %s -> %s.%s.%s @ node: %s", os.Getenv("HOSTNAME"), pod.Name,
		pod.Namespace,
		pod.Status.PodIP,
		pod.Spec.NodeName,
	)
	req, _ := http.NewRequest("GET", fmt.Sprintf("http://%s:%d", "localhost", port), nil)
	var r1, d1, d2, c1, c2, c3, fb time.Time
	trace := &httptrace.ClientTrace{
		DNSStart: func(_ httptrace.DNSStartInfo) {
			d1 = time.Now()
		},
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			d2 = time.Now()
			log.Printf("\tlatency dns: %s", d2.Sub(d1))
			if dnsInfo.Err != nil {
				log.Println("An error occured while handling DNS")
			}
		},
		ConnectStart: func(_, _ string) {
			if c1.IsZero() {
				// connecting to IP
				c1 = time.Now()
			}
			// log.Printf("\tconnection start: %s", c1)
		},
		ConnectDone: func(net, addr string, err error) {
			if err != nil {
				log.Printf("unable to connect to host %v: %v\n", addr, err)
			}
			c2 = time.Now()
			// log.Printf("\tconnection establish: %s", c2)
			log.Printf("\tlatency connection: %s", c2.Sub(c1))
		},
		GotConn: func(connInfo httptrace.GotConnInfo) {
			c3 = time.Now()
		},
		WroteRequest: func(_ httptrace.WroteRequestInfo) {
			r1 = time.Now()
			log.Printf("\tlatency write request: %s", r1.Sub(c3))
		},
		GotFirstResponseByte: func() {
			fb = time.Now()
			log.Printf("\tlatency server processing: %s", fb.Sub(c3))
		},
	}
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	// resp, err := http.DefaultTransport.RoundTrip(req)
	// create new transport, based on default transport
	// clone, modify, create new client using modified transport
	// use new client in get request
	newTransport := http.DefaultTransport.(*http.Transport).Clone()
	newTransport.DisableKeepAlives = true
	httpClient := &http.Client{Transport: newTransport}
	resp, err := httpClient.Do(req)

	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	endTime := time.Now()
	log.Printf("\tlatency content transfer: %s", endTime.Sub(fb))
	log.Printf("\tlatency total: %s", endTime.Sub(d1))

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	log.Printf("\tresponse: %s", string(body))
}

//probe executes the sequences of steps needed to probe an associated scout pod
func probe() {
	// time.Sleep(5 * time.Second)
	pods := getPods(getClientsetExternal())
	// pods := getPods(getClientsetInternal())
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
		time.Sleep(10 * time.Second)
	}
}

// main initializes handlers, invokes the probe mechanism, and starts a server
func main() {
	handlers()
	go probe()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		log.Printf("server is listening at http://localhost:%d", port)
		http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	}()
	wg.Wait()
}
