package queries

import (
	"database/sql"
	"fmt"
	"log"
)

func insertTenant(db *sql.DB) {
	t := Tenant{
		ID:   "4",
		Name: "Tenant 4",
	}
	queryStr := "INSERT INTO public.tenants(id, name) VALUES ($1, $2);"
	_, err := db.Exec(queryStr, t.ID, t.Name)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Printf("Successfully inserted object with ID: %s, Name: %s...\n", t.ID, t.Name)
	}
}
