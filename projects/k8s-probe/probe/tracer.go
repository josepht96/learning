package main

import (
	"log"
	"net/http"
	"net/http/httptrace"
	"time"
)

type Tracer struct {
	url                                         string
	ClientTrace                                 *httptrace.ClientTrace
	r1, d1, d2, c0, c1, c2, c3, fb              time.Time
	dnsDur, connDur, totalDur, serverprocessDur time.Duration
}

func CreateTransportObj() *http.Transport {
	// resp, err := http.DefaultTransport.RoundTrip(req)
	// create new transport, based on default transport
	// clone, modify, create new client using modified transport
	// use new client in get request
	newTransport := http.DefaultTransport.(*http.Transport).Clone()
	newTransport.DisableKeepAlives = true
	return newTransport
}

func CreateTraceObj(traceURL string) *Tracer {
	t := &Tracer{}
	t.url = traceURL
	t.ClientTrace = &httptrace.ClientTrace{
		GetConn: func(hostPort string) {
			t.c0 = time.Now()
			log.Printf("connecting to: %s", t.url)
			log.Printf("\tconnection start: %s", t.c0)
		},
		DNSStart: func(_ httptrace.DNSStartInfo) {
			t.d1 = time.Now()
		},
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			t.d2 = time.Now()
			t.dnsDur = t.d2.Sub(t.d1)
			log.Printf("\tlatency dns: %s", t.dnsDur)
			if dnsInfo.Err != nil {
				log.Println("An error occured while resolving DNS")
			}
			log.Printf("\t\tdnsInfo: %v", dnsInfo.Addrs)
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
			t.connDur = t.c2.Sub(t.c1)
			log.Printf("\tlatency connection established: %s", t.connDur)
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
			t.serverprocessDur = t.fb.Sub(t.c3)
			log.Printf("\tlatency server processing: %s", t.serverprocessDur)
		},
	}
	return t
}
