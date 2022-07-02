package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/josepht96/learning/projects/backend/controllers"
	"github.com/josepht96/learning/projects/backend/database"
)

func main() {
	controllers.RegisterControllers()
	database.Initialize()
	fmt.Println("Listening at: http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
