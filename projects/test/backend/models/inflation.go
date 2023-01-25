package models

import (
	"fmt"
	"log"

	"github.com/josepht96/learning/projects/backend/database"
)

type inflationEntry struct {
	Date      string  `json:"date"`
	CPI       float32 `json:"cpi"`
	Timestamp string  `json:"timestamp"`
}

func GetInflation() []inflationEntry {
	rows, err := database.DbConn.Query("SELECT * FROM public.inflation")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var respObject []inflationEntry
	for rows.Next() {
		var i inflationEntry
		rows.Scan(&i.Date, &i.CPI, &i.Timestamp)
		if err != nil {
			log.Fatal(err)
		}
		respObject = append(respObject, i)
		fmt.Printf("Data: %v\n", i)
	}
	return respObject
}
