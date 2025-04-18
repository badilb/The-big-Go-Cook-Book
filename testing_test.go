package main

import (
	"cookBook/handling_errors"
	"testing"
)

func TestParseStringToInt(t *testing.T) {
	num, err := handling_errors.ParseStringToInt("123456789")
	if err != nil {
		t.Fatal(err)
	}
	if num != 123456789 {
		t.Fatal("Number don't match")
	}
}

func TestFailParseStringToInt(t *testing.T) {
	_, err := handling_errors.ParseStringToInt("")
	if err == nil {
		t.Fatal(err)
	}
}
