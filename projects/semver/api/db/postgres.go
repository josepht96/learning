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

type DatabaseConnection struct {
	connString string
	conn *sql.DB
}

func Ping(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Successfully pinged %v:%v...\n", host, port)
}

func Initialize() error {
	db := DatabaseConnection {
		connString: "",
		conn:nil,
	}
	var err error
	fmt.Printf("Connecting to postgres server at: %s:%d...\n", host, port)
	db.connString = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db.conn, err = sql.Open("postgres", db.connString)
	if err != nil {
		panic(err)
		return err
	}
	defer db.conn.Close()
	fmt.Printf("Successfully connected to: %v:%v...\n", host, port)

	

	//Ping(db)
	//queries.SelectTenants(db)
	//queries.SelectIdnTenants(db)
	//insertTenant(db)
	return nil
}
