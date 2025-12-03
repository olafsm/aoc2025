package main

import (
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	f, _ := os.Open("input_test.txt")
	defer func() {
		f.Close()
	}()
	expected := 3121910778619
	actual := run(f)
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
