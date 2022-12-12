package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func openInput(path string) *os.File {
	f, err := os.Open(path)
	panicOnErr(err)
	return f
}

func parseInt(input string) int64 {
	i, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

type Monkey struct {
	items  []int64
	apply  func(item int64) int64
	test   func(item int64) int64
	tested int64
}

func parseMonkey(scanner *bufio.Scanner) *Monkey {
	// Ignore header text
	scanner.Scan()
	startingItemsRaw := scanner.Text()[len("  Starting items: "):]
	scanner.Scan()
	operationRaw := scanner.Text()[len("  Operation: new = old "):]
	scanner.Scan()
	testRaw := scanner.Text()[len("  Test: divisible by "):]
	scanner.Scan()
	testTrueRaw := scanner.Text()[len("    If true: throw to monkey "):]
	scanner.Scan()
	testFalseRaw := scanner.Text()[len("    If false: throw to monkey "):]

	var items []int64
	for _, item := range strings.Split(startingItemsRaw, ", ") {
		items = append(items, parseInt(item))
	}

	var apply func(item int64) int64
	if strings.HasPrefix(operationRaw, "+ ") {
		value := parseInt(operationRaw[2:])
		apply = func(item int64) int64 {
			return item + value
		}
	} else if strings.HasPrefix(operationRaw, "* ") {
		if operationRaw == "* old" {
			apply = func(item int64) int64 {
				return item * item
			}
		} else {
			value := parseInt(operationRaw[2:])
			apply = func(item int64) int64 {
				return item * value
			}
		}

	} else {
		panic(errors.New(operationRaw))
	}

	test := parseInt(testRaw)
	testTrue := parseInt(testTrueRaw)
	testFalse := parseInt(testFalseRaw)

	return &Monkey{
		items: items,
		apply: apply,
		test: func(item int64) int64 {
			if item%test == 0 {
				return testTrue
			}
			return testFalse
		},
		tested: 0,
	}
}

func DoIt(inputName string, loops int, worryReduction int64) int64 {
	f := openInput(inputName)
	scanner := bufio.NewScanner(f)

	var monkeys []*Monkey

	for scanner.Scan() {
		monkeys = append(monkeys, parseMonkey(scanner))
		if !scanner.Scan() {
			break
		}
	}

	for i := 0; i < loops; i++ {
		for _, m := range monkeys {
			for _, item := range m.items {
				item = m.apply(item)
				item /= worryReduction
				destination := monkeys[m.test(item)]
				destination.items = append(destination.items, item)
				m.tested++
			}
			m.items = make([]int64, 0)
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].tested > monkeys[j].tested
	})

	fmt.Println(monkeys[0].tested, monkeys[1].tested)
	return monkeys[0].tested * monkeys[1].tested
}

func main() {
	result := DoIt("main/2022-11/input.txt", 20, 3)
	fmt.Println(result)
	result = DoIt("main/2022-11/input.txt", 10000, 1)
	fmt.Println(result)
}
