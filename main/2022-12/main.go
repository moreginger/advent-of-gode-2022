package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

type Node struct {
	x         int
	y         int
	height    uint8
	neighbors []*Node
}

func ReadNodes(inputName string) ([]*Node, *Node, *Node) {
	bytes, err := os.ReadFile(inputName)
	panicOnErr(err)

	lines := strings.Split(string(bytes), "\n")
	lines = lines[:len(lines)-1]
	nodes := make([][]*Node, len(lines))
	for i := range nodes {
		nodes[i] = make([]*Node, len(lines[0]))
	}

	var startNode, endNode *Node
	for y := 0; y < len(nodes); y++ {
		for x := 0; x < len(nodes[y]); x++ {
			node := &Node{
				x:         x,
				y:         y,
				height:    lines[y][x],
				neighbors: []*Node{},
			}
			if node.height == 'S' {
				node.height = 'a'
				startNode = node
			} else if node.height == 'E' {
				node.height = 'z'
				endNode = node
			}

			nodes[y][x] = node
		}
	}

	for y := 0; y < len(nodes); y++ {
		for x := 0; x < len(nodes[y]); x++ {
			node := nodes[y][x]
			if x-1 >= 0 {
				node.neighbors = append(node.neighbors, nodes[y][x-1])
			}
			if x+1 < len(nodes[y]) {
				node.neighbors = append(node.neighbors, nodes[y][x+1])
			}
			if y-1 >= 0 {
				node.neighbors = append(node.neighbors, nodes[y-1][x])
			}
			if y+1 < len(nodes) {
				node.neighbors = append(node.neighbors, nodes[y+1][x])
			}
		}
	}

	var flatNodes []*Node
	for _, row := range nodes {
		flatNodes = append(flatNodes, row...)
	}

	// DEBUG
	//for i, n := range flatNodes {
	//	fmt.Print(*n, ", ")
	//	if i%8 == 7 {
	//		fmt.Println()
	//	}
	//}

	return flatNodes, startNode, endNode
}

func getIndex(nodes []*Node, query *Node) int {
	for i, node := range nodes {
		if node == query {
			return i
		}
	}
	panic(errors.New("not found"))
}

func DoIt(inputName string) int {
	nodes, startNode, endNode := ReadNodes(inputName)

	//fmt.Println("startNode", startNode)
	//fmt.Println("endNode", endNode)

	startNodeIndex := getIndex(nodes, startNode)
	endNodeIndex := getIndex(nodes, endNode)
	distance := make([]int, len(nodes))
	for i := range nodes {
		distance[i] = math.MaxInt
	}
	distance[startNodeIndex] = 0

	previous := make([]int, len(nodes))
	for i := range nodes {
		previous[i] = -1
	}

	queue := make([]int, len(nodes))
	for i, _ := range nodes {
		queue[i] = i
	}
	for len(queue) > 0 {
		sort.Slice(queue, func(i, j int) bool {
			return distance[queue[i]] < distance[queue[j]]
		})
		nodeIndex := queue[0]
		queue = queue[1:]

		if nodeIndex != startNodeIndex && previous[nodeIndex] == -1 {
			// Unreachable
			continue
		}

		for _, v := range nodes[nodeIndex].neighbors {
			vertexIndex := getIndex(nodes, v)
			inQueue := false
			for _, qi := range queue {
				if vertexIndex == qi {
					inQueue = true
					break
				}
			}
			if inQueue && (nodes[vertexIndex].height-nodes[nodeIndex].height) <= 1 {
				d := distance[nodeIndex] + 1
				if d < distance[vertexIndex] {
					distance[vertexIndex] = d
					previous[vertexIndex] = nodeIndex
				}
			}
		}
	}

	// DEBUG
	//for i, v := range distance {
	//	char := "."
	//	if v == math.MaxInt {
	//		char = "X"
	//	}
	//	fmt.Print(char)
	//	if i%181 == 180 {
	//		fmt.Println()
	//	}
	//}

	return distance[endNodeIndex]
}

func main() {
	result := DoIt("main/2022-12/input.txt")
	fmt.Println(result)
}
