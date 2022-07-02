package controllers

import (
	"fmt"
	"net/http"
	"github.com/josepht96/learning/projects/backend/models"
)

type inflationHandler struct {
}

func (i *inflationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/inflation" {
		switch r.Method {
		case http.MethodGet:
			getInflationData(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		switch r.Method {
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}

}

func getInflationData(w http.ResponseWriter, r *http.Request) {
	models.GetInflation()
	w.Write([]byte(rMsg))
}

func inflationResponse(w http.ResponseWriter, r *http.Request) {
	rMsg := fmt.Sprintf("Accessing inflation controller\n")
	w.Write([]byte(rMsg))
}

func newInflationHandler() *inflationHandler {
	return &inflationHandler{
	}
}