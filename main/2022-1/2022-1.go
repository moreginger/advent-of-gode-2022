package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func panic(err error) {
	log.Fatal(err)
	os.Exit(1)
}

func openInput(path string) *os.File {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return f
}

func getTotalCalories(f *os.File) []uint64 {
	var result []uint64
	scanner := bufio.NewScanner(f)
	var total uint64 = 0
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			result = append(result, total)
			total = 0
		} else {
			i, err := strconv.ParseUint(text, 10, 64)
			if err != nil {
				panic(err)
			}
			total += i
		}

	}
	return result
}

func main() {
	f := openInput("main/2022-1/input1.txt")
	totals := getTotalCalories(f)
	sort.Slice(totals, func(i, j int) bool {
		return totals[i] > totals[j]
	})
	println(totals[0])

	var top3Total uint64 = 0
	for _, cals := range totals[0:3] {
		top3Total += cals
	}
	println(top3Total)
}
