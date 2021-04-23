package controllers

import "fmt"

func Hello(name string) (string, error) {
	message := fmt.Sprintf("Hello, %v!", name)
	return message, nil
}
