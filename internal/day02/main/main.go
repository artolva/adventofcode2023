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
	//allowStringNumbers = false
	fileName   = "misc/cubeGames"
	redLimit   = 12
	greenLimit = 13
	blueLimit  = 14
)

func main() {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	var total, setTotal int
	for scanner.Scan() {
		text := scanner.Text()
		gameIdValue := gameWasPossible(text)
		gameSetValue := minimumRequiredCount(text)
		total = total + gameIdValue
		setTotal = setTotal + gameSetValue

	}

	fmt.Printf("Total Ids: %d\n", total)
	fmt.Printf("Total Sets: %d\n", setTotal)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func minimumRequiredCount(gameDetail string) int {
	gameOutline := strings.Split(gameDetail, ": ")
	//gameId, _ := strconv.Atoi(gameOutline[0][4:])

	//gameLine := strings.ReplaceAll(gameOutline[1], " ", "")

	minRed, minGreen, minBlue := 0, 0, 0
	for _, gameRecord := range strings.Split(gameOutline[1], "; ") {
		gameRecord = gameRecord
		for _, gameCount := range strings.Split(gameRecord, ", ") {
			split := strings.Split(gameCount, " ")
			count, _ := strconv.Atoi(split[0])
			colour := split[1]

			if colour == "red" && count > minRed {
				minRed = count
			} else if colour == "green" && count > minGreen {
				minGreen = count
			} else if colour == "blue" && count > minBlue {
				minBlue = count
			}
		}
	}

	return minOne(minRed) * minOne(minBlue) * minOne(minGreen)
}

func minOne(val int) int {
	if val == 0 {
		return 1
	}
	return val
}

func gameWasPossible(gameDetail string) int {
	gameOutline := strings.Split(gameDetail, ": ")
	gameId, _ := strconv.Atoi(gameOutline[0][5:])

	for _, gameRecord := range strings.Split(gameOutline[1], "; ") {
		for _, gameCount := range strings.Split(gameRecord, ", ") {
			split := strings.Split(gameCount, " ")
			count, _ := strconv.Atoi(split[0])
			colour := split[1]

			if colour == "red" && count > redLimit {
				return 0
			} else if colour == "green" && count > greenLimit {
				return 0
			} else if colour == "blue" && count > blueLimit {
				return 0
			}
		}
	}
	return gameId
}
