package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readFile(path string) (digits []string, err error) {
	fd, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(fd)
	digits = make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		end := strings.Trim(strings.Split(line, "|")[1], " ")

		for _, digit := range strings.Split(end, " ") {
			digits = append(digits, digit)
		}
	}
	return digits, nil
}

func solve(path string) (int, error) {
	digits, err := readFile(path)
	if err != nil {
		return 0, err
	}

	result := 0
	for _, digit := range digits {
		if len(digit) == 7 || len(digit) == 2 || len(digit) == 3 || len(digit) == 4 {
			result += 1
		}
	}
	return result, nil
}

func main() {
	path := "day8/in-day8.txt"
	res, err := solve(path)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d\n", res)
}
