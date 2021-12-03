package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFile(path string) (nums []string, err error) {
	fileReader, err := os.Open(path)
	defer fileReader.Close()

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(fileReader)

	nums = make([]string, 0)

	for scanner.Scan() {
		nums = append(nums, scanner.Text())
	}
	return nums, nil
}

func solve(path string) (result int, err error) {
	nums, err := readFile(path)

	if err != nil {
		return 0, err
	}

	gamma, epsilon := 0, 0
	for i := len(nums[0]) - 1; i >= 0; i-- {
		countZero := 0
		countOne := 0
		for idx := 0; idx < len(nums); idx++ {
			num := nums[idx]
			if num[i] == '0' {
				countZero += 1
			} else {
				countOne += 1
			}

		}
		// i = 4 - one: 5, zero: 7, gamma: 00000, epsilon: 00001
		// i = 3 - one: 7, zero: 5, gamma: 00010, epsolon: 00001
		// i = 2 - one: 8, zero: 4, gamma: 00110, epsilon: 00001
		// i = 1 - one: 5, zero: 7, gamma: 00110, epsilon: 01001
		// i = 0 - one: 7, zero: 5, gamma: 10110, epsilon: 01001
		if countOne > countZero {
			gamma |= 1 << (len(nums[0]) - 1 - i)
		} else if countZero > countOne {
			epsilon |= 1 << (len(nums[0]) - 1 - i)
		}
	}

	return gamma * epsilon, nil
}

func main() {
	path := "day3/in-day3.txt"
	result, err := solve(path)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%d", result)
}
