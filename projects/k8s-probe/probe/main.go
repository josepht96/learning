package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptrace"
	"sync"
	"text/template"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var promTotalReq = createMetricTotalRequests()
var promTotalReqLatency = createMetricTotalLatency()
var promTotalDNSDur = createMetricDNSDur()
var promTotalConnDur = createMetricConnDur()
var promTotalServerProcessingDur = createMetricServerProcessingDur()
var promCertNotBefore = createMetricCertNotBefore()
var promCertNotAfter = createMetricCertNotAfter()

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	data := struct {
		Name string
	}{
		Name: "Gopher",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func probe(config Config) {
	for key, ep := range config.Endpoints {
		log.Printf("Starting probe for %s", key)
		URL := fmt.Sprintf("%s://%s", ep.Protocol, ep.URL)
		labels := prometheus.Labels{
			"endpoint": URL,
		}
		req, _ := http.NewRequest("GET", URL, nil)
		tracer := CreateTraceObj(URL)
		req = req.WithContext(httptrace.WithClientTrace(req.Context(), tracer.ClientTrace))
		httpClient := &http.Client{Transport: CreateTransportObj()}
		resp, err := httpClient.Do(req)
		if err != nil {
			panic(err)
		}
		if resp.TLS != nil {
			certs := resp.TLS.PeerCertificates
			for _, cert := range certs {
				labels := prometheus.Labels{
					"endpoint": URL,
					"issuer":   cert.Issuer.String(),
				}
				log.Printf("Issuer Name: %s\n", cert.Issuer.String())
				log.Printf("Expiry: %s \n", cert.NotBefore.Format("2006-January-02"))
				promCertNotBefore.With(labels).Set(float64(cert.NotBefore.Unix()))
				log.Printf("Expiry: %s \n", cert.NotAfter.Format("2006-January-02"))
				promCertNotAfter.With(labels).Set(float64(cert.NotAfter.Unix()))
				log.Printf("Common Name: %s \n", cert.Issuer.CommonName)
			}
		}
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()

		promTotalReq.With(labels).Inc()
		// these are done in nanonseconds, number will get large very quickly
		// rounding will cause problems
		// really short durations will potentially never leave 0 as theyll be rounded down
		promTotalReqLatency.With(labels).Add(float64(tracer.totalDur.Milliseconds()))
		promTotalDNSDur.With(labels).Add(float64(tracer.dnsDur.Microseconds()))
		promTotalConnDur.With(labels).Add(float64(tracer.connDur.Microseconds()))
		promTotalServerProcessingDur.With(labels).Add(float64(tracer.serverprocessDur.Microseconds()))
	}
}

// handlers handles the endpoint for remote scout instances
func handlers() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", serveTemplate)
}

func main() {
	config := parseConfig()
	handlers()
	log.Println("Serving on http://localhost:8080...")
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			fmt.Println("Server error:", err)
		}
	}()

	probe(*config)

	wg.Wait()
}
