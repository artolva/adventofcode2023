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
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)

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

		if leftVal < 0 {
			v := splitString[i]
			checkLeft, err := strconv.Atoi(v)
			if err == nil {
				leftString = ""
				leftVal = checkLeft
			} else {
				leftString += v
				if number, err := getNumber(leftString); err == nil {
					leftVal = number
				}
			}
		}
		if rightVal < 0 {
			v := splitString[rightKey]
			checkRight, err := strconv.Atoi(v)
			if err == nil {
				rightString = ""
				rightVal = checkRight
			} else {
				rightString = fmt.Sprintf("%s%s", v, rightString)
				if number, err := getNumber(rightString); err == nil {
					rightVal = number
				}
			}
		}

		if leftVal > -1 && rightVal > -1 {
			atoi, _ := strconv.Atoi(fmt.Sprintf("%d%d", leftVal, rightVal))
			return atoi
		}
	}

	panic("shouldn't happen")
}

func getNumber(value string) (int, error) {
	if !allowStringNumbers {
		return 0, fmt.Errorf("dissallow")
	}

	if strings.HasPrefix(value, "one") || strings.HasSuffix(value, "one") {
		return 1, nil
	} else if strings.HasPrefix(value, "two") || strings.HasSuffix(value, "two") {
		return 2, nil
	} else if strings.HasPrefix(value, "three") || strings.HasSuffix(value, "three") {
		return 3, nil
	} else if strings.HasPrefix(value, "four") || strings.HasSuffix(value, "four") {
		return 4, nil
	} else if strings.HasPrefix(value, "five") || strings.HasSuffix(value, "five") {
		return 5, nil
	} else if strings.HasPrefix(value, "six") || strings.HasSuffix(value, "six") {
		return 6, nil
	} else if strings.HasPrefix(value, "seven") || strings.HasSuffix(value, "seven") {
		return 7, nil
	} else if strings.HasPrefix(value, "eight") || strings.HasSuffix(value, "eight") {
		return 8, nil
	} else if strings.HasPrefix(value, "nine") || strings.HasSuffix(value, "nine") {
		return 9, nil
	}

	return 0, fmt.Errorf("not a number")
}
