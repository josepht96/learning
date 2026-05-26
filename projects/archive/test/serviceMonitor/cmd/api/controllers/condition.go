package controllers

import (
	"net/http"
	"regexp"
	"strconv"

	"github.com/josepht96/learning/go/serviceMonitor/cmd/api/models/condition"
)

type conditionC struct {
	patternID *regexp.Regexp
}

func (api *conditionC) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/condition" {
		switch r.Method {
		case http.MethodGet:
			api.getCondtion(w, r)
		default:
			api.defaultReponse(w)
		}
	} else {
		switch r.Method {
		case http.MethodGet:
			api.getConditions(w, r)
		default:
			api.defaultReponse(w)
		}
	}
}
func (api *conditionC) getCondtion(w http.ResponseWriter, r *http.Request) {
	response, err := condition.Get()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write([]byte(response))
}

func (api *conditionC) getConditions(w http.ResponseWriter, r *http.Request) {
	id := api.getURLParam(w, r.URL.Path)
	if id == -1 {
		return
	}
	response, err := condition.GetID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write([]byte(response))
}

func (api *conditionC) defaultReponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (api *conditionC) getURLParam(w http.ResponseWriter, url string) int {
	matches := api.patternID.FindStringSubmatch(url)
	if len(matches) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return -1
	} else {
		id, err := strconv.Atoi(matches[1])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return id
	}

}
func newConditionController() *conditionC {
	return &conditionC{
		patternID: regexp.MustCompile(`^/condition/(\d+$)/?`),
	}
}
