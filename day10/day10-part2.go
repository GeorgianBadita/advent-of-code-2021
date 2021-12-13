package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func readFile(path string) (chunks []string, err error) {
	fd, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(fd)

	chunks = make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		chunks = append(chunks, strings.Trim(line, " "))
	}
	return chunks, nil
}

func solve(path string) (int, error) {
	chunks, err := readFile(path)
	if err != nil {
		return 0, err
	}

	pointSystem := map[string]int{"(": 1, "[": 2, "{": 3, "<": 4}
	lookupTable := map[string]string{")": "(", "]": "[", "}": "{", ">": "<"}

	goodChunks := make([]string, 0)

	for _, chunk := range chunks {
		stack := make([]string, 0)
		isGoodChunk := true
		for _, elem := range chunk {
			currentValue := string([]rune{elem})
			if currentValue == "(" || currentValue == "[" || currentValue == "{" || currentValue == "<" {
				stack = append(stack, currentValue)
			} else {
				expected := lookupTable[currentValue]
				if stack[len(stack)-1] != expected {
					isGoodChunk = false
					break
				} else {
					stack = stack[:len(stack)-1]
				}
			}
		}
		if isGoodChunk {
			goodChunks = append(goodChunks, chunk)
		}
	}
	scores := make([]int, 0)

	for _, chunk := range goodChunks {
		stack := make([]string, 0)
		idx := 0
		for idx < len(chunk) {
			moved := false
			currVal := string([]byte{chunk[idx]})
			if currVal == "(" || currVal == "[" || currVal == "{" || currVal == "<" {
				stack = append(stack, currVal)
			} else {
				expected := lookupTable[currVal]
				for len(stack) > 0 && idx < len(chunk) && stack[len(stack)-1] == expected {
					stack = stack[:len(stack)-1]
					idx += 1
					moved = true
					if idx == len(chunk) {
						break
					}
					currVal = string([]byte{chunk[idx]})
					expected = lookupTable[currVal]
				}
			}
			if !moved {
				idx++
			}

		}
		score := 0
		for i := len(stack) - 1; i >= 0; i-- {
			score = score*5 + pointSystem[stack[i]]
		}
		scores = append(scores, score)
	}
	sort.Ints(scores)
	return scores[len(scores)/2], nil
}

func main() {
	path := "day10/in-day10.txt"
	res, err := solve(path)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d\n", res)
}
