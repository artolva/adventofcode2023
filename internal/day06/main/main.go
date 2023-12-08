package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// With Numbers: 54450
// Without: 	 54265

const (
	allowStringNumbers = false
	fileName           = "misc/boatRecords"
)

func main() {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example

	timeLine := extractNumbers(getLine(scanner))
	distanceLine := extractNumbers(getLine(scanner))

	fmt.Printf("timeline: %+v\n", timeLine)
	fmt.Printf("distanceLine: %+v\n", distanceLine)

	var validOptions []int
	for i := 0; i < len(timeLine); i++ {
		validOption := 0
		fmt.Println("==========================")
		timeVal := timeLine[i]
		distanceVal := distanceLine[i]

		for i := 0; i < timeVal; i++ {
			remainder := timeVal - i

			consider := remainder*i > distanceVal
			if consider {
				fmt.Printf("%d speed for %d seconds totals %d\n", i, remainder, remainder*i)
				validOption++
			}
		}

		fmt.Printf("Valid options: %d\n", validOption)
		validOptions = append(validOptions, validOption)
	}

	total := 1
	for _, validCombinations := range validOptions {
		fmt.Printf("+++++\nCombination Value: %d\n", validCombinations)
		total = total * validCombinations
	}

	fmt.Printf("Total: %d", total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getLine(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func extractNumbers(line string) []int {
	var nextVal string
	var results []int
	fmt.Printf("Line: %s\n", line)
	s := strings.Split(line, ":")[1]
	split := strings.Split(s, "")
	for i := 0; i < len(split); i++ {
		char := split[i]

		if intVal, err := strconv.Atoi(char); err == nil {
			nextVal = fmt.Sprintf("%s%d", nextVal, intVal)
		}

		if (i+1) == len(split) || (char != " " && split[i+1] == " ") {
			intVal, _ := strconv.Atoi(nextVal)
			results = append(results, intVal)
			nextVal = ""
		}
	}
	return results
}
