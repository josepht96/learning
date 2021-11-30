package models

import (
	//"fmt"
	"log"
	"github.com/josepht96/learning/projects/semver/api/db"
)

//GetProjects returns project list
func GetProjects() {
	err := db.Initialize()
	if err != nil {
		log.Println("error initializing database connection")
	}
}