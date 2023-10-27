package main

import (
	"context"
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

	"github.com/fatih/color"
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

const (
	httpTemplate = `` +
		"\t" + `printing table...` + "\n" +
		`   DNS Lookup   TCP Connection   Server Processing   Content Transfer` + "\n" +
		`[ %-9s  |     %-9s  |        %-9s  |       %-9s  ]` + "\n" +
		`             |                |                   |                  |` + "\n" +
		`    namelookup:%-9s      |                   |                  |` + "\n" +
		`                        connect:%-9s         |                  |` + "\n" +
		`                                      starttransfer:%-9s        |` + "\n" +
		`                                                                 total:%-9s` + "\n"
)

func printf(format string, a ...interface{}) (n int, err error) {
	fmt.Println("In printf")
	return fmt.Fprintf(color.Output, format, a...)
}

func grayscale(code color.Attribute) func(string, ...interface{}) string {
	return color.New(code + 232).SprintfFunc()
}

//getClientsetInternal creates a clientset config when inside in the cluster
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

//getClientsetInternal creates a clientset config when outside in the cluster
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
	req, _ := http.NewRequest("GET", fmt.Sprintf("http://%s:8080", "localhost"), nil)
	var d1, d2, c1, c2, fb time.Time
	trace := &httptrace.ClientTrace{
		DNSStart: func(_ httptrace.DNSStartInfo) {
			d1 = time.Now()
		},
		DNSDone: func(_ httptrace.DNSDoneInfo) {
			d2 = time.Now()
			log.Printf("\tlatency dns: %s", d2.Sub(d1))
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
		GotFirstResponseByte: func() {
			fb = time.Now()
			log.Printf("\tlatency firstByte: %s", fb.Sub(c2))
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
	}
	defer resp.Body.Close()
	now := time.Now()

	log.Printf("\tlatency total: %s", now.Sub(d1))

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	log.Printf("\tresponse: %s", string(body))

	log.Printf(httpTemplate, d2.Sub(d1), c2.Sub(c1), fb.Sub(c2), now.Sub(fb), d2.Sub(d1), c2.Sub(d1), fb.Sub(d1), now.Sub(d1))
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
	var wg sync.WaitGroup
	wg.Add(1)
	go probe()
	wg.Wait()
}
