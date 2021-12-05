package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Line struct {
	p1 Point
	p2 Point
}

func min(a int, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func readFile(path string) (lines []Line, err error) {
	fd, err := os.Open(path)

	defer fd.Close()

	if err != nil {
		return nil, err
	}

	lines = make([]Line, 0)

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := scanner.Text()
		sides := strings.Split(line, "->")
		leftSide := strings.Split(sides[0], ",")
		rightSide := strings.Split(sides[1], ",")

		x1, err := strconv.Atoi(strings.TrimSpace(leftSide[0]))

		if err != nil {
			return nil, err
		}

		y1, err := strconv.Atoi(strings.TrimSpace(leftSide[1]))

		if err != nil {
			return nil, err
		}

		x2, err := strconv.Atoi(strings.TrimSpace(rightSide[0]))

		if err != nil {
			return nil, err
		}

		y2, err := strconv.Atoi(strings.TrimSpace(rightSide[1]))

		if err != nil {
			return nil, err
		}

		lines = append(lines, Line{Point{x1, y1}, Point{x2, y2}})
	}
	return lines, nil
}

func keepOnlyVerticalHorizontal(lines []Line) (filteredLines []Line) {
	filteredLines = make([]Line, 0)
	for _, line := range lines {
		if line.p1.x == line.p2.x || line.p1.y == line.p2.y {
			filteredLines = append(filteredLines, line)
		}
	}
	return filteredLines
}

func solve(path string) (int, error) {
	lines, err := readFile(path)
	if err != nil {
		return 0, err
	}

	visitedCount := make(map[Point]int)
	multipleTouches := 0
	fltLines := keepOnlyVerticalHorizontal(lines)
	for _, line := range fltLines {
		if line.p1.x == line.p2.x {
			for tmpY := min(line.p1.y, line.p2.y); tmpY <= max(line.p1.y, line.p2.y); tmpY++ {
				currPoint := Point{line.p1.x, tmpY}
				_, found := visitedCount[currPoint]
				if found {
					visitedCount[currPoint]++
				} else {
					visitedCount[currPoint] = 1
				}
			}
		} else {
			for tmpX := min(line.p1.x, line.p2.x); tmpX <= max(line.p1.x, line.p2.x); tmpX++ {
				currPoint := Point{tmpX, line.p1.y}
				_, found := visitedCount[currPoint]
				if found {
					visitedCount[currPoint]++
				} else {
					visitedCount[currPoint] = 1
				}
			}
		}
	}

	for _, value := range visitedCount {
		if value > 1 {
			multipleTouches += 1
		}
	}
	return multipleTouches, nil
}

func main() {
	path := "day5/in-day5.txt"

	result, err := solve(path)

	if err != nil {
		panic(err)
	}

	print(result)
}
