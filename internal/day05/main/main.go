package main

import (
	"adventofcode2023/util"
	"fmt"
	"strconv"
	"strings"
)

const (
	//allowStringNumbers = false
	fileName = "misc/almanac"
)

//seed-to-soil map:
//soil-to-fertilizer map:
//fertilizer-to-water map:
//water-to-light map:
//light-to-temperature map:
//temperature-to-humidity map:
//humidity-to-location map:

type MapRef struct {
	destinationRange int
	rangeStart       int
	rangeLength      int
}

func main() {
	file, scanner := util.GetFile(fileName)
	defer file.Close()

	scanner.Scan()
	seedList := strings.Split(strings.Split(scanner.Text(), ": ")[1], " ")

	lowestNumber := -1
	listMap := make(map[string][]MapRef)
	scanner.Scan()

	mapHeader := ""
	for scanner.Scan() {
		textLine := scanner.Text()

		fmt.Printf("text line: %s\n", textLine)
		if textLine == "" {
			continue
		} else if strings.Contains(textLine, ":") {
			mapHeader = strings.Split(textLine, " ")[0]
		} else {
			components := strings.Split(textLine, " ")
			listMap[mapHeader] = append(listMap[mapHeader], MapRef{
				destinationRange: getInt(components[0]),
				rangeStart:       getInt(components[1]),
				rangeLength:      getInt(components[2]),
			})
		}

	}

	//for key, list := range listMap {
	//	fmt.Printf("key map %s, %+v\n", key, list)
	//}
	fmt.Printf("seed seedList: %+v\n", seedList)
	for _, seedStr := range seedList {
		seed := getInt(seedStr)
		seed = getNextMapKey(seed, listMap["seed-to-soil"])
		seed = getNextMapKey(seed, listMap["soil-to-fertilizer"])
		seed = getNextMapKey(seed, listMap["fertilizer-to-water"])
		seed = getNextMapKey(seed, listMap["water-to-light"])
		seed = getNextMapKey(seed, listMap["light-to-temperature"])
		seed = getNextMapKey(seed, listMap["temperature-to-humidity"])
		seed = getNextMapKey(seed, listMap["humidity-to-location"])

		if lowestNumber == -1 || seed < lowestNumber {
			lowestNumber = seed
		}
	}

	fmt.Printf("Lowest location: %d", lowestNumber)
}

func getNextMapKey(sourceRef int, options []MapRef) int {
	for _, option := range options {
		if sourceRef >= option.rangeStart && sourceRef <= (option.rangeStart+option.rangeLength) {
			fmt.Printf("Key ref %d is in range of %+v\n", sourceRef, option)
			offset := sourceRef - option.rangeStart

			return option.destinationRange + offset
		}
	}

	return sourceRef
}

func getInt(val string) int {
	intVal, _ := strconv.Atoi(val)
	return intVal
}
