package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func Ping(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully pinged server...")
}

func Initialize() error {
	fmt.Printf("Connecting to postgres server at %s:%d...\n", host, port)
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlConn)
	if err != nil {
		panic(err)
		return err
	}
	defer db.Close()
	fmt.Println("Successfully connected...")

	return nil

	//ping(db)
	//queries.SelectTenants(db)
	//queries.SelectIdnTenants(db)
	//insertTenant(db)
}
