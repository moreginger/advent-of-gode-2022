package main

import (
	"bufio"
	"fmt"
	"os"
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

func readTrees() [][]int {
	var trees [][]int

	f := openInput("input.txt")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		treeLine := make([]int, len(line))
		for i, t := range line {
			treeLine[i] = int(t)
		}
		trees = append(trees, treeLine)
	}

	return trees
}

type Coordinate struct {
	x int
	y int
}

func coordinate(x int, y int) Coordinate {
	return Coordinate{
		x: x,
		y: y,
	}
}

func main() {
	trees := readTrees()
	visibleTrees := make(map[Coordinate]struct{})

	rows := len(trees)
	columns := len(trees[0])
	max := 0

	for row := 0; row < rows; row++ {
		max = 0
		for col := 0; col < columns; col++ {
			if trees[row][col] > max {
				max = trees[row][col]
				visibleTrees[coordinate(row, col)] = struct{}{}
			}
		}
		max = 0
		for col := columns - 1; col >= 0; col-- {
			if trees[row][col] > max {
				max = trees[row][col]
				visibleTrees[coordinate(row, col)] = struct{}{}
			}
		}
	}

	for col := 0; col < columns; col++ {
		max = 0
		for row := 0; row < rows; row++ {
			if trees[row][col] > max {
				max = trees[row][col]
				visibleTrees[coordinate(row, col)] = struct{}{}
			}
		}
		max = 0
		for row := rows - 1; row >= 0; row-- {
			if trees[row][col] > max {
				max = trees[row][col]
				visibleTrees[coordinate(row, col)] = struct{}{}
			}
		}
	}

	fmt.Println(rows * columns)
	fmt.Println(len(visibleTrees))
}
