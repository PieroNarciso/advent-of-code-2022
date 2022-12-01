package main

import (
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)


func main() {
	file, err := ioutil.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}
	content := string(file)
	content = strings.TrimSpace(content)

	values := strings.Split(content, "\n")


	elvesValues := []int{}
	for i := 0; i < len(values); {
		currElve := 0
		for ; i < len(values) && values[i] != ""; {
			val, err := strconv.Atoi(values[i])
			if err != nil {
				log.Println(err)
			}
			currElve += val
			i++
		}
		elvesValues = append(elvesValues, currElve)
		i++
	}

	sort.Slice(elvesValues, func(i, j int) bool {
		return elvesValues[i] > elvesValues[j]
	})
	
	maxValue := elvesValues[0] + elvesValues[1] + elvesValues[2]

	log.Printf("Max value: %v\n", maxValue)
}
