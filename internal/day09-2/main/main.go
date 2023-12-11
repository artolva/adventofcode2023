package main

import (
	"adventofcode2023/util"
	"fmt"
	"log"
	"time"
)

// Last attempt: 1746657573

const (
	// allowStringNumbers = false
	fileName = "misc/oasisSequence"
)

func main() {
	now := time.Now()
	file, scanner := util.GetRowsFromFile(fileName)
	defer file.Close()

	var totalSequence int
	//var numberSequences [][]int
	for scanner.Scan() {
		text := scanner.Text()

		values := util.ExtractNumbersByDelimiter(text, " ")

		sequenceProgression := findNextInSequence(values)

		fmt.Printf("======Sequence Val: %d\n\n", sequenceProgression)
		totalSequence = totalSequence + sequenceProgression
	}

	fmt.Printf("Total Sequence: %d\n", totalSequence)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Processing time: %d\n", time.Now().UnixMilli()-now.UnixMilli())
}

func findNextInSequence(sequence []int) int {
	var intervals []int
	var deeper bool
	sequenceLen := len(sequence)
	for index := 0; index < sequenceLen-1; index++ {
		nextVal := sequence[index+1]
		currentVal := sequence[index]

		newInterval := nextVal - currentVal
		if newInterval != 0 {
			deeper = true
		}

		intervals = append(intervals, newInterval)
	}

	firstVal := sequence[0]
	if !deeper {
		return firstVal
	}

	fmt.Printf("Current Intervals: %+v\n", intervals)
	if len(intervals) > 2 {
		a := sequence[1] + (-1 * intervals[0])
		fmt.Printf("Check Value: %d is %+v\n", a, a == sequence[0])
	}

	nextInSequence := -1 * findNextInSequence(intervals)

	fmt.Printf("Next: %d, First: %d\n", nextInSequence, firstVal)
	return firstVal + nextInSequence
}
