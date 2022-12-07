package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func isMarker(quad string) bool {
	set := map[rune]bool{}

	for _, r := range quad {
		set[r] = true
	}

	if len(set) == 14 {
		return true
	}
	return false
}

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	content := strings.TrimSpace(string(b))
	lines := strings.Split(content, "\n")

	result := 0

	for _, line := range lines {
		for i := 0; i < len(line)-13; i++ {
			if isMarker(line[i:i+14]) {
				result += i+14
				fmt.Println(line[i:i+14], i+14)
				break
			}
		}
	}
	fmt.Println("Result is:", result)
}
