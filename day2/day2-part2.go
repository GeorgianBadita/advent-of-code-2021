package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile(path string) (commands []string, steps []int, err error) {
	fileReader, err := os.Open(path)
	defer fileReader.Close()

	if err != nil {
		return nil, nil, err
	}

	scanner := bufio.NewScanner(fileReader)
	steps = make([]int, 0)
	commands = make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, " ")
		step, err := strconv.Atoi(lineSplit[1])

		if err != nil {
			return nil, nil, err
		}

		steps = append(steps, step)
		commands = append(commands, lineSplit[0])
	}
	return commands, steps, nil
}

func solve(path string) (countAsc int, err error) {
	commands, steps, err := readFile(path)

	if err != nil {
		return -1, err
	}

	horizontalPos := 0
	depth := 0
	aim := 0

	for i, command := range commands {
		cmd := strings.ToLower(strings.Trim(command, " \r\n\t"))
		if strings.Compare(cmd, "forward") == 0 {
			horizontalPos += steps[i]
			depth += aim * steps[i]
		} else if strings.Compare(cmd, "down") == 0 {
			aim += steps[i]
		} else if strings.Compare(cmd, "up") == 0 {
			aim -= steps[i]
		}
	}

	return horizontalPos * depth, nil
}

func main() {
	inPath := "day2/in-day2.txt"
	result, err := solve(inPath)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d\n", result)
}
