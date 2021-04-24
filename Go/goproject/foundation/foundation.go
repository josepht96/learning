package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/josepht96/learning/go/goproject/controllers"
)

func main() {
	controllers.RegisterControllers()
	fmt.Println("Listening at: localhost:5000")
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
