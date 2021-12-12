package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFile(path string) (heights [][]int, err error) {
	fd, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(fd)

	heights = make([][]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		height := make([]int, 0)

		for _, elem := range line {
			height = append(height, int(elem-'0'))
		}
		heights = append(heights, height)
	}
	return heights, nil
}

func validCoords(array [][]int, x int, y int) bool {
	return x >= 0 && x < len(array) && y >= 0 && y < len(array[0])
}

func solve(path string) (int, error) {
	heights, err := readFile(path)

	if err != nil {
		return 0, nil
	}

	result := 0
	var dx = []int{-1, 0, 1, 0}
	var dy = []int{0, 1, 0, -1}

	for rowIdx := 0; rowIdx < len(heights); rowIdx++ {
		for colIdx := 0; colIdx < len(heights[0]); colIdx++ {
			lowPoint := true
			for idx := 0; idx < len(dx); idx++ {
				row := rowIdx + dx[idx]
				col := colIdx + dy[idx]
				if validCoords(heights, row, col) {
					if heights[rowIdx][colIdx] >= heights[row][col] {
						lowPoint = false
						break
					}
				}
			}
			if lowPoint {
				result += 1 + heights[rowIdx][colIdx]
			}
		}
	}
	return result, nil
}

func main() {
	path := "day9/in-day9.txt"
	res, err := solve(path)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d\n", res)
}
