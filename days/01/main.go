package main

import (
	_ "embed"
	"fmt"
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

	mostCaloriesBeingCarried := 0
	currentElfCalorieSum := 0

	for _, foodCalorieLine := range elfsFood {
		if isEmptyLine(foodCalorieLine) {
			if currentElfCalorieSum > mostCaloriesBeingCarried {
				mostCaloriesBeingCarried = currentElfCalorieSum
			}

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

	fmt.Println("The most calories being carried is:", mostCaloriesBeingCarried)
}

func isEmptyLine(s string) bool {
	return s == ""
}
