package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readFile(path string) (nums []int, err error) {
	fileReader, err := os.Open(path)
	defer fileReader.Close()

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(fileReader)
	nums = make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)

		if err != nil {
			return nil, err
		}

		nums = append(nums, num)
	}
	return nums, nil
}

func solve(path string) (countAsc int, err error) {
	nums, err := readFile(path)

	if err != nil {
		return -1, err
	}

	countAsc = 0
	mapped := make([]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		mapped = append(mapped, nums[i]+nums[i+1]+nums[i+2])
	}

	for i := 1; i < len(mapped); i++ {
		if mapped[i-1] < mapped[i] {
			countAsc++
		}
	}
	return countAsc, nil
}

func main() {
	inPath := "day1/in-day1.txt"
	result, err := solve(inPath)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d\n", result)
}
