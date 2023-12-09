package main

import (
	"adventofcode2023/util"
	"fmt"
	"log"
	"math"
	"time"
)

// Last attempt: 1746657573

const (
	// allowStringNumbers = false
	fileName = "misc/oasisSequence"
)

func main() {
	now := time.Now()
	file, scanner := util.GetFile(fileName)
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
	sequenceLen := len(sequence)
	for index := 0; index < (sequenceLen - 1); index++ {
		nextVal := sequence[index+1]
		currentVal := sequence[index]

		if nextVal < currentVal {
			//fmt.Printf("Next: %d, Last: %d\n", nextVal, currentVal)
			intervals = append(intervals, -1*int(math.Abs(float64(currentVal)-float64(nextVal))))
		} else {
			intervals = append(intervals, nextVal-currentVal)
		}
	}

	var deeper bool
	for _, interval := range intervals {
		if interval != 0 {
			deeper = true
		}
	}

	lastVal := sequence[sequenceLen-1]
	if !deeper {
		return lastVal
	}

	fmt.Printf("Current Intervals: %+v\n", intervals)
	if len(intervals) > 2 {
		a := sequence[sequenceLen-2] + intervals[len(intervals)-1]
		fmt.Printf("Check Value: %d is %+v\n", a, a == sequence[sequenceLen-1])
	}

	nextInSequence := findNextInSequence(intervals)
	if lastVal < nextInSequence {
		fmt.Printf("Next: %d, Last: %d\n", nextInSequence, lastVal)

		return nextInSequence + lastVal
	}
	fmt.Printf("Next: %d, Last: %d\n", nextInSequence, lastVal)
	return lastVal + nextInSequence
}
