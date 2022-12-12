package main

import (
	"testing"
)

func TestDoit(t *testing.T) {
	result := DoIt("test_input.txt")
	if result != 10605 {
		t.Errorf("Result was %d", result)
	}
}
