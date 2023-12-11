package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetRowsFromFile(fileName string) []string {
	file, err := os.Open(fileName)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	var rows []string
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}

	return rows
}

func ExtractNumbersByDelimiter(line, delimiter string) []int {
	var nextVal string
	var results []int
	fmt.Printf("Line: %s\n", line)
	split := strings.Split(line, "")
	for i := 0; i < len(line); i++ {
		char := split[i]

		if char == "-" {
			nextVal = "-"
			continue
		}

		if intVal, err := strconv.Atoi(char); err == nil {
			nextVal = fmt.Sprintf("%s%d", nextVal, intVal)
		}

		if (i+1) == len(line) || (char != delimiter && split[i+1] == delimiter) {
			intVal, _ := strconv.Atoi(nextVal)
			results = append(results, intVal)
			nextVal = ""
		}
	}
	return results
}
