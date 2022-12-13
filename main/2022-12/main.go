package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"time"
)

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

type Node struct {
	x         int
	y         int
	height    int
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
				height:    int(lines[y][x]),
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

func getDistances(nodes []*Node, startNode *Node) map[*Node]int {
	distance := make(map[*Node]int, len(nodes))
	for _, node := range nodes {
		distance[node] = math.MaxInt
	}
	distance[startNode] = 0

	previous := make(map[*Node]*Node, len(nodes))
	for _, node := range nodes {
		previous[node] = nil
	}

	queue := make([]*Node, len(nodes))
	for i, node := range nodes {
		queue[i] = node
	}
	for len(queue) > 0 {
		qiMin, distanceMin := 0, math.MaxInt
		var node *Node
		for qi, n := range queue {
			if distance[n] < distanceMin {
				qiMin = qi
				distanceMin = distance[n]
				node = n
			}
		}
		queue = append(queue[:qiMin], queue[qiMin+1:]...)

		if node != startNode && previous[node] == nil {
			// Unreachable
			break
		}

		d := distanceMin + 1

		for _, v := range node.neighbors {
			inQueue := false
			for _, qi := range queue {
				if v == qi {
					inQueue = true
					break
				}
			}
			if inQueue && (node.height-v.height) <= 1 {
				if d < distance[v] {
					distance[v] = d
					previous[v] = node
				}
			}
		}
	}

	return distance
}

func DoIt(inputName string) (int, int) {
	nodes, startNode, endNode := ReadNodes(inputName)
	distances := getDistances(nodes, endNode)

	fmt.Printf("Nodes: %d\n", len(nodes))
	result1 := distances[startNode]
	result2 := math.MaxInt
	for node, distance := range distances {
		if node.height == 'a' {
			if result2 > distance {
				result2 = distance
			}
		}
	}

	return result1, result2

}

func main() {
	start := time.Now()
	result1, result2 := DoIt("main/2022-12/input.txt")
	elapsed := time.Since(start)
	log.Printf("Elapsed %s", elapsed)
	fmt.Println(result1)
	fmt.Println(result2)
}
