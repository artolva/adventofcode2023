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
	for scanner.Scan() {
		total = total + findCardValue(scanner.Text())
	}

	fmt.Printf("Total value: %d", total)
}

func findCardValue(cardText string) int {
	rowSplit := strings.Split(cardText, ": ")
	gameDetail := strings.Split(rowSplit[1], " | ")

	winningNumbers := extractValues(gameDetail[0])
	ticketNumbers := extractValues(gameDetail[1])

	fmt.Printf("winningNumbers string: %+v\n", winningNumbers)
	fmt.Printf("ticketNumbers string: %+v\n", ticketNumbers)

	var matches int
	for _, number := range ticketNumbers {
		for _, winningNumber := range winningNumbers {
			if number == winningNumber {
				fmt.Printf("%d is a winner %d\n", number, winningNumber)
				matches++
			}
		}
	}

	if matches > 0 {
		total := 1
		for i := 1; i < matches; i++ {
			total = total * 2
		}

		fmt.Printf("%d total matches for %d points\n", matches, total)
		return total
	}

	return 0
}

func extractValues(gameDetail string) []int {
	fmt.Printf("====== Extract Values from: %s\n", gameDetail)
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
