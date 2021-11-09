package postgresgo

import (
	"database/sql"
	"fmt"

	//blank import
	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	hostname     = "localhost"
	hostport     = 5432
	username     = "postgres"
	password     = "pswd"
	databasename = "postgres"
)

func init() {
	var err error
	pgConn := fmt.Sprintf("port=%d host=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		hostport, hostname, username, password, databasename)

	db, err = sql.Open("postgres", pgConn)
	if err != nil {
		panic(err)
	}
	//defer db.Close()W

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Connected postgres server: port=%d host=%s user=%s "+
		"dbname=%s sslmode=disable\n",
		hostport, hostname, username, databasename)
}
func getSensors() {
	q := `
		SELECT id, name
		FROM public."sensor"
		`

	rows, err := db.Query(q)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		fmt.Printf("Id: %v, name: %v\n", id, name)
	}
}

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
