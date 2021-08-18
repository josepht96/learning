package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
)

func main() {
	f, err := os.Create("./data/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	i := 0
	for i < 1000000 {
		_, err := f.WriteString(fmt.Sprintf("%d,", rand.Intn(10000)))
		if err != nil {
			fmt.Println(err)
		}
		i++
	}
}
