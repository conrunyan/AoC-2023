package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func loadData(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Failed to open file")
	}
	var lines []string
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func isValidGame(line string) (int, bool) {
	regex := regexp.MustCompile(`Game (\d+):(.*)`)

	// Find submatches in the email string
	matches := regex.FindStringSubmatch(line)
	if len(matches) > 0 {
		// The first element in matches is the entire match,
		// subsequent elements are the captured groups
		gameId, _ := strconv.Atoi(matches[1])
		gameDraws := strings.Split(matches[2], ";")

		validTotals := map[string]int{
			"red":   12,
			"green": 13,
			"blue":  14,
		}

		// Check each draw for a valid game
		for _, draw := range gameDraws {
			// Check each color for a valid amount
			for key, maxNum := range validTotals {
				numRegex := regexp.MustCompile(`(\d+) ` + key)
				numMatches := numRegex.FindStringSubmatch(draw)
				if len(numMatches) > 0 {
					matchedNum, _ := strconv.Atoi(numMatches[1])
					// Break early if the number is past the max
					if matchedNum > maxNum {
						return gameId, false
					}
				}
			}

		}

		fmt.Println("gameId:", gameId)
		fmt.Println("gameDraws:", gameDraws)

		return gameId, true
	} else {
		fmt.Println("No matches found.")
	}

	return 0, false
}

func getMinCubesPower(line string) int {
	// var minCubes []int
	draws := strings.Split(line, ";")
	colorMaxes := map[string]int{"green": 1, "red": 1, "blue": 1}
	for _, draw := range draws {
		for color := range colorMaxes {
			regex := regexp.MustCompile(`(\d+) ` + color)
			matches := regex.FindStringSubmatch(draw)
			if len(matches) > 0 {
				num, _ := strconv.Atoi(matches[1])
				colorMaxes[color] = max(num, colorMaxes[color])
			}
		}
	}
	fmt.Println("colorMaxes", colorMaxes)

	drawPower := 1
	for _, value := range colorMaxes {
		drawPower *= value
	}

	return drawPower
}

func main() {
	inputData := loadData("./input.txt")
	total := 0
	for _, line := range inputData {
		// Part 1
		// gameId, isValid := isValidGame(line)
		// if isValid {
		// 	total += gameId
		// }
		cubePower := getMinCubesPower(line)
		total += cubePower
	}
	fmt.Println("Total", total)
}
