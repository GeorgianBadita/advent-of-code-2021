package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func findRatingRec(nums []string, currentBit int, ratingType string) string {
	if len(nums) == 1 {
		return nums[0]
	}

	cntZero, cntOne := 0, 0
	for idx := 0; idx < len(nums); idx++ {
		bit := nums[idx][currentBit]

		if bit == '0' {
			cntZero += 1
		} else {
			cntOne += 1
		}
	}

	newNums := make([]string, 0)
	compareBy := '0'

	if strings.Compare(ratingType, "O2") == 0 {
		if cntZero == cntOne {
			compareBy = '1'
		} else if cntZero > cntOne {
			compareBy = '0'
		} else {
			compareBy = '1'
		}
	}

	if strings.Compare(ratingType, "CO2") == 0 {
		if cntZero == cntOne {
			compareBy = '0'
		} else if cntZero > cntOne {
			compareBy = '1'
		} else {
			compareBy = '0'
		}
	}

	for _, num := range nums {
		if num[currentBit] == byte(compareBy) {
			newNums = append(newNums, num)
		}
	}

	return findRatingRec(newNums, currentBit+1, ratingType)
}

func convert(num string) int {
	res := 0
	currPow := 0
	for idx := len(num) - 1; idx >= 0; idx-- {
		if num[idx] == '1' {
			res |= 1 << currPow
		}
		currPow += 1
	}
	return res
}

func findO2Co2Ratings(nums []string) (int, int) {
	o2Rating := findRatingRec(nums, 0, "O2")
	co2Rating := findRatingRec(nums, 0, "CO2")
	return convert(o2Rating), convert(co2Rating)
}

func solve(path string) (result int, err error) {
	nums, err := readFile(path)

	if err != nil {
		return 0, err
	}

	o2Rating, co2Rating := findO2Co2Ratings(nums)

	return o2Rating * co2Rating, nil
}

func main() {
	path := "day3/in-day3.txt"
	result, err := solve(path)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%d", result)
}
