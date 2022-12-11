package main

import (
	"bufio"
	"errors"
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

type Coordinate struct {
	x int
	y int
}

type Movement struct {
	x int
	y int
}

func (m Movement) apply(coordinate Coordinate) Coordinate {
	return Coordinate{
		x: coordinate.x + m.x,
		y: coordinate.y + m.y,
	}
}

func generateMovements(x int, y int, steps int) []Movement {
	var movements []Movement
	for i := 0; i < steps; i++ {
		movements = append(movements, Movement{
			x: x,
			y: y,
		})
	}
	return movements
}

func readMovements(line string) []Movement {
	movementRegex := regexp.MustCompile("^([UDLR]) ([0-9]+)$")
	movementMatch := movementRegex.FindStringSubmatch(line)
	direction := movementMatch[1]
	steps := parseInt(movementMatch[2])

	if direction == "U" || direction == "D" {
		y := 1
		if direction == "D" {
			y = -1
		}
		return generateMovements(0, y, steps)
	} else if direction == "L" || direction == "R" {
		x := 1
		if direction == "L" {
			x = -1
		}
		return generateMovements(x, 0, steps)
	} else {
		panic(errors.New(direction))
	}
}

func sign(num int) int {
	if num > 0 {
		return 1
	} else if num < 0 {
		return -1
	} else {
		return 0
	}
}

func adjustTail(head Coordinate, tail Coordinate) Coordinate {
	hx, hy := head.x, head.y
	tx, ty := tail.x, tail.y
	moving := hx-tx > 1 || hx-tx < -1 || hy-ty > 1 || hy-ty < -1
	if !moving {
		return tail
	}
	if hx != tx {
		tx = tx + sign(hx-tx)
	}
	if hy != ty {
		ty = ty + sign(hy-ty)
	}
	return Coordinate{
		x: tx,
		y: ty,
	}
}

func Doit(inputName string, knots int) int {
	f := openInput(inputName)
	scanner := bufio.NewScanner(f)

	rope := make([]Coordinate, knots)
	for i := range rope {
		rope[i] = Coordinate{
			x: 0, y: 0,
		}
	}

	visited := make(map[Coordinate]struct{})
	visited[rope[knots-1]] = struct{}{}

	for scanner.Scan() {
		line := scanner.Text()
		movements := readMovements(line)
		for _, movement := range movements {
			rope[0] = movement.apply(rope[0])
			for i := 1; i < knots; i++ {
				rope[i] = adjustTail(rope[i-1], rope[i])
			}
			visited[rope[knots-1]] = struct{}{}
		}
	}

	return len(visited)
}

func main() {
	fmt.Println(Doit("main/2022-9/input.txt", 1))
	fmt.Println(Doit("main/2022-9/input.txt", 10))
}
