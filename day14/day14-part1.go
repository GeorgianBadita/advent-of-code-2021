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

func expand(template string, rules []Rule) string {
	newTemplate := ""

	for idx := 0; idx < len(template)-1; idx++ {
		currChars := template[idx : idx+2]
		hasRule := false
		for _, rule := range rules {
			if rule.left == currChars {
				newTemplate += string([]byte{currChars[0]}) + rule.right
				hasRule = true
				break
			}
		}
		if !hasRule {
			newTemplate += currChars
		}
	}
	return newTemplate + string([]byte{template[len(template)-1]})
}

func solve(path string, steps int) (int, error) {
	template, rules, err := readFile(path)

	if err != nil {
		return 0, err
	}

	for idx := 0; idx < steps; idx++ {
		template = expand(template, rules)
	}

	occMap := make(map[string]int)

	for _, chr := range template {
		currChar := string([]rune{chr})
		occ, ok := occMap[currChar]
		if !ok {
			occMap[currChar] = 1
		} else {
			occMap[currChar] = occ + 1
		}
	}

	maxOcc := 0
	minOcc := len(template)

	for _, occ := range occMap {
		if occ > maxOcc {
			maxOcc = occ
		}
		if occ < minOcc {
			minOcc = occ
		}
	}
	return maxOcc - minOcc, nil
}

func main() {
	path := "day14/in-day14.txt"
	steps := 20
	res, err := solve(path, steps)

	if err != nil {
		panic(err)
	}
	fmt.Printf("%d\n", res)
}
