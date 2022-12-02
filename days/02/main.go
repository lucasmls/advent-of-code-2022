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

const (
	Rock    = "A"
	Paper   = "B"
	Scissor = "C"

	RockResponse    = "X"
	PaperResponse   = "Y"
	ScissorResponse = "Z"

	LoseOutcome = "lose"
	DrawOutcome = "draw"
	WinOutcome  = "win"

	RockScore    = 1
	PaperScore   = 2
	ScissorScore = 3

	LoseScore = 0
	DrawScore = 3
	WinScore  = 6
)

func main() {
	games := strings.Split(input, "\n")

	finalScore := 0

	for _, game := range games {
		options := strings.Split(game, " ")

		oponentChoice := options[0]
		response := responseToOption(options[1])

		gameResult := calculateGameResult(oponentChoice, response)
		gameScore := calculateScore(response, gameResult)

		finalScore += gameScore
	}

	fmt.Println("The final score is:", finalScore)
}

func responseToOption(response string) string {
	switch response {
	case RockResponse:
		return Rock
	case PaperResponse:
		return Paper
	default:
		return Scissor
	}
}

func calculateGameResult(oponentChoice string, response string) string {
	if oponentChoice == response {
		return DrawOutcome
	}

	if oponentChoice == Rock && response == Paper {
		return WinOutcome
	}

	if oponentChoice == Paper && response == Scissor {
		return WinOutcome
	}

	if oponentChoice == Scissor && response == Rock {
		return WinOutcome
	}

	return LoseOutcome
}

func calculateScore(shape string, outcome string) int {
	score := 0
	switch shape {
	case Rock:
		score += RockScore
	case Paper:
		score += PaperScore
	case Scissor:
		score += ScissorScore
	}

	switch outcome {
	case LoseOutcome:
		score += LoseScore
	case DrawOutcome:
		score += DrawScore
	case WinOutcome:
		score += WinScore
	}

	return score
}
