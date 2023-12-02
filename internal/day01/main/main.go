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
	fileName           = "misc/calibrationDocument"
)

func main() {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	var total int
	for scanner.Scan() {
		text := scanner.Text()
		sum := findFirstAndLastDigit(text)
		total = total + sum

		fmt.Printf("===\nGot %d from %s\n", sum, text)
	}

	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func findFirstAndLastDigit(value string) int {
	leftVal, rightVal := -1, -1

	var leftString, rightString string

	maxIndex := len(value) - 1
	splitString := strings.Split(value, "")
	for i := 0; i < len(value); i++ {
		rightKey := maxIndex - i

		leftVal = setValue(leftVal, splitString, i, leftString)
		rightVal = setValue(rightVal, splitString, rightKey, rightString)

		if leftVal > -1 && rightVal > -1 {
			atoi, _ := strconv.Atoi(fmt.Sprintf("%d%d", leftVal, rightVal))
			return atoi
		}
	}

	panic("shouldn't happen")
}

func setValue(val int, splitString []string, i int, leftBuilder string) int {
	if val < 0 {
		v := splitString[i]
		checkLeft, err := strconv.Atoi(v)
		if err == nil {
			leftBuilder = ""
			val = checkLeft
		} else {
			leftBuilder += v
			if number, err := getNumber(leftBuilder); err == nil {
				val = number
			}
		}
	}
	return val
}

func getNumber(value string) (int, error) {
	if !allowStringNumbers {
		return 0, fmt.Errorf("dissallow")
	}

	if strings.Contains(value, "one") {
		return 1, nil
	} else if strings.Contains(value, "two") {
		return 2, nil
	} else if strings.Contains(value, "three") {
		return 3, nil
	} else if strings.Contains(value, "four") {
		return 4, nil
	} else if strings.Contains(value, "five") {
		return 5, nil
	} else if strings.Contains(value, "six") {
		return 6, nil
	} else if strings.Contains(value, "seven") {
		return 7, nil
	} else if strings.Contains(value, "eight") {
		return 8, nil
	} else if strings.Contains(value, "nine") {
		return 9, nil
	}

	return 0, fmt.Errorf("not a number")
}
