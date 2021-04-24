package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"

	"github.com/josepht96/learning/go/goproject/foundation/models"
)

type fooHandler struct {
	idPattern *regexp.Regexp
}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/foo" {
		switch r.Method {
		case http.MethodGet:
			f.getFoo(w, r)
		case http.MethodPost:
			f.postFoo(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		switch r.Method {
		case http.MethodGet:
			f.getAltFoo(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}

}
func (f *fooHandler) getFoo(w http.ResponseWriter, r *http.Request) {
	fi, err := models.GetFoo()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data, _ := json.Marshal(fi)
	w.Write([]byte(data))
}

func (f *fooHandler) getAltFoo(w http.ResponseWriter, r *http.Request) {
	//Get ULR id
	id := f.getURLParam(w, r)
	fi, err := models.GetAltFoo(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data, _ := json.Marshal(fi)
	w.Write([]byte(data))
}
func (f *fooHandler) postFoo(w http.ResponseWriter, r *http.Request) {
	fi, err := f.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data, _ := json.Marshal(fi)
	w.Write([]byte(data))
}

func (f *fooHandler) parseRequest(r *http.Request) (models.Foo, error) {
	//declare instance
	s := models.Foo{}
	bodyBytes, err := ioutil.ReadAll(r.Body)
	//cast from json to type Foo
	err = json.Unmarshal(bodyBytes, &s)
	if err != nil {
		return models.Foo{}, err
	}
	return s, nil
}
func (f *fooHandler) getURLParam(w http.ResponseWriter, r *http.Request) int {
	matches := f.idPattern.FindStringSubmatch(r.URL.Path)
	fmt.Printf("URL path is: %v\n", r.URL.Path)
	if len(matches) == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	id, _ := strconv.Atoi(matches[1])
	return id
}
func newFooController() *fooHandler {
	return &fooHandler{
		idPattern: regexp.MustCompile(`^/foo/(\d+)/?`),
	}
}
