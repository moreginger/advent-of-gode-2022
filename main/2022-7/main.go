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

type Content struct {
	name string
	size int
}

func main() {
	f := openInput("input.txt")
	scanner := bufio.NewScanner(f)
	scanner.Scan() // Ignore initial "$ cd /"

	commandRe := regexp.MustCompile("^\\$ ((ls)|(cd \\.\\.)|(cd [a-z]+))$")
	contentsRe := regexp.MustCompile("^((dir [a-z]+)|(([0-9]+) ([a-z.]+)))$")

	root := Content{
		name: "",
		size: 0,
	}
	path := []Content{root}
	dirs := map[string]Content{"": root}

	for scanner.Scan() {
		line := scanner.Text()
		commandMatch := commandRe.FindStringSubmatch(line)
		contentsMatch := contentsRe.FindStringSubmatch(line)
		if len(commandMatch) > 0 {
			if len(commandMatch[2]) > 0 {
				// ls
				// FIXME: cope with revisiting for ls.
			} else if len(commandMatch[3]) > 0 {
				// cd ..
				path = path[:len(path)-1]
			} else if len(commandMatch[4]) > 0 {
				// cd abcd
				name := commandMatch[4][3:]
				pathString := ""
				for _, dir := range path {
					pathString += dir.name + "/"
				}
				pathString += name

				fmt.Println(pathString)
				dir, ok := dirs[pathString]
				if !ok {
					dir = Content{
						name: name,
						size: 0,
					}
					dirs[pathString] = dir
				}
				path = append(path, dir)
			}
		} else if len(contentsMatch) > 0 {
			if len(contentsMatch[2]) > 0 {
				// dir
			} else if len(contentsMatch[3]) > 0 {
				// file
				size := parseInt(contentsMatch[4])

				pathString := ""
				for i, dir := range path {
					pathString += dir.name
					path[i].size += size
					dirs[pathString] = path[i]
					pathString += "/"
				}

			}
		}
	}

	total := 0
	for _, dir := range dirs {
		fmt.Println(dir)
		if dir.size <= 100000 {
			total += dir.size
		}
	}

	fmt.Println(total)
}
