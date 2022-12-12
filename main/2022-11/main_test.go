package main

import (
	"testing"
)

func TestDoit1(t *testing.T) {
	result := DoIt("test_input.txt", 20, 3)
	if result != 10605 {
		t.Errorf("Result was %d", result)
	}
}

func TestDoit2(t *testing.T) {
	result := DoIt("test_input.txt", 10000, 1)
	if result != 2713310158 {
		t.Errorf("Result was %d", result)
	}
}
