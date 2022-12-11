package main

import (
	"os"
	"testing"
)

func writeTestFile(t *testing.T, name string, content string) string {
	inputFile := t.TempDir() + "/" + name
	f, err := os.Create(inputFile)
	if err != nil {
		panic(err)
	}
	f.WriteString(content)
	return inputFile
}

func TestDoit2(t *testing.T) {
	inputPath := writeTestFile(t, "input", `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2
`)

	result := Doit(inputPath, 2)
	if result != 13 {
		t.Errorf("Result was %d", result)
	}
}

func TestDoit10(t *testing.T) {
	inputPath := writeTestFile(t, "input", `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20
`)

	result := Doit(inputPath, 10)
	if result != 13 {
		t.Errorf("Result was %d", result)
	}
}
