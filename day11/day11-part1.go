package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x int
	y int
}

func readFile(path string) (flash [][]int, err error) {
	fd, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(fd)

	flash = make([][]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		height := make([]int, 0)

		for _, elem := range line {
			height = append(height, int(elem-'0'))
		}
		flash = append(flash, height)
	}
	return flash, nil
}

func validCoords(arr [][]int, point Point) bool {
	return point.x >= 0 && point.x < len(arr) && point.y >= 0 && point.y < len(arr[0])
}

func makeFlashed(flashes [][]int) [][]bool {
	flashed := make([][]bool, len(flashes))
	for idx := 0; idx < len(flashes); idx++ {
		flashed[idx] = make([]bool, len(flashes[0]))
		for jdx := 0; jdx < len(flashes[0]); jdx++ {
			flashed[idx][jdx] = false
		}
	}
	return flashed
}

func solve(path string, steps int) (int, error) {
	flashes, err := readFile(path)
	if err != nil {
		return 0, err
	}

	dx := []int{-1, 0, 0, 1, 1, -1, 1, -1}
	dy := []int{0, 1, -1, 0, 1, 1, -1, -1}

	result := 0
	for step := 0; step < steps; step++ {
		flashedThisStep := make([]Point, 0)
		flashed := makeFlashed(flashes)

		for row := 0; row < len(flashes); row++ {
			for col := 0; col < len(flashes[0]); col++ {
				flashes[row][col] += 1
				if flashes[row][col] > 9 {
					result++
					flashes[row][col] = 0
					flashed[row][col] = true
					flashedThisStep = append(flashedThisStep, Point{row, col})
				}
			}
		}

		for len(flashedThisStep) > 0 {
			current := flashedThisStep[0]
			flashedThisStep = flashedThisStep[1:]
			for idx := 0; idx < len(dx); idx++ {
				currX := current.x + dx[idx]
				currY := current.y + dy[idx]
				if validCoords(flashes, Point{currX, currY}) && !flashed[currX][currY] {
					flashes[currX][currY] += 1
					if flashes[currX][currY] > 9 {
						result++
						flashes[currX][currY] = 0
						flashed[currX][currY] = true
						flashedThisStep = append(flashedThisStep, Point{currX, currY})
					}
				}
			}
		}
	}
	return result, nil
}

func main() {
	path := "day11/in-day11.txt"
	steps := 100

	result, err := solve(path, steps)

	if err != nil {
		panic(nil)
	}

	fmt.Printf("%d\n", result)
}
