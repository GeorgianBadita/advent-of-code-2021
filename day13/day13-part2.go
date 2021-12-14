package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Fold struct {
	foldType string
	value    int
}

func readFile(path string) (points map[Point]bool, folds []Fold, err error) {
	fd, err := os.Open(path)

	if err != nil {
		return nil, nil, err
	}

	scanner := bufio.NewScanner(fd)
	points = make(map[Point]bool, 0)

	for scanner.Scan() {
		lineStr := scanner.Text()
		if lineStr == "" {
			break
		}
		line := strings.Split(strings.Trim(lineStr, " "), ",")
		x, err := strconv.Atoi(line[0])

		if err != nil {
			return nil, nil, err
		}

		y, err := strconv.Atoi(line[1])

		if err != nil {
			return nil, nil, err
		}

		points[Point{x, y}] = true
	}

	folds = make([]Fold, 0)
	for scanner.Scan() {
		lineStr := scanner.Text()
		foldStr := strings.Split(strings.Trim(lineStr, " "), " ")[2]
		foldType := strings.Split(foldStr, "=")[0]
		foldValue, err := strconv.Atoi(strings.Split(foldStr, "=")[1])

		if err != nil {
			return nil, nil, err
		}

		folds = append(folds, Fold{foldType, foldValue})
	}

	return points, folds, nil
}

func printPoints(pointsMap map[Point]bool) {
	points := make([]Point, 0)

	for point := range pointsMap {
		points = append(points, point)
	}

	if len(points) == 0 {
		return
	}
	maxX := points[0].x
	maxY := points[0].y

	for _, p := range points {
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	toPrint := make([][]string, maxY+1)
	for idx := 0; idx <= maxY; idx++ {
		toPrint[idx] = make([]string, maxX+1)
		for jdx := 0; jdx <= maxX; jdx++ {
			toPrint[idx][jdx] = "."
		}
	}

	for _, p := range points {
		toPrint[p.y][p.x] = "#"
	}

	fmt.Print("\n")
	for row := 0; row <= maxY; row++ {
		for col := 0; col <= maxX; col++ {
			fmt.Printf("%s", toPrint[row][col])
		}
		fmt.Print("\n")
	}
}

func makeFoldY(points map[Point]bool, value int) {
	for point := range points {
		if point.y == value {
			delete(points, point)
		}
		if point.y > value {
			delete(points, point)
			point.y -= 2 * (point.y - value)
			points[point] = true
		}
	}
}

func makeFoldX(points map[Point]bool, value int) {
	for point := range points {
		if point.x == value {
			delete(points, point)
		}
		if point.x > value {
			delete(points, point)
			point.x -= 2 * (point.x - value)
			points[point] = true
		}
	}
}

func solve(path string) (int, error) {
	points, folds, err := readFile(path)

	if err != nil {
		return 0, err
	}

	for _, fold := range folds {
		if fold.foldType == "y" {
			makeFoldY(points, fold.value)
		} else {
			makeFoldX(points, fold.value)
		}
	}
	printPoints(points)
	return len(points), nil
}

func main() {
	path := "day13/in-day13.txt"

	res, err := solve(path)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", res)
}
