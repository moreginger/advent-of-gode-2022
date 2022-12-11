package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

type Registers struct {
	X int
}

type Instruction interface {
	step(r *Registers) bool
}

type NoopInstruction struct {
	worked int
}

func (i *NoopInstruction) step(_ *Registers) bool {
	i.worked++
	return i.worked == 1
}

type AddxInstruction struct {
	value  int
	worked int
}

func (i *AddxInstruction) step(r *Registers) bool {
	i.worked++
	if i.worked == 2 {
		r.X += i.value
		return true
	}
	return false
}

func parseInstruction(line string) Instruction {
	if line == "noop" {
		return &NoopInstruction{}
	}
	addxRegex := regexp.MustCompile("^addx ([0-9-]+)$")
	addxMatch := addxRegex.FindStringSubmatch(line)
	value := parseInt(addxMatch[1])
	return &AddxInstruction{
		value:  value,
		worked: 0,
	}
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func DoIt(inputName string) int {
	f := openInput(inputName)
	scanner := bufio.NewScanner(f)

	registers := &Registers{
		X: 1,
	}

	cycle := 1
	signalSum := 0

	for scanner.Scan() {
		line := scanner.Text()
		instruction := parseInstruction(line)
		for complete := false; complete == false; complete = instruction.step(registers) {
			if (cycle-20)%40 == 0 {
				signalSum += registers.X * cycle
			}

			position := (cycle - 1) % 40
			char := '.'
			if abs(registers.X-position) <= 1 {
				char = '#'
			}
			fmt.Printf("%c", char)
			if position == 39 {
				fmt.Println()
			}
			cycle++
			//fmt.Printf("%s %+v %+v\n", reflect.TypeOf(instruction), instruction, registers)
		}

	}

	return signalSum
}

func main() {
	result := DoIt("main/2022-10/input.txt")
	fmt.Println(result)
}
