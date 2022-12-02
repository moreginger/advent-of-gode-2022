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

func getScore1(text string) (int, error) {
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

func getScore2(text string) (int, error) {
	switch text {
	case "A X":
		return 3, nil
	case "A Y":
		return 4, nil
	case "A Z":
		return 8, nil
	case "B X":
		return 1, nil
	case "B Y":
		return 5, nil
	case "B Z":
		return 9, nil
	case "C X":
		return 2, nil
	case "C Y":
		return 6, nil
	case "C Z":
		return 7, nil
	default:
		return -1, errors.New(fmt.Sprintf("Cannot parse %v", text))
	}
}

func main() {
	f := openInput("main/2022-2/input1.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	totalScore1, totalScore2 := 0, 0
	for scanner.Scan() {
		text := scanner.Text()
		score, getScoreErr := getScore1(text)
		if getScoreErr != nil {
			panic(getScoreErr)
		}
		totalScore1 += score

		score, getScoreErr = getScore2(text)
		if getScoreErr != nil {
			panic(getScoreErr)
		}
		totalScore2 += score
	}

	fmt.Println(totalScore1)
	fmt.Println(totalScore2)
}
