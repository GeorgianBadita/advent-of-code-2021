package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Point struct {
	x int
	y int
}

type QueuePoint struct {
	point Point
	dist  int
}

func readFile(path string) (risks [][]int, err error) {
	fd, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(fd)

	risks = make([][]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		height := make([]int, 0)

		for _, elem := range line {
			height = append(height, int(elem-'0'))
		}
		risks = append(risks, height)
	}
	return risks, nil
}

func validCoords(arr [][]int, point Point) bool {
	return point.x >= 0 && point.x < len(arr) && point.y >= 0 && point.y < len(arr[0])
}

func pickSmallest(queue []QueuePoint) QueuePoint {
	best := queue[0]
	for idx := 1; idx < len(queue); idx++ {
		if queue[idx].dist < best.dist {
			best = queue[idx]
		}
	}

	return best
}

func removeBest(queue []QueuePoint, best QueuePoint) []QueuePoint {
	newSlice := make([]QueuePoint, len(queue)-1)
	currIdx := 0
	for _, elem := range queue {
		if elem.dist == best.dist && elem.point.x == best.point.x && elem.point.y == best.point.y {
			continue
		}
		newSlice[currIdx] = elem
		currIdx++
	}
	return newSlice
}

func solve(path string) (int, error) {
	risks, err := readFile(path)

	if err != nil {
		return 0, err
	}

	dx := []int{-1, 0, 0, 1}
	dy := []int{0, 1, -1, 0}

	queue := make([]QueuePoint, 0)
	dst := make([][]int, len(risks))
	for idx := 0; idx < len(dst); idx++ {
		dst[idx] = make([]int, len(risks[0]))
		for jdx := 0; jdx < len(risks[0]); jdx++ {
			dst[idx][jdx] = math.MaxInt
		}
	}

	dst[0][0] = 0
	queue = append(queue, QueuePoint{Point{0, 0}, dst[0][0]})

	for len(queue) > 0 {
		curr := pickSmallest(queue)
		queue = removeBest(queue, curr)

		if curr.dist != dst[curr.point.x][curr.point.y] {
			continue
		}

		for idx := 0; idx < len(dx); idx++ {
			newX := curr.point.x + dx[idx]
			newY := curr.point.y + dy[idx]
			if validCoords(risks, Point{newX, newY}) {
				if dst[newX][newY] > dst[curr.point.x][curr.point.y]+risks[newX][newY] {
					dst[newX][newY] = dst[curr.point.x][curr.point.y] + risks[newX][newY]
					queue = append(queue, QueuePoint{Point{newX, newY}, dst[newX][newY]})
				}
			}
		}
	}

	return dst[len(risks)-1][len(risks[0])-1], nil
}

func main() {
	path := "day15/in-day15.txt"
	res, err := solve(path)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", res)
}
