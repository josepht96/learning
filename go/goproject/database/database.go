package database

import "fmt"

//DBConnection is db connection
var DBConnection *sql.DB

//SetupDatabase setups database connection
func SetupDatabase() error {
	fmt.Println("Hello from database package")
	return nil
}