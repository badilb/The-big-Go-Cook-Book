package main

import (
	"testing"
)

func TestParseStringToInt(t *testing.T) {
	num, err := ParseStringToInt("123456789")
	if err != nil {
		t.Fatal(err)
	}
	if num != 123456789 {
		t.Fatal("Number don't match")
	}
}

func TestFailParseStringToInt(t *testing.T) {
	_, err := ParseStringToInt("")
	if err == nil {
		t.Fatal(err)
	}
}
