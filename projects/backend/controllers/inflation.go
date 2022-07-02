package controllers

import (
	"encoding/json"
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
	respObject := models.GetInflation()
	respJson, _ := json.Marshal(respObject)
	w.Write([]byte(respJson))
}

func newInflationHandler() *inflationHandler {
	return &inflationHandler{}
}
