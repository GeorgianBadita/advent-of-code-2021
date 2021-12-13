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

func countPaths(caveSystem CaveSystem, sol []string, used map[string]int, countSol *int) {
	if sol[len(sol)-1] == "end" {
		*countSol++
		return
	}
	for _, neighbour := range caveSystem.graph[sol[len(sol)-1]] {
		if !neighbour.isSmall {
			sol = append(sol, neighbour.name)
			countPaths(caveSystem, sol, used, countSol)
			sol = sol[:len(sol)-1]
		} else {
			isAlready2 := false
			for key, occ := range used {
				if key != "start" && occ == 2 {
					isAlready2 = true
					break
				}
			}
			currCount, _ := used[neighbour.name]
			if (isAlready2 && currCount >= 1) || neighbour.name == "start" {
				continue
			}

			used[neighbour.name] += 1
			sol = append(sol, neighbour.name)

			countPaths(caveSystem, sol, used, countSol)
			sol = sol[:len(sol)-1]
			used[neighbour.name] -= 1
		}
	}
}

func main() {
	path := "day12/in-day12.txt"

	cave, err := readFile(path)
	if err != nil {
		panic(err)
	}
	result := 0
	used := make(map[string]int)
	sol := []string{"start"}
	used["start"] = 2
	for nodeName := range cave.graph {
		if nodeName != "start" && nodeName != "end" && "a" <= nodeName && nodeName <= "z" {
			used[nodeName] = 0
		}
	}

	countPaths(cave, sol, used, &result)

	fmt.Printf("%d\n", result)
}
