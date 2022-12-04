package main

import (
	_ "embed"
	"fmt"
	"strings"
)

var (
	//go:embed input.txt
	input string
)

func main() {
	var (
		sacks  = strings.Split(input, "\n")
		result = 0
	)

	groups := chunk(sacks, 3)

	for _, group := range groups {
		groupSet := map[int]map[string]int{}

		for i, item := range group {
			groupSet[i] = stringCharacterSet(item)
		}

		for k, _ := range groupSet[0] {
			if _, ok := groupSet[1][k]; !ok {
				continue
			}

			if _, ok := groupSet[2][k]; !ok {
				continue
			}

			result += calcLetterPriority(k)
		}
	}

	fmt.Println("Result: ", result)
}

func stringCharacterSet(s string) map[string]int {
	set := map[string]int{}

	for _, r := range s {
		char := string(r)
		if _, ok := set[char]; !ok {
			set[char] = 0
		}

		set[char] += 1
	}

	return set
}

func chunk(slice []string, size int) [][]string {
	var (
		chunk  []string
		result [][]string
	)

	for i, item := range slice {
		chunk = append(chunk, item)

		if (i+1)%size == 0 {
			result = append(result, chunk)

			chunk = []string{}
			continue
		}
	}

	if len(chunk) > 0 && len(chunk) < size {
		result = append(result, chunk)
	}

	return result
}

func calcLetterPriority(letter string) int {
	char := rune(letter[0])

	// Uppercase
	if char >= 65 && char <= 90 {
		return int(char-65+1) + 26
	}

	// Lowercase
	return int(char - 97 + 1)
}
