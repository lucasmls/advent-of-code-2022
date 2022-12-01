package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"go.uber.org/zap"
)

var (
	//go:embed input.txt
	input string
)

var (
	logger *zap.SugaredLogger
)

func init() {
	baseLogger, _ := zap.NewProduction()
	defer baseLogger.Sync()
	logger = baseLogger.Sugar()
}

func main() {
	elfsFood := strings.Split(input, "\n")

	elfsCaloriesSum := []int{}
	currentElfCalorieSum := 0

	for _, foodCalorieLine := range elfsFood {
		if isEmptyLine(foodCalorieLine) {
			elfsCaloriesSum = append(elfsCaloriesSum, currentElfCalorieSum)
			currentElfCalorieSum = 0
			continue
		}

		calorie, err := strconv.Atoi(foodCalorieLine)
		if err != nil {
			logger.Error(
				"Unable to parse string to integer",
				zap.Error(err),
				zap.String("calorie", foodCalorieLine),
			)

			continue
		}

		currentElfCalorieSum += calorie
	}

	sort.Slice(elfsCaloriesSum, func(i, j int) bool {
		return elfsCaloriesSum[i] > elfsCaloriesSum[j]
	})

	sum := elfsCaloriesSum[0] + elfsCaloriesSum[1] + elfsCaloriesSum[2]

	fmt.Println("The top 3 calories being carried sums up to:", sum)
}

func isEmptyLine(s string) bool {
	return s == ""
}
