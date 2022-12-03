package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func openInput(path string) *os.File {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return f
}

func getPriority(item int32) (int32, error) {
	if 64 < item && item <= 90 {
		return item - 64 + 26, nil
	}
	if 96 < item && item <= 122 {
		return item - 96, nil
	}
	return -1, errors.New(fmt.Sprintf("Invalid item %v", item))
}

func getItems(rucksack string) []int32 {
	var items []int32
	for _, item := range rucksack {
		priority, err := getPriority(item)
		if err != nil {
			panic(err)
		}
		items = append(items, priority)
	}
	return items
}

func getCompartment(items []int32) map[int32]struct{} {
	compartment := make(map[int32]struct{})
	for _, item := range items {
		compartment[item] = struct{}{}
	}
	return compartment
}

func getCompartments(items []int32) []map[int32]struct{} {
	var compartments []map[int32]struct{}
	compartments = append(compartments, getCompartment(items[:len(items)/2]))
	compartments = append(compartments, getCompartment(items[len(items)/2:]))
	return compartments
}

func Keys[T comparable](m map[T]struct{}) []T {
	var result []T
	for k := range m {
		result = append(result, k)
	}
	return result
}

func Intersect[T comparable](l map[T]struct{}, r map[T]struct{}) map[T]struct{} {
	result := make(map[T]struct{})
	for key, _ := range l {
		if _, ok := r[key]; ok {
			result[key] = struct{}{}
		}
	}
	return result
}

func getPriorityPart1(rucksack string) int {
	items := getItems(rucksack)
	compartments := getCompartments(items)
	common := Intersect(compartments[0], compartments[1])
	priority := Keys(common)[0]
	return int(priority)
}

func getPriorityPart2(rucksacks []string) int {
	var items []map[int32]struct{}
	for _, rucksack := range rucksacks {
		items = append(items, getCompartment(getItems(rucksack)))
	}
	common := Intersect(Intersect(items[0], items[1]), items[2])
	priority := Keys(common)[0]
	return int(priority)
}

func main() {
	f := openInput("main/2022-3/input.txt")
	defer f.Close()

	sum1, sum2 := 0, 0
	var group []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		rucksack := scanner.Text()
		sum1 += getPriorityPart1(rucksack)

		group = append(group, rucksack)
		if len(group) == 3 {
			sum2 += getPriorityPart2(group)
			group = nil
		}
	}

	fmt.Println(sum1)
	fmt.Println(sum2)
}
