package controllers

import (
	//"encoding/json"
	//"fmt"
	//"io/ioutil"
	"net/http"
	"regexp"
	//"strconv"
	"github.com/josepht96/learning/projects/semver/api/models"
)

type projectHandler struct {
	guidPattern *regexp.Regexp
}

func (p *projectHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/project/" {
		switch r.Method {
		case http.MethodGet:
			getProject(w, r)
		case http.MethodPut:
			updateProject(w, r)
		case http.MethodPost:
			createProject(w, r)
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

func getProject(w http.ResponseWriter, r *http.Request) {
	models.GetProject()
	w.WriteHeader(http.StatusOK)
}

func updateProject(w http.ResponseWriter, r *http.Request) {
	models.UpdateProject()
	w.WriteHeader(http.StatusOK)
}
func createProject(w http.ResponseWriter, r *http.Request) {
	models.CreateProject()
	w.WriteHeader(http.StatusOK)
}

func newProjectHandler() *projectHandler {
	return &projectHandler{
		guidPattern: regexp.MustCompile(`^/project/(\d+)/?`),
	}
}

// func (f *fooHandler) getFoo(w http.ResponseWriter, r *http.Request) {
// 	fi, err := models.GetFoo()
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	data, _ := json.Marshal(fi)
// 	w.Write([]byte(data))
// }

// func (f *fooHandler) getAltFoo(w http.ResponseWriter, r *http.Request) {
// 	//Get ULR id
// 	id := f.getURLParam(w, r)
// 	fi, err := models.GetAltFoo(id)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	data, _ := json.Marshal(fi)
// 	w.Write([]byte(data))
// }
// func (f *fooHandler) postFoo(w http.ResponseWriter, r *http.Request) {
// 	fi, err := f.parseRequest(r)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	fi, err = models.PostFoo(fi)
// 	data, _ := json.Marshal(fi)
// 	w.Write([]byte(data))
// }

// func (f *fooHandler) parseRequest(r *http.Request) (models.Foo, error) {
// 	//declare instance
// 	s := models.Foo{}
// 	bodyBytes, err := ioutil.ReadAll(r.Body)
// 	//cast from json to type Foo
// 	err = json.Unmarshal(bodyBytes, &s)
// 	if err != nil {
// 		return models.Foo{}, err
// 	}
// 	return s, nil
// }
// func (f *fooHandler) getURLParam(w http.ResponseWriter, r *http.Request) int {
// 	matches := f.idPattern.FindStringSubmatch(r.URL.Path)
// 	fmt.Printf("URL path is: %v\n", r.URL.Path)
// 	if len(matches) == 0 {
// 		w.WriteHeader(http.StatusNotFound)
// 	}
// 	id, _ := strconv.Atoi(matches[1])
// 	return id
// }
