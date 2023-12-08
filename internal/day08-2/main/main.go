package main

import (
	"adventofcode2023/util"
	"fmt"
	"log"
	"strings"
	"time"
)

// last guess: 6325589301922576693

const (
	allowStringNumbers = false
	fileName           = "misc/mapInput"
)

type Decision struct {
	left  string
	right string
}

func main() {
	now := time.Now()
	file, scanner := util.GetFile(fileName)
	defer file.Close()

	scanner.Scan()
	instructions := strings.Split(scanner.Text(), "")

	// skip one line
	scanner.Scan()

	var startKey int
	startingMap := make(map[int]string)
	decisionMap := make(map[string]Decision)
	for scanner.Scan() {
		text := scanner.Text()
		options := strings.Split(text, " = ")

		decisionMap[options[0]] = Decision{
			left:  options[1][1:4],
			right: options[1][6:9],
		}

		if options[0][2:] == "A" {
			startingMap[startKey] = options[0]
			startKey++
		}
	}
	fmt.Printf("startingNodes: %+v\n", startingMap)

	index := 0
	zMapping := make(map[int]int)
	for {
		currentIndex := index % len(instructions)
		instruction := instructions[currentIndex]

		index++
		for i := 0; i < len(startingMap); i++ {
			if zMapping[i] > 0 {
				continue
			}

			if instruction == "L" {
				startingMap[i] = decisionMap[startingMap[i]].left
			} else {
				startingMap[i] = decisionMap[startingMap[i]].right
			}
			if startingMap[i][2:] == "Z" {
				//fmt.Printf("starting map %s\n", startingMap[i])
				//fmt.Printf("Index: %d\n", index)
				//fmt.Printf("CurrentIndex: %+v\n", currentIndex+1 == len(instructions))
				zMapping[i] = index
			}
		}

		//fmt.Printf("index %d and Z Mapping: %+v\n", index, zMapping)

		if len(zMapping) == len(startingMap) {
			break
		}
	}

	var iterationList []int
	for _, loopCount := range zMapping {
		iterationList = append(iterationList, loopCount)
	}

	lcm := LCM(iterationList[0], iterationList[1], iterationList[2:]...)

	fmt.Printf("Reached ZZZ in %d\n", lcm)
	fmt.Printf("Processing time: %d\n", time.Now().UnixMilli()-now.UnixMilli())

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
