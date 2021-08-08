package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/josepht96/learning/go/serviceMonitor/cmd/api/controllers"
)

func main() {
	controllers.RegisterControllers()
	fmt.Println("Listening at: http://localhost:5000")
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
