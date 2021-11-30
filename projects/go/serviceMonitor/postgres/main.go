package main

import (
	"database/sql"
	"fmt"

	"github.com/josepht96/learning/go/serviceMonitor/postgres/queries"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func ping(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully pinged server...")
}

func main() {
	fmt.Printf("Connecting to postgres server at %s:%d...\n", host, port)
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlConn)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("Successfully connected...")

	ping(db)
	queries.SelectTenants(db)
	queries.SelectIdnTenants(db)
	//insertTenant(db)
}
