package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode"
)


func buildCrane(crane string) map[int][]string {
	lines := strings.Split(crane, "\n")
	lines = lines[:len(lines)-1]

	result := make(map[int][]string)

	for i := len(lines)-1; i >= 0; i-- {
		craneIdx := 1
		current := lines[i]
		for j := 0; j < len(current); {
			c := string(current[j+1])
			if unicode.IsLetter(rune(c[0])) {
				result[craneIdx] = append(result[craneIdx], c)
			}

			craneIdx++
			j += 4
		}
	}

	return result
}

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	content := string(b)

	splited := strings.Split(content, "\n\n")
	crane := splited[0]
	moves := splited[1]
	moves = strings.TrimSpace(moves)

	craneMap := buildCrane(crane)

	for _, move := range strings.Split(moves, "\n") {
		cols := strings.Split(move, " ")
		count, _ := strconv.Atoi(cols[1])
		from , _ := strconv.Atoi(cols[3])
		target, _ := strconv.Atoi(cols[5])

		moved := make([]string, 0)
		for i := 0; i < count; i++ {
			top := craneMap[from][len(craneMap[from])-1]
			craneMap[from] = craneMap[from][:len(craneMap[from])-1]
			moved = append(moved, top)
		}
		for i := len(moved)-1; i >= 0; i-- {
			craneMap[target] = append(craneMap[target], moved[i])
		}
		
	}

	for i := 1; i <= len(craneMap); i++ {
		fmt.Print(craneMap[i][len(craneMap[i])-1])
	}
}
