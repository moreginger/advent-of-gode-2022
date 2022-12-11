package main

import (
	"os"
	"testing"
)

func TestDoit(t *testing.T) {
	inputFile := t.TempDir() + "/input"
	f, err := os.Create(inputFile)
	if err != nil {
		panic(err)
	}
	f.WriteString(`R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2
`)

	result := Doit(inputFile)
	if result != 13 {
		t.Errorf("Result was %d", result)
	}
}
