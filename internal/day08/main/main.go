package main

import (
	"adventofcode2023/util"
	"fmt"
	"log"
	"strings"
)

// With Numbers: 54450
// Without: 	 54265

const (
	allowStringNumbers = false
	fileName           = "misc/mapInput"
)

type Decision struct {
	left  string
	right string
}

func main() {
	//now := time.Now()
	file, scanner := util.GetRowsFromFile(fileName)
	defer file.Close()

	scanner.Scan()
	instructions := strings.Split(scanner.Text(), "")

	// skip one line
	scanner.Scan()

	decisionMap := make(map[string]Decision)
	for scanner.Scan() {
		text := scanner.Text()
		options := strings.Split(text, " = ")

		decisionMap[options[0]] = Decision{
			left:  options[1][1:4],
			right: options[1][6:9],
		}
	}

	for s, decision := range decisionMap {
		fmt.Printf("%s = (%s, %s)\n", s, decision.left, decision.right)
	}

	var index int
	currentDecision := "AAA"
	for {
		currentIndex := index % len(instructions)
		instruction := instructions[currentIndex]

		if instruction == "L" {
			currentDecision = decisionMap[currentDecision].left
		} else {
			currentDecision = decisionMap[currentDecision].right
		}

		index++
		//fmt.Printf("Decision %d moving to decision: %s\n", index, currentDecision)
		if currentDecision == "ZZZ" {
			break
		}
	}

	fmt.Printf("Reached ZZZ in %d\n", index)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
