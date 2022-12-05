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

	for _, line := range lines {
		middleIdx := len(line) / 2
		left := line[:middleIdx]
		right := line[middleIdx:]

		storage := map[rune]bool{}
		alreadyStorage := map[rune]bool{}

		for _, c := range left {
			storage[c] = true
		}

		for _, c := range right {
			if storage[c] {
				alreadyStorage[c] = true
			}
		}

		for key := range alreadyStorage {
			fmt.Println("Putting:", string(key))
			result += getPriority(key)
		}
	}
	fmt.Println("Total value:", result)
}
