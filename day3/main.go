package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func getPriority(letter rune) int {
	if letter >= 97 {
		return int(letter) - 96
	}
	return int(letter) - 38
}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	fileContent := string(content)
	fileContent = strings.TrimSpace(fileContent)
	lines := strings.Split(fileContent, "\n")

	result := 0

	for i := 0; i < len(lines); {
		countGroup := map[rune]int{}
		for j := 0; j < 3; j++ {
			lineValues := map[rune]bool{}

			for _, c := range lines[i] {
				if !lineValues[c] {
					lineValues[c] = true
					countGroup[c] += 1
				}
			}

			i++
		}

		for key, val := range countGroup {
			if val == 3 {
				result += getPriority(key)
				continue
			}
		}
	}

	fmt.Println("Total value:", result)
}
