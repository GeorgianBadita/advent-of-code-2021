package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Rule struct {
	left  string
	right string
}

type ExpandKey struct {
	from string
	to   string
}

func readFile(path string) (template string, rules []Rule, err error) {
	fd, err := os.Open(path)

	if err != nil {
		return "", nil, err
	}

	scanner := bufio.NewScanner(fd)

	scanner.Scan()
	template = strings.Trim(scanner.Text(), " ")

	rules = make([]Rule, 0)

	scanner.Scan()
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, "->")
		rules = append(rules, Rule{strings.Trim(lineSplit[0], " "), strings.Trim(lineSplit[1], " ")})
	}
	return template, rules, nil
}

func expand(template map[string]int, rules map[string]string) map[string]int {
	update := make(map[string]int)

	for pair, occ := range template {
		update[string(pair[0])+rules[pair]] += occ
		update[rules[pair]+string(pair[1])] += occ
	}

	return update
}

func solve(path string, steps int) (int, error) {
	template, rules, err := readFile(path)

	if err != nil {
		return 0, err
	}

	templateMap := make(map[string]int)
	rulesMap := make(map[string]string)

	for idx := 0; idx < len(template)-1; idx++ {
		templateMap[template[idx:idx+2]]++
	}
	for idx := 0; idx < len(rules); idx++ {
		rulesMap[rules[idx].left] = rules[idx].right
	}

	for idx := 0; idx < steps; idx++ {
		templateMap = expand(templateMap, rulesMap)
	}

	occMap := make(map[string]int)

	for pair, occ := range templateMap {
		occMap[string(pair[0])] += occ
	}

	max, min := 0, math.MaxInt

	for _, occ := range occMap {
		if occ < min {
			min = occ
		}
		if occ > max {
			max = occ
		}
	}

	return max - min + 1, nil
}

func main() {
	path := "day14/in-day14.txt"
	steps := 40
	res, err := solve(path, steps)

	if err != nil {
		panic(err)
	}
	fmt.Printf("%d\n", res)
}
