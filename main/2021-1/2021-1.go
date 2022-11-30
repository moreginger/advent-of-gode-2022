package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func panic(err error) {
	log.Fatal(err)
	os.Exit(1)
}

func getIncreases(f *os.File) int {
	increases := 0
	previous := ^uint64(0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		i, err := strconv.ParseUint(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		if i > previous {
			increases++
		}
		previous = i
	}
	return increases
}

func main() {
	input1 := "main/2021-1/input1.txt"
	f, err := os.Open(input1)
	if err != nil {
		panic(err)
	}

	fmt.Println(getIncreases(f))
}
