package main

import (
	"adventofcode2023/types"
	"adventofcode2023/util"
	"fmt"
	"strconv"
	"time"
)

// Correct: 	 726820169514

const (
	fileName = "misc/galaxies"
)

type Galaxy struct {
	x       int
	y       int
	id      int
	xOffset int
	yOffset int
}

func main() {
	now := time.Now()
	lines := util.GetRowsFromFile(fileName)

	columnWidth := len(lines[0])
	numberOfRows := len(lines)

	galaxyRows := make(map[int]struct{})
	galaxyColumns := make(map[int]struct{})

	var galaxyId int
	var coordinates []types.Coordinate
	for row, line := range lines {
		for col, char := range line {
			if char == '#' {
				galaxyRows[row] = struct{}{}
				galaxyColumns[col] = struct{}{}
				coordinates = append(coordinates, types.Coordinate{
					X: col,
					Y: row,
				})
				galaxyId++
			}
		}
	}

	galaxyMap, galaxies := buildGalaxyMap(coordinates, numberOfRows, columnWidth, galaxyRows, galaxyColumns)

	fmt.Println("UpdatedMap")
	for y := 0; y < len(galaxyMap); y++ {
		fmt.Println("")
		rowMap := galaxyMap[y]
		for x := 0; x < len(rowMap); x++ {
			fmt.Printf("%s", rowMap[x])
		}

	}

	var total, comparisons int
	for sourceId := 0; sourceId < len(galaxies); sourceId++ {
		for targetId := sourceId + 1; targetId < len(galaxies); targetId++ {
			source := galaxies[sourceId]
			target := galaxies[targetId]

			galaxy := findDistanceToGalaxy(source, target)

			fmt.Printf("\nDistance from %d to %d: %d", source.id, target.id, galaxy)
			total += galaxy
			comparisons++
		}
	}
	fmt.Printf("\ntotal: %d\n", total)
	fmt.Printf("\ncomparisons: %d\n", comparisons)

	fmt.Printf("Processing time: %d\n", time.Now().UnixMilli()-now.UnixMilli())
}

func findDistanceToGalaxy(source, destination Galaxy) int {
	var lowX, highX, lowY, highY int
	if source.x < destination.x {
		lowX = source.x
		highX = destination.x
	} else {
		lowX = destination.x
		highX = source.x
	}

	if source.y < destination.y {
		lowY = source.y
		highY = destination.y
	} else {
		lowY = destination.y
		highY = source.y
	}

	return (highX - lowX) + (highY - lowY)
}

func buildGalaxyMap(coordinates []types.Coordinate, rowCount, colCount int, galaxyRows, galaxyColumns map[int]struct{}) (map[int]map[int]string, []Galaxy) {
	galaxyMap := make(map[int]map[int]string)

	galaxyId := 1
	stepSize := 999999

	var galaxies []Galaxy
	var rowOffset int
	for row := 0; row < rowCount; row++ {
		newRowMap := make(map[int]string)
		_, rightRow := galaxyRows[row]
		var colOffset int
		for column := 0; column < colCount; column++ {
			var isGalaxy bool
			for _, coordinate := range coordinates {
				if coordinate.X == column && coordinate.Y == row {
					isGalaxy = true
					break
				}
			}

			_, rightCol := galaxyColumns[column]
			if isGalaxy {
				newRowMap[column+colOffset] = strconv.Itoa(galaxyId)
				galaxies = append(galaxies, Galaxy{
					x:  column + colOffset,
					y:  row + rowOffset,
					id: galaxyId,
				})
				galaxyId++
				continue
			} else if !rightCol {
				//for i := 0; i < stepSize; i++ {
				//	newRowMap[column+colOffset+i] = "."
				//}
				colOffset += stepSize
			}
			newRowMap[column+colOffset] = "."
		}

		if !rightRow {
			//for i := 0; i < stepSize; i++ {
			//	galaxyMap[row+rowOffset+i] = newRowMap
			//}
			rowOffset += stepSize
		}

		galaxyMap[row+rowOffset] = newRowMap
	}

	return galaxyMap, galaxies
}
