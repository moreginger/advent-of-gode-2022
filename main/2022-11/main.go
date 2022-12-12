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

func parseInt(input string) int {
	i, err := strconv.ParseInt(input, 10, 0)
	if err != nil {
		panic(err)
	}
	return int(i)
}

type Monkey struct {
	items  []int
	apply  func(item int) int
	test   func(item int) int
	tested int
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

	var items []int
	for _, item := range strings.Split(startingItemsRaw, ", ") {
		items = append(items, parseInt(item))
	}

	var apply func(item int) int
	if strings.HasPrefix(operationRaw, "+ ") {
		value := parseInt(operationRaw[2:])
		apply = func(item int) int {
			return item + value
		}
	} else if strings.HasPrefix(operationRaw, "* ") {
		if operationRaw == "* old" {
			apply = func(item int) int {
				return item * item
			}
		} else {
			value := parseInt(operationRaw[2:])
			apply = func(item int) int {
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
		test: func(item int) int {
			if item%test == 0 {
				return testTrue
			}
			return testFalse
		},
		tested: 0,
	}
}

func DoIt(inputName string) int {
	f := openInput(inputName)
	scanner := bufio.NewScanner(f)

	var monkeys []*Monkey

	for scanner.Scan() {
		monkeys = append(monkeys, parseMonkey(scanner))
		if !scanner.Scan() {
			break
		}
	}

	for i := 0; i < 20; i++ {
		for _, m := range monkeys {
			for _, item := range m.items {
				item = m.apply(item)
				item /= 3
				destination := monkeys[m.test(item)]
				destination.items = append(destination.items, item)
				m.tested++
			}
			m.items = make([]int, 0)
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].tested > monkeys[j].tested
	})

	return monkeys[0].tested * monkeys[1].tested
}

func main() {
	result := DoIt("main/2022-11/input.txt")
	fmt.Println(result)
}
