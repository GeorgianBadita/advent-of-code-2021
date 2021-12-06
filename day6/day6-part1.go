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

	fish = make([]int, len(fishSplit))
	for i, elem := range fishSplit {
		fish[i], err = strconv.Atoi(elem)
		if err != nil {
			return nil, err
		}
	}
	return fish, nil
}

func simulate(path string, iterations int) (int, error) {
	fish, err := readFile(path)
	if err != nil {
		return 0, err
	}

	for it := 0; it < iterations; it += 1 {
		newFish := make([]int, 0)
		for idx := 0; idx < len(fish); idx++ {
			currFishVal := fish[idx]
			if currFishVal == 0 {
				currFishVal = 6
				newFish = append(newFish, 8)
			} else {
				currFishVal -= 1
			}
			fish[idx] = currFishVal
		}
		for idx := 0; idx < len(newFish); idx++ {
			fish = append(fish, newFish[idx])
		}
	}
	return len(fish), nil
}

func main() {
	path := "day6/in-day6.txt"
	simDays := 80
	result, err := simulate(path, simDays)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", result)
}
