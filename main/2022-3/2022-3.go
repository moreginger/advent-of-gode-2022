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

func getCompartment(items []int32) map[int32]struct{} {
	compartment := make(map[int32]struct{})
	for _, item := range items {
		compartment[item] = struct{}{}
	}
	return compartment
}

func getCompartments(rucksack string) []map[int32]struct{} {
	var items []int32
	for _, item := range rucksack {
		priority, err := getPriority(item)
		if err != nil {
			panic(err)
		}
		items = append(items, priority)
	}

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

func main() {
	f := openInput("main/2022-3/input.txt")

	sum := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		rucksack := scanner.Text()
		compartments := getCompartments(rucksack)
		common := Intersect(compartments[0], compartments[1])
		priority := Keys(common)[0]
		sum += int(priority)
	}

	fmt.Println(sum)
}
