package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	isSmall bool
	name    string
}

type CaveSystem struct {
	graph map[string][]Node
}

func insertNode(nodeName string, toAdd string, isSmall bool, caveSystem CaveSystem) {
	neighbours, exists := caveSystem.graph[nodeName]
	if !exists {
		caveSystem.graph[nodeName] = make([]Node, 0)
		neighbours = caveSystem.graph[nodeName]
	}
	neighbours = append(neighbours, Node{isSmall, toAdd})
	caveSystem.graph[nodeName] = neighbours
}

func printCaveSystem(caveSystem CaveSystem) {
	for nodeName, neighbours := range caveSystem.graph {
		fmt.Printf("%s: %v\n", nodeName, neighbours)
	}
}

func readFile(path string) (caveSystem CaveSystem, err error) {
	fd, err := os.Open(path)
	if err != nil {
		return CaveSystem{}, err
	}

	scanner := bufio.NewScanner(fd)
	caveSystem = CaveSystem{}
	caveSystem.graph = make(map[string][]Node)

	for scanner.Scan() {
		line := strings.Split(strings.Trim(scanner.Text(), " "), "-")
		leftNode := line[0]
		rightNode := line[1]

		isRightSmall := "a" <= rightNode && rightNode <= "z"
		isLeftSmall := "a" <= leftNode && leftNode <= "z"

		insertNode(leftNode, rightNode, isRightSmall, caveSystem)
		insertNode(rightNode, leftNode, isLeftSmall, caveSystem)
	}

	return caveSystem, nil
}

func main() {
	path := "day12/in-day12.txt"

	cave, err := readFile(path)
	if err != nil {
		panic(err)
	}

	printCaveSystem(cave)
}
