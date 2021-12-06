package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile(path string) (fish []int, err error) {
	fd, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(fd)
	scanner.Scan()

	line := scanner.Text()
	fishSplit := strings.Split(line, ",")

	fish = make([]int, 9)

	for _, elem := range fishSplit {
		age, err := strconv.Atoi(elem)
		if err != nil {
			return nil, err
		}
		fish[age]++
	}
	return fish, nil
}

func simulate(path string, iterations int) (int, error) {
	fish, err := readFile(path)

	if err != nil {
		return 0, err
	}

	for it := 0; it < iterations; it += 1 {
		justMade := fish[0]
		for d := range fish[:len(fish)-1] {
			fish[d] = fish[d+1]
		}
		fish[6] += justMade
		fish[8] = justMade
	}
	fmt.Printf("%v\n", fish)
	numFish := 0
	for _, num := range fish {
		numFish += num
	}
	return numFish, nil
}

func main() {
	path := "day6/in-day6.txt"
	simDays := 256
	result, err := simulate(path, simDays)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", result)
}
