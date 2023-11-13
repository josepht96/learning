package main

import (
	"context"
	"encoding/json"
	"errors"
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

	"github.com/prometheus/client_golang/prometheus/promhttp"
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
}

type Tracer struct {
	ClientTrace                    *httptrace.ClientTrace
	r1, d1, d2, c0, c1, c2, c3, fb time.Time
}

var port = 8080

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data := Response{
		Status:     "OK",
		StatusCode: 200,
		Body: Message{
			Message: "connected to scout",
		},
	}
	json.NewEncoder(w).Encode(data)
}

// handlers handles the endpoint for remote scout instances
func handlers() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", rootHandler)
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
		log.Fatalf("Error fetching pods: %s\n", err)
	}
	return pods
}

func createRequestObj(insideCluster bool, pod v1.Pod) *http.Request {
	var req *http.Request
	if insideCluster {
		req, _ = http.NewRequest("GET", fmt.Sprintf("http://%s:%d", pod.Status.PodIP, port), nil)
	} else {
		req, _ = http.NewRequest("GET", fmt.Sprintf("http://%s:%d", "localhost", port), nil)
	}
	return req
}

func createTransportObj() *http.Transport {
	// resp, err := http.DefaultTransport.RoundTrip(req)
	// create new transport, based on default transport
	// clone, modify, create new client using modified transport
	// use new client in get request
	newTransport := http.DefaultTransport.(*http.Transport).Clone()
	newTransport.DisableKeepAlives = true
	return newTransport
}

func createTraceObj(insideCluster bool) *Tracer {
	t := &Tracer{}
	// var r1, d1, d2, c0, c1, c2, c3, fb *time.Time
	t.ClientTrace = &httptrace.ClientTrace{
		GetConn: func(hostPort string) {
			t.c0 = time.Now()
			log.Printf("\tconnection start: %s", t.c0)
		},
		DNSStart: func(_ httptrace.DNSStartInfo) {
			t.d1 = time.Now()
		},
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			if insideCluster {
				return
			}
			t.d2 = time.Now()
			log.Printf("\tlatency dns: %s", t.d2.Sub(t.d1))
			if dnsInfo.Err != nil {
				log.Println("An error occured while handling DNS")
			}
		},
		ConnectStart: func(_, _ string) {
			if t.c1.IsZero() {
				// connecting to IP
				t.c1 = time.Now()
			}
			// log.Printf("\tconnection start: %s", c1)
		},
		ConnectDone: func(net, addr string, err error) {
			if err != nil {
				log.Printf("unable to connect to host %v: %v\n", addr, err)
			}
			t.c2 = time.Now()
			// log.Printf("\tconnection establish: %s", c2)
			log.Printf("\tlatency connection: %s", t.c2.Sub(t.c1))
		},
		GotConn: func(connInfo httptrace.GotConnInfo) {
			t.c3 = time.Now()
		},
		WroteRequest: func(_ httptrace.WroteRequestInfo) {
			t.r1 = time.Now()
			log.Printf("\tlatency write request: %s", t.r1.Sub(t.c3))
		},
		GotFirstResponseByte: func() {
			t.fb = time.Now()
			log.Printf("\tlatency server processing: %s", t.fb.Sub(t.c3))
		},
	}
	return t
}

//make request executes the get request to arg pod and prints output to stdout
func makeRequest(insideCluster bool, pod v1.Pod) error {
	var req *http.Request
	// var tracer *Tracer
	if pod.Status.PodIP == "" {
		return errors.New(fmt.Sprintf("IP address for %s could not be determined", pod.Name))
	}
	log.Printf("probing: %s -> %s.%s.%s @ node: %s", os.Getenv("HOSTNAME"),
		pod.Name,
		pod.Namespace,
		pod.Status.PodIP,
		pod.Spec.NodeName,
	)
	req = createRequestObj(insideCluster, pod)
	tracer := createTraceObj(insideCluster)

	req = req.WithContext(httptrace.WithClientTrace(req.Context(), tracer.ClientTrace))
	httpClient := &http.Client{Transport: createTransportObj()}
	resp, err := httpClient.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()
	endTime := time.Now()
	log.Printf("\tlatency content transfer: %s", endTime.Sub(tracer.fb))
	log.Printf("\tlatency total: %s", endTime.Sub(tracer.c0))

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Printf("\tresponse: %s", string(body))

	return nil
}

//probe executes the sequences of steps needed to probe an associated scout pod
func probe() {
	insideCluster := true
	var clientset *kubernetes.Clientset

	if os.Getenv("HOSTNAME") != "" {
		clientset = getClientsetInternal()
	} else {
		insideCluster = false
		clientset = getClientsetExternal()
	}

	for {
		pods := getPods(clientset)
		log.Println("scout pods:")
		for _, pod := range pods.Items {
			log.Printf("\tnode: %s", pod.Spec.NodeName)
			log.Printf("\t\t%s.%s.%s", pod.Name,
				pod.Namespace,
				pod.Status.PodIP,
			)
		}
		for _, pod := range pods.Items {
			err := makeRequest(insideCluster, pod)
			if err != nil {
				log.Println(err)
			}

		}
		time.Sleep(15 * time.Second)
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
		err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
		if err != nil {
			log.Fatal(err)
		}
	}()
	wg.Wait()
}
