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
		min: ints[0],
		max: ints[1],
	}
}

func getRanges(input string) []Range {
	var ranges []Range
	for _, s := range strings.Split(input, ",") {
		ranges = append(ranges, getRange(s))
	}
	return ranges
}

func (x Range) IsSuperset(other Range) bool {
	return x.min >= other.min && x.max <= other.max
}

func (x Range) Overlaps(other Range) bool {
	return x.IsSuperset(other) ||
		other.IsSuperset(x) ||
		(x.min >= other.min && x.min <= other.max) ||
		(x.max >= other.min && x.max <= other.max)
}

func main() {
	f := openInput("main/2022-4/input.txt")
	defer f.Close()

	subsets := 0
	overlaps := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		ranges := getRanges(line)
		if ranges[0].IsSuperset(ranges[1]) || ranges[1].IsSuperset(ranges[0]) {
			subsets++
		}
		if ranges[0].Overlaps(ranges[1]) {
			overlaps++
		}
	}

	fmt.Println(subsets)
	fmt.Println(overlaps)
}
