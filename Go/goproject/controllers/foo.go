package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type fooHandler struct {
	Message string `json:"message1"`
}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/foo" {
		switch r.Method {
		case http.MethodGet:
			data, _ := json.Marshal(f)
			w.Write([]byte(data))
		case http.MethodPost:
			s := fooHandler{}
			bodyBytes, err := ioutil.ReadAll(r.Body)
			err = json.Unmarshal(bodyBytes, &s)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
			data, _ := json.Marshal(s)
			w.Write([]byte(data))
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		switch r.Method {
		case http.MethodGet:
			f.Message = "not /foo route"
			data, _ := json.Marshal(f)
			w.Write([]byte(data))
		case http.MethodPost:
			s := fooHandler{}
			bodyBytes, err := ioutil.ReadAll(r.Body)
			err = json.Unmarshal(bodyBytes, &s)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
			s.Message = "also not /foo route"
			data, _ := json.Marshal(s)
			w.Write([]byte(data))
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}

}
func NewFooController() *fooHandler {
	return &fooHandler{Message: "hello world"}
}
