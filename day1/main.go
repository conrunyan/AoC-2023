package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getNumIfIsStartOfWord(idx int, line string) string {
	var words = map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	// check if current idx + len of key equals the key
	for word, num := range words {
		section := line[idx:min(idx+len(word), len(line))]
		if section == word {
			return num
		}
	}
	return ""
}

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

func getFirstLastInts(line string) int {
	var first string
	var last string
	var firstFound bool
	var foundOne bool
	for idx, c := range line {
		var currentChar = string(c)
		_, err := strconv.Atoi(string(currentChar))
		if err != nil {
			// Not a digit, could it be the start of a word?
			num := getNumIfIsStartOfWord(idx, line)
			if num != "" {
				currentChar = num
				foundOne = true
			}
			// Is a digit, so we'll mark to save the value
		} else {
			foundOne = true
		}

		// First go round is the only time we'll modify the first value.
		// Otherwise we'll just move the last value forward to the latest
		// found one.
		if firstFound == false && foundOne == true {
			first = currentChar
			last = currentChar
			firstFound = true
		} else if foundOne == true {
			last = currentChar
		}

		foundOne = false
	}

	val, err := strconv.Atoi(first + last)
	if err != nil {
		fmt.Println("Something went horribly wrong...")
	}

	return val

}

func main() {
	var sum int
	fileLines := loadData("./input2.txt")
	for _, thing := range fileLines {
		value := getFirstLastInts(thing)
		sum += value
	}
	fmt.Printf("Sum %v\n", sum)
}
