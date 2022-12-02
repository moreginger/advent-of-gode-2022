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

func getScore(text string) (int, error) {
	switch text {
	case "A X":
		return 4, nil
	case "A Y":
		return 8, nil
	case "A Z":
		return 3, nil
	case "B X":
		return 1, nil
	case "B Y":
		return 5, nil
	case "B Z":
		return 9, nil
	case "C X":
		return 7, nil
	case "C Y":
		return 2, nil
	case "C Z":
		return 6, nil
	default:
		return -1, errors.New(fmt.Sprintf("Cannot parse %v", text))
	}
}

func main() {
	f := openInput("main/2022-2/input1.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	totalScore := 0
	for scanner.Scan() {
		text := scanner.Text()
		score, getScoreErr := getScore(text)
		if getScoreErr != nil {
			panic(getScoreErr)
		}
		totalScore += score
	}

	fmt.Println(totalScore)
}
