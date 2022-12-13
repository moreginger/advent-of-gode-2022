package main

import (
	"testing"
)

func TestDoit1(t *testing.T) {
	result := DoIt("test_input.txt")
	if result != 31 {
		t.Errorf("Result was %d", result)
	}
}
