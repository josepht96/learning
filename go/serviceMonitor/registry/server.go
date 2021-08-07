package registry

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

//ServerPort const
const ServerPort = ":3000"

//ServicesURL const
const ServicesURL = "http://localhost" + ServerPort + "/services"

//ExposeURL const
const ExposeURL = "http://172.17.0.1:8081/services"

type registry struct {
	registrations []Registration
	mutex         *sync.Mutex
}

func (r *registry) add(reg Registration) error {
	r.mutex.Lock()
	r.registrations = append(r.registrations, reg)
	r.mutex.Unlock()
	return nil
}

func (r *registry) remove(url string) error {
	for i := range r.registrations {
		if r.registrations[i].ServiceURL == url {
			r.mutex.Lock()
			r.registrations = append(r.registrations[:i], r.registrations[i+1:]...)
			r.mutex.Unlock()
			return nil
		}
	}
	return fmt.Errorf("service at URL %v not found", url)
}

var r = registry{registrations: make([]Registration, 0),
	mutex: new(sync.Mutex),
}

//RegistryService struct
type RegistryService struct{}

func (s RegistryService) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Println("Request received")
	switch req.Method {

	case http.MethodGet:
		fmt.Println("Incoming Get request...")
		_, err := w.Write([]byte("Request service"))
		if err != nil {
			log.Println(err)
		}

	case http.MethodPost:
		dec := json.NewDecoder(req.Body)
		var reg Registration
		err := dec.Decode(&reg)
		fmt.Printf("Incoming request values: %v\n", reg)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Printf("Adding service: %v with URL: %v\n", reg.ServiceName,
			reg.ServiceURL)
		err = r.add(reg)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

	case http.MethodDelete:
		payload, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		url := string(payload)
		log.Printf("Removing service at URL: %v", url)
		err = r.remove(url)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
