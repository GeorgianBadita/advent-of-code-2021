package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile(path string) (nums []int, err error) {
	fd, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(fd)
	scanner.Scan()
	numStr := scanner.Text()

	numsSplit := strings.Split(numStr, ",")
	nums = make([]int, len(numsSplit))

	for idx, num := range numsSplit {
		numInt, err := strconv.Atoi(num)

		if err != nil {
			return nil, err
		}
		nums[idx] = numInt
	}
	return nums, nil
}

func max(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	res := nums[0]
	for idx := 1; idx < len(nums); idx++ {
		if nums[idx] > res {
			res = nums[idx]
		}
	}
	return res
}

func computeRes(nums []int, target int) int {
	res := 0
	for _, num := range nums {
		val := target - num
		if target < num {
			val = num - target
		}
		res += (val * (val + 1)) / 2
	}
	return res
}

func solve(path string) (int, error) {
	nums, err := readFile(path)

	if err != nil {
		return 0, err
	}

	maxNum := max(nums)
	bestRes := computeRes(nums, 0)
	for target := 1; target <= maxNum; target += 1 {
		currRes := computeRes(nums, target)
		if currRes <= bestRes {
			bestRes = currRes
		}
	}
	return bestRes, err
}

func main() {
	path := "day7/in-day7.txt"
	res, err := solve(path)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d\n", res)
}
