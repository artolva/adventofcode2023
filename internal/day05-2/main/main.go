package main

import (
	"adventofcode2023/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
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

type DestinationRef struct {
	destinationStart int
	rangeStart       int
	rangeLength      int
}

type RangeRef struct {
	seedValue int
	seedRange int
}

func main() {
	now := time.Now()
	file, scanner := util.GetRowsFromFile(fileName)
	defer file.Close()

	scanner.Scan()

	var rangeReferences []RangeRef
	seedPairList := strings.Split(strings.Split(scanner.Text(), ": ")[1], " ")
	for i := 0; i < len(seedPairList); i = i + 2 {
		rangeReferences = append(rangeReferences, RangeRef{
			seedValue: getInt(seedPairList[i]),
			seedRange: getInt(seedPairList[i+1]),
		})
	}

	lowestNumber := -1
	listMap := make(map[string][]DestinationRef)
	scanner.Scan()

	mapHeader := ""
	for scanner.Scan() {
		textLine := scanner.Text()

		if textLine == "" {
			continue
		} else if strings.Contains(textLine, ":") {
			mapHeader = strings.Split(textLine, " ")[0]
		} else {
			components := strings.Split(textLine, " ")
			listMap[mapHeader] = append(listMap[mapHeader], DestinationRef{
				destinationStart: getInt(components[0]),
				rangeStart:       getInt(components[1]),
				rangeLength:      getInt(components[2]),
			})
		}

	}

	rangeReferences = getKeysFromSet(rangeReferences, listMap["seed-to-soil"])
	rangeReferences = getKeysFromSet(rangeReferences, listMap["soil-to-fertilizer"])
	rangeReferences = getKeysFromSet(rangeReferences, listMap["fertilizer-to-water"])
	rangeReferences = getKeysFromSet(rangeReferences, listMap["water-to-light"])
	rangeReferences = getKeysFromSet(rangeReferences, listMap["light-to-temperature"])
	rangeReferences = getKeysFromSet(rangeReferences, listMap["temperature-to-humidity"])
	rangeReferences = getKeysFromSet(rangeReferences, listMap["humidity-to-location"])

	for _, rangeRef := range rangeReferences {
		if lowestNumber == -1 || rangeRef.seedValue < lowestNumber {
			lowestNumber = rangeRef.seedValue
		}
	}
	fmt.Printf("Lowest location: %d\n", lowestNumber)
	// SecureBanking#330169
	fmt.Printf("Processing time: %d\n", time.Now().UnixMilli()-now.UnixMilli())
}

func getKeysFromSet(rangeReferences []RangeRef, destinationReferences []DestinationRef) []RangeRef {
	var resultSet []RangeRef
	fmt.Println("\n\n\n\n==========================\nRange Start")

	sort.Slice(destinationReferences, func(i, j int) bool {
		return destinationReferences[i].rangeStart < destinationReferences[j].rangeStart
	})
	for _, rangeRef := range rangeReferences {
		seedStart := rangeRef.seedValue
		seedEnd := rangeRef.seedValue + rangeRef.seedRange
		fmt.Printf("==========================\nFor Seed Start: %d\nFor Seed End: %d\n", seedStart, seedEnd)

		for _, destinationRef := range destinationReferences {
			rangeStart := destinationRef.rangeStart
			rangeEnd := rangeStart + destinationRef.rangeLength
			// encompass
			if seedStart <= rangeStart && seedEnd >= rangeEnd {
				resultSet = append(resultSet, RangeRef{
					seedValue: destinationRef.destinationStart,
					seedRange: destinationRef.rangeLength,
				})
			} else if seedStart < rangeStart && seedEnd > rangeStart && seedEnd < rangeEnd {
				offset := rangeEnd - seedEnd
				resultSet = append(resultSet, RangeRef{
					seedValue: destinationRef.destinationStart,
					seedRange: destinationRef.rangeLength - offset,
				})
			} else if seedStart > rangeStart && seedStart < rangeEnd && seedEnd > rangeEnd {
				offset := seedStart - rangeStart
				resultSet = append(resultSet, RangeRef{
					seedValue: destinationRef.destinationStart + offset,
					seedRange: destinationRef.rangeLength - offset,
				})
			} else if seedStart > rangeStart && seedEnd < rangeEnd {
				offset := seedStart - rangeStart
				resultSet = append(resultSet, RangeRef{
					seedValue: destinationRef.destinationStart + offset,
					seedRange: rangeRef.seedRange,
				})
			}
		}
	}

	return resultSet
}

func getInt(val string) int {
	intVal, _ := strconv.Atoi(val)
	return intVal
}
