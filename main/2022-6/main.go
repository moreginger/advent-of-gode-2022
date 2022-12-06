package main

import (
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

func main() {

	windowSize := 4
	window := make([]rune, windowSize)
	for i := range window {
		window[i] = -1
	}
	windowIndex := 0

	counts := make(map[rune]int, 52)
	multiples := 0

	f := openInput("main/2022-6/input.txt")
	b := make([]byte, 16)
	fileIndex := 0

	done := false

	for {
		n, err := f.Read(b)
		if n == 0 {
			break
		}
		panicOnErr(err)

		for i := 0; i < n; i++ {
			ro := window[windowIndex]
			rn := rune(b[i])
			window[windowIndex] = rn
			windowIndex = (windowIndex + 1) % windowSize

			counts[ro] = counts[ro] - 1
			if counts[ro] == 1 {
				multiples--
			}
			counts[rn] = counts[rn] + 1
			if counts[rn] == 2 {
				multiples++
			}

			if multiples == 0 && fileIndex+1 >= windowSize {
				done = true
				break
			}

			fileIndex++
		}

		if done {
			break
		}
	}

	fmt.Println(fileIndex + 1)
}
