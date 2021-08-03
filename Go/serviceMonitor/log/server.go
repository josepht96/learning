package log

import (
	"fmt"
	"io/ioutil"
	stlog "log"
	"net/http"
	"os"
)

var log *stlog.Logger

type fileLog string

func (fl fileLog) Write(data []byte) (int, error) {
	fmt.Println("in Write")
	f, err := os.OpenFile(string(fl), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return f.Write(data)
}

//Run func
func Run(destination string) {
	fmt.Println("in Run")
	log = stlog.New(fileLog(destination), "", stlog.LstdFlags)
}

//RegisterHandlers func
func RegisterHandlers() {
	fmt.Println("in RegisterHandlers")
	http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		msg, err := ioutil.ReadAll(r.Body)
		if err != nil || len(msg) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		write(string(msg))
	})
}

func write(message string) {
	fmt.Println("in write")
	fmt.Printf("%v\n", message)
	log.Printf("%v\n", message)
}
