package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

var DbConn *sql.DB

func PingDatabase() {
	err := DbConn.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Successfully pinged %v:%v...\n", host, port)
}

func Initialize() {
	var err error
	fmt.Printf("Connecting to postgres server: %s:%d...\n", host, port)
	connString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	DbConn, err = sql.Open("postgres", connString)
	if err != nil {
		log.Fatal("Failed to execute query: ", err)
	}
	//defer db.Conn.Close()
	fmt.Printf("Successfully connected: %v:%v\n", host, port)
}
