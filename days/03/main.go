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

func itemToPriority(item rune) int {
	// Uppercase
	if item >= 65 && item <= 90 {
		return int(item-65+1) + 26
	}

	// Lowercase
	return int(item - 97 + 1)
}

type Item struct {
	Char     rune
	Quantity int
	Priority int
}

func main() {
	rucksacks := strings.Split(input, "\n")

	result := 0

	for _, rucksack := range rucksacks {
		rucksackSize := len(rucksack)

		firstCompartment := map[rune]*Item{}
		secondCompartment := map[rune]*Item{}

		for i := 0; i < (rucksackSize / 2); i++ {
			item := rune(rucksack[i])

			x, ok := firstCompartment[item]
			if !ok {
				firstCompartment[item] = &Item{
					Char:     item,
					Quantity: 1,
					Priority: itemToPriority(item),
				}

				continue
			}

			x.Quantity += 1
		}

		for i := rucksackSize / 2; i < rucksackSize; i++ {
			item := rune(rucksack[i])

			x, ok := secondCompartment[item]
			if !ok {
				secondCompartment[item] = &Item{
					Char:     item,
					Quantity: 1,
					Priority: itemToPriority(item),
				}

				continue
			}

			x.Quantity += 1
		}

		for _, item := range firstCompartment {
			_, ok := secondCompartment[item.Char]
			if !ok {
				continue
			}

			result += item.Priority
		}
	}

	fmt.Println("Result: ", result)
}
