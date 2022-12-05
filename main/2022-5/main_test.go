package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseCrates(t *testing.T) {
	result, _ := ParseCrates("    [F] [T] [B] [D]     [P]     [P]")
	if !reflect.DeepEqual(result, []string{"", "F", "T", "B", "D", "", "P", "", "P"}) {
		t.Errorf("Failed")
	}
}

func TestApplyInstructions(t *testing.T) {
	stacks := [][]string{{"A", "B"}, {"C"}}
	instructions := []Instruction{{move: 2, from: 0, to: 1}, {move: 2, from: 1, to: 0}}
	result := applyInstructions(stacks, instructions, true)

	fmt.Println(result)
}
