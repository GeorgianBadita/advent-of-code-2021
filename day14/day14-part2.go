package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Rule struct {
	left  string
	right string
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

func expand(template map[string]int, rules []Rule) map[string]int {
	for _, rule := range rules {
		occ, ok := template[rule.left]
		if ok {
			first := string([]byte{rule.left[0]}) + rule.right
			second := rule.right + string([]byte{rule.left[1]})

			occFirst, okFirst := template[first]
			if okFirst {
				template[first] += occFirst + occFirst
			} else {
				template[first] = occ
			}
			occSecond, okSecond := template[second]
			if okSecond {
				template[second] += occSecond + occFirst
			} else {
				template[second] = occ
			}
		}
	}
	return template
}

func solve(path string, steps int) (int, error) {
	template, rules, err := readFile(path)

	if err != nil {
		return 0, err
	}

	templateMap := make(map[string]int)

	for idx := 0; idx < len(template)-1; idx++ {
		templateMap[template[idx:idx+1]] = 1
	}

	for step := 0; step < steps; steps++ {
		templateMap = expand(templateMap, rules)
	}

	return 0, nil
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
