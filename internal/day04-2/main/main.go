package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	//allowStringNumbers = false
	fileName = "misc/lotteryTickets"
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
	copyMap := make(map[int]int)
	for scanner.Scan() {
		total = total + findCardValue(copyMap, scanner.Text())
	}

	fmt.Printf("Total value: %d", total)
}

func findCardValue(repeatMap map[int]int, cardText string) int {
	rowSplit := strings.Split(cardText, ": ")
	gameNumber := getNumber(rowSplit[0][4:])
	gameDetail := strings.Split(rowSplit[1], " | ")

	winningNumbers := extractValues(gameDetail[0])
	ticketNumbers := extractValues(gameDetail[1])

	//fmt.Printf("winningNumbers string: %+v\n", winningNumbers)
	//fmt.Printf("ticketNumbers string: %+v\n", ticketNumbers)

	var matches int
	for _, number := range ticketNumbers {
		for _, winningNumber := range winningNumbers {
			if number == winningNumber {
				//fmt.Printf("%d is a winner %d\n", number, winningNumber)
				matches++
			}
		}
	}

	if matches > 0 {
		fmt.Printf("Game %d has %d matches\n", gameNumber, matches)

		total := 1

		copyCount := repeatMap[gameNumber]
		for i := 1; i <= matches; i++ {
			total = total + 1

			var repeatCount int
			if value, ok := repeatMap[gameNumber+i]; ok {
				repeatCount = value
			}
			repeatMap[gameNumber+i] = repeatCount + copyCount + 1
		}

		points := 1 + copyCount
		//if gameNumber < 20 {
		fmt.Printf("Game Map: %+v", repeatMap)
		fmt.Printf("%d repeats for %d points\n========\n", copyCount, points)
		//}
		return points
	}

	fmt.Printf("game %d includes %d default copies\n========\n", gameNumber, repeatMap[gameNumber])
	return repeatMap[gameNumber] + 1
}

func extractValues(gameDetail string) []int {
	//fmt.Printf("====== Extract Values from: %s\n", gameDetail)
	var caret int
	var numbers []int
	var valueString string
	for _, char := range strings.Split(gameDetail, "") {
		if caret == 2 {
			numbers = append(numbers, getNumber(valueString))

			caret = 0
			valueString = ""
		} else {
			valueString = fmt.Sprintf("%s%s", valueString, char)
			caret++
		}
	}

	numbers = append(numbers, getNumber(valueString))

	return numbers
}

func getNumber(ticketNumber string) int {
	atoi, _ := strconv.Atoi(strings.ReplaceAll(ticketNumber, " ", ""))

	return atoi
}
