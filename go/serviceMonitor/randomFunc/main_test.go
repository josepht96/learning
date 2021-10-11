package main

import "testing"

func TestPrintString(t *testing.T) {
	s := PrintString("test")
	if s != "test" {
		t.Errorf(s)
	}
}
