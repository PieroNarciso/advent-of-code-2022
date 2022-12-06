package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getIds(lower int, upper int) []int {
	result := make([]int, 0)

	for i := lower; i <= upper; i++ {
		result = append(result, i)
	}
	return result
}

func contains(s1 []int, v int) bool {
	for _, e := range s1 {
		if e == v {
			return true
		}
	}
	return false
}

func subslice(s1 []int, s2 []int) bool {
	if len(s1) > len(s2) {
		return false
	}
	for _, e := range s1 {
		if !contains(s2, e) {
			return false
		}
	}
	return true
}

func overlaps(s1 []int, s2 []int) bool {
	for _, e := range s1 {
		if contains(s2, e) {
			return true
		}
	}
	return false
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 0

	for scanner.Scan() {
		pairs := strings.Split(scanner.Text(), ",")
		left := pairs[0]
		right := pairs[1]

		leftPairs := strings.Split(left, "-")
		idLower, _ := strconv.Atoi(leftPairs[0])
		idUpper, _ := strconv.Atoi(leftPairs[1])
		leftIds := getIds(idLower, idUpper)

		rightPairs := strings.Split(right, "-")
		idLower, _ = strconv.Atoi(rightPairs[0])
		idUpper, _ = strconv.Atoi(rightPairs[1])
		rightIds := getIds(idLower, idUpper)

		if overlaps(leftIds, rightIds) {
			// fmt.Println(leftIds, rightIds)
			count++
		}

	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println("Total Value:", count)
}
