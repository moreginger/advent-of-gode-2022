package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func openInput(path string) *os.File {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return f
}

func ParseCrates(input string) ([]string, bool) {
	var cratePattern = regexp.MustCompile("(    |\\[([A-Z])\\] ?)")
	var crates []string
	isCrates := false
	for _, crate := range cratePattern.FindAllStringSubmatch(input, 10) {
		if crate[2] == "" {
			crates = append(crates, "")
		} else {
			isCrates = true
			crates = append(crates, crate[2])
		}
	}
	return crates, isCrates
}

func padStacks(stacks [][]string, index int) [][]string {
	for i := len(stacks); i <= index; i++ {
		stacks = append(stacks, []string{})
	}
	return stacks
}

func parseStacks(scanner *bufio.Scanner) [][]string {
	var stacks [][]string

	for scanner.Scan() {
		line := scanner.Text()
		crates, isCrates := ParseCrates(line)
		if !isCrates {
			break
		}
		for i, crate := range crates {
			stacks = padStacks(stacks, i)
			if crate != "" {
				stacks[i] = append(stacks[i], crate)
			}
		}
	}

	return stacks
}

type Instruction struct {
	move int
	from int
	to   int
}

func parseInt(input string) int {
	i, err := strconv.ParseInt(input, 10, 0)
	if err != nil {
		panic(err)
	}
	return int(i)
}

func parseInstructions(scanner *bufio.Scanner) []Instruction {
	var instructionPattern = regexp.MustCompile("move (\\d+) from (\\d+) to (\\d+)")
	var instructions []Instruction
	for scanner.Scan() {
		line := scanner.Text()
		instructionMatch := instructionPattern.FindStringSubmatch(line)
		instructions = append(instructions, Instruction{move: parseInt(instructionMatch[1]), from: parseInt(instructionMatch[2]) - 1, to: parseInt(instructionMatch[3]) - 1})
	}
	return instructions
}

func reverse(stack []string) []string {
	var result []string
	for i := len(stack) - 1; i >= 0; i-- {
		result = append(result, stack[i])
	}
	return result
}

func applyInstructions(stacks [][]string, instructions []Instruction, blockMove bool) [][]string {
	result := make([][]string, len(stacks))
	copy(result, stacks)
	for _, instruction := range instructions {
		toMove := make([]string, instruction.move)
		copy(toMove, result[instruction.from][:instruction.move])
		if !blockMove {
			toMove = reverse(toMove)
		}
		result[instruction.to] = append(toMove, result[instruction.to]...)
		result[instruction.from] = result[instruction.from][instruction.move:]
	}
	return result
}

func main() {
	f := openInput("main/2022-5/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	stacks := parseStacks(scanner)
	scanner.Scan() // Throw blank line away.
	instructions := parseInstructions(scanner)
	part1 := applyInstructions(stacks, instructions, false)
	part2 := applyInstructions(stacks, instructions, true)

	for _, stack := range part1 {
		fmt.Print(stack[0])
	}
	fmt.Println()
	for _, stack := range part2 {
		fmt.Print(stack[0])
	}
}
