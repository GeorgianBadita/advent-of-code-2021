package main

import (
	"bufio"
	"fmt"
	"os"
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

	pointSystem := map[string]int{")": 3, "]": 57, "}": 1197, ">": 25137}
	lookupTable := map[string]string{")": "(", "]": "[", "}": "{", ">": "<"}
	result := 0
	for _, chunk := range chunks {
		stack := make([]string, 0)
		for _, elem := range chunk {
			currentValue := string([]rune{elem})
			if currentValue == "(" || currentValue == "[" || currentValue == "{" || currentValue == "<" {
				stack = append(stack, currentValue)
			} else {
				expected := lookupTable[currentValue]
				if len(stack) == 0 {
					break
				}
				if stack[len(stack)-1] != expected {
					result += pointSystem[currentValue]
					break
				} else {
					stack = stack[:len(stack)-1]
				}
			}
		}
	}
	return result, nil
}

func main() {
	path := "day10/in-day10.txt"
	res, err := solve(path)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d\n", res)
}
