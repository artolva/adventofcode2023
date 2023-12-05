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

type MapRef struct {
	destinationStart int
	rangeStart       int
	rangeLength      int
}

type SeedSet struct {
	seedValue int
	seedRange int
}

func main() {
	now := time.Now()
	file, scanner := util.GetFile(fileName)
	defer file.Close()

	scanner.Scan()

	var seedSets []SeedSet
	seedPairList := strings.Split(strings.Split(scanner.Text(), ": ")[1], " ")
	for i := 0; i < len(seedPairList); i = i + 2 {
		seedSets = append(seedSets, SeedSet{
			seedValue: getInt(seedPairList[i]),
			seedRange: getInt(seedPairList[i+1]),
		})
	}

	lowestNumber := -1
	listMap := make(map[string][]MapRef)
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
			listMap[mapHeader] = append(listMap[mapHeader], MapRef{
				destinationStart: getInt(components[0]),
				rangeStart:       getInt(components[1]),
				rangeLength:      getInt(components[2]),
			})
		}

	}

	seedSets = getKeysFromSet(seedSets, listMap["seed-to-soil"])
	seedSets = getKeysFromSet(seedSets, listMap["soil-to-fertilizer"])
	seedSets = getKeysFromSet(seedSets, listMap["fertilizer-to-water"])
	seedSets = getKeysFromSet(seedSets, listMap["water-to-light"])
	seedSets = getKeysFromSet(seedSets, listMap["light-to-temperature"])
	seedSets = getKeysFromSet(seedSets, listMap["temperature-to-humidity"])
	seedSets = getKeysFromSet(seedSets, listMap["humidity-to-location"])

	for _, seedSet := range seedSets {
		if lowestNumber == -1 || seedSet.seedValue < lowestNumber {
			lowestNumber = seedSet.seedValue
		}
	}
	fmt.Printf("Lowest location: %d\n", lowestNumber)

	fmt.Printf("Processing time: %d\n", time.Now().UnixMilli()-now.UnixMilli())
}

func getKeysFromSet(seedSets []SeedSet, options []MapRef) []SeedSet {
	var resultSet []SeedSet
	fmt.Println("\n\n\n\n==========================\nRange Start")

	sort.Slice(options, func(i, j int) bool {
		return options[i].rangeStart < options[j].rangeStart
	})
	for _, seedSet := range seedSets {
		seedStart := seedSet.seedValue
		seedEnd := seedSet.seedValue + seedSet.seedRange
		fmt.Printf("==========================\nFor Seed Start: %d\nFor Seed End: %d\n", seedStart, seedEnd)

		for _, option := range options {
			rangeStart := option.rangeStart
			rangeEnd := rangeStart + option.rangeLength
			// encompass
			if seedStart <= rangeStart && seedEnd >= rangeEnd {
				resultSet = append(resultSet, SeedSet{
					seedValue: option.destinationStart,
					seedRange: option.rangeLength,
				})
			} else if seedStart < rangeStart && seedEnd > rangeStart && seedEnd < rangeEnd {
				offset := rangeEnd - seedEnd
				resultSet = append(resultSet, SeedSet{
					seedValue: option.destinationStart,
					seedRange: option.rangeLength - offset,
				})
			} else if seedStart > rangeStart && seedStart < rangeEnd && seedEnd > rangeEnd {
				offset := seedStart - rangeStart
				resultSet = append(resultSet, SeedSet{
					seedValue: option.destinationStart + offset,
					seedRange: option.rangeLength - offset,
				})
			} else if seedStart > rangeStart && seedEnd < rangeEnd {
				offset := seedStart - rangeStart
				resultSet = append(resultSet, SeedSet{
					seedValue: option.destinationStart + offset,
					seedRange: seedSet.seedRange,
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
