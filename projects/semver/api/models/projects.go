package models

import (
	"fmt"
	"log"
	"github.com/josepht96/learning/projects/semver/api/database"

	_ "github.com/lib/pq"
)
type Projects struct {
	ID int
	Name string
	Semvar string
	Description string
	Timestamp string

}
//GetProjects returns project list
func GetProjects() {
	db, err := database.Initialize()
	if err != nil {
		log.Println("error initializing database connection")
	}
	projectQuery := "SELECT * FROM PROJECTS;"
	var p Projects
	err = db.Conn.QueryRow(projectQuery).Scan(&p.ID, &p.Name, &p.Semvar, &p.Description, &p.Timestamp)
	if err != nil {
		log.Fatal("Failed to execute query: ", err)
	}
	db.Conn.Close()
	fmt.Println(p)
}