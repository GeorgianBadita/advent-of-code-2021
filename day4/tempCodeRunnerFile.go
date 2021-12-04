package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Cell struct {
	value  int
	marked bool
}

func makeArrayFromLine(arrStr string, sep string) (arr []int, err error) {
	numStr := strings.Split(strings.TrimSpace(arrStr), sep)
	arr = make([]int, 0)
	for _, num := range numStr {
		if num == "" {
			continue
		}
		numInt, err := strconv.Atoi(num)
		if err != nil {
			return nil, err
		}
		arr = append(arr, numInt)
	}
	return arr, nil
}

func readFile(path string) (drawn []int, tables [][][]Cell, err error) {
	fd, err := os.Open(path)

	if err != nil {
		return nil, nil, err
	}

	scanner := bufio.NewScanner(fd)

	scanner.Scan()
	numStr := scanner.Text()
	drawn, err = makeArrayFromLine(numStr, ",")
	if err != nil {
		return nil, nil, err
	}

	tables = make([][][]Cell, 0)
	for scanner.Scan() {
		currLine := scanner.Text()
		if currLine == "" {
			continue
		}
		table := make([][]Cell, 5)
		arr, err := makeArrayFromLine(currLine, " ")

		if err != nil {
			return nil, nil, err
		}

		line := make([]Cell, len(arr))
		for i, elem := range arr {
			line[i] = Cell{elem, false}
		}
		table[0] = line

		for idx := 0; idx < 4; idx++ {
			scanner.Scan()
			newLine := scanner.Text()
			arr, err := makeArrayFromLine(newLine, " ")

			if err != nil {
				return nil, nil, err
			}
			line := make([]Cell, len(arr))
			for i, elem := range arr {
				line[i] = Cell{elem, false}
			}
			table[idx+1] = line
			if err != nil {
				return nil, nil, err
			}
		}

		tables = append(tables, table)
	}

	return drawn, tables, nil
}

func checkWin(table [][]Cell, x int, y int) bool {
	cntMarkedLine := 0
	for yy := 0; yy < len(table[0]); yy++ {
		if table[x][yy].marked {
			cntMarkedLine += 1
		}
	}
	cntMarkedCol := 0
	for xx := 0; xx < len(table); xx++ {
		if table[xx][y].marked {
			cntMarkedCol += 1
		}
	}
	if cntMarkedLine == len(table[0]) || cntMarkedCol == len(table) {
		return true
	}
	return false
}

func computeSolution(table [][]Cell, lastDrawn int) int {
	sum := 0
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[0]); j++ {
			if table[i][j].marked == false {
				sum += table[i][j].value
			}
		}
	}
	return sum * lastDrawn
}

func solve(path string) (int, error) {
	draw, tables, err := readFile(path)
	if err != nil {
		return 0, err
	}

	for _, drawn := range draw {
		currentDrawn := drawn
		for tableIdx := 0; tableIdx < len(tables); tableIdx++ {
			currentTable := tables[tableIdx]
			for i := 0; i < len(currentTable); i++ {
				for j := 0; j < len(currentTable[0]); j++ {
					if currentTable[i][j].value == currentDrawn {
						currentTable[i][j].marked = true
						if checkWin(currentTable, i, j) {
							return computeSolution(currentTable, currentDrawn), nil
						}
					}

				}
			}
		}
	}
	return 0, nil
}

func main() {
	path := "day4/in-day4.txt"

	sol, err := solve(path)

	if err != nil {
		panic(err)
	}
	print(sol)
}
