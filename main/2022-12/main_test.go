package main

import (
	"testing"
)

func TestDoit1(t *testing.T) {
	result1, result2 := DoIt("test_input.txt")
	if result1 != 31 {
		t.Errorf("Result1 was %d", result1)
	}
	if result2 != 29 {
		t.Errorf("Result2 was %d", result2)
	}
}

func TestDoit1Real(t *testing.T) {
	result1, _ := DoIt("input.txt")
	if result1 != 528 {
		t.Errorf("Result1 was %d", result1)
	}
}
