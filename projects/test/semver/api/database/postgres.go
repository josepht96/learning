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

type DatabaseConnection struct {
	ConnString string
	Conn *sql.DB
}

func Ping(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Successfully pinged %v:%v...\n", host, port)
}

func Initialize() (DatabaseConnection, error) {
	db := DatabaseConnection {
		ConnString: "",
		Conn:nil,
	}
	var err error
	fmt.Printf("Connecting to postgres server at: %s:%d...\n", host, port)
	db.ConnString = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db.Conn, err = sql.Open("postgres", db.ConnString)
	if err != nil {
		log.Fatal("Failed to execute query: ", err)
	}
	//defer db.Conn.Close()
	fmt.Printf("Successfully connected to: %v:%v...\n", host, port)

	return db, nil
}
