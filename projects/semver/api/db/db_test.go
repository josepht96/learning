package db

import (
	"testing"
)

func TestInitialize(t *testing.T) {
	err := Initialize()
	if err != nil {
		t.Errorf("Error initializing database: %v", err)
	  }
}