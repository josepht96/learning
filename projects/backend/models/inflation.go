package models

import (
	"fmt"

	"github.com/josepht96/learning/projects/backend/database"
)

func GetInflation() {
	fmt.Println("inflation")
	database.PingDatabase()
}
