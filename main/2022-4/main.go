package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func openInput(path string) *os.File {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return f
}

type Range struct {
	min int
	max int
}

func getRange(input string) Range {
	var ints []int
	for _, s := range strings.Split(input, "-") {
		i, err := strconv.ParseInt(s, 10, 0)
		if err != nil {
			panic(err)
		}
		ints = append(ints, int(i))
	}
	return Range{
		max: ints[0],
		min: ints[1],
	}
}

func getRanges(input string) []Range {
	var ranges []Range
	for _, s := range strings.Split(input, ",") {
		ranges = append(ranges, getRange(s))
	}
	return ranges
}

func (x Range) isSuperset(other Range) bool {
	return x.min >= other.min && x.max <= other.max
}

func main() {
	f := openInput("main/2022-4/input.txt")
	defer f.Close()

	subsets := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		ranges := getRanges(line)
		if ranges[0].isSuperset(ranges[1]) || ranges[1].isSuperset(ranges[0]) {
			subsets++
		}
	}

	fmt.Println(subsets)
}
