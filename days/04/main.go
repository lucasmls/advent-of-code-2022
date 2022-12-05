package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

var (
	//go:embed input.txt
	input string
)

func main() {
	var (
		// 38-41,38-38
		// 18-65,18-65
		assignmentGroups = strings.Split(input, "\n")
		result           = 0
	)

	for _, group := range assignmentGroups {
		// [38-41 38-38]
		elfAssignments := strings.Split(group, ",")

		// 38-41
		firstElfAssignment := NewAssignment(elfAssignments[0])
		// 38-38
		secondElfAssignment := NewAssignment(elfAssignments[1])

		if firstElfAssignment.PartiallyOverlap(secondElfAssignment) || secondElfAssignment.PartiallyOverlap(firstElfAssignment) {
			result += 1
		}
	}

	fmt.Println("Result:", result)
}

type Assignment struct {
	Raw   string
	Start int
	End   int
}

func NewAssignment(raw string) Assignment {
	strStartAndEnd := strings.Split(raw, "-")

	start, _ := strconv.Atoi(strStartAndEnd[0])
	end, _ := strconv.Atoi(strStartAndEnd[1])

	return Assignment{
		Raw:   raw,
		Start: start,
		End:   end,
	}
}

func (a Assignment) FullyOverlap(counterpartyAssignment Assignment) bool {
	if a.Start == counterpartyAssignment.Start && a.End == counterpartyAssignment.End {
		return true
	}

	if a.Start <= counterpartyAssignment.Start && a.End >= counterpartyAssignment.End {
		return true
	}

	return false
}

func (a Assignment) PartiallyOverlap(counterpartyAssignment Assignment) bool {
	if a.Start == counterpartyAssignment.Start && a.End == counterpartyAssignment.End {
		return true
	}

	if inBetweenRange(a.Start, a.End, counterpartyAssignment.Start) {
		return true
	}

	if inBetweenRange(a.Start, a.End, counterpartyAssignment.End) {
		return true
	}

	return false
}

func inBetweenRange(start, end, value int) bool {
	if value >= start && value <= end {
		return true
	}

	return false
}
