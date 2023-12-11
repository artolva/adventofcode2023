package main

import (
	"adventofcode2023/util"
	"fmt"
	"log"
	"slices"
	"strings"
	"time"
)

const (
	// Last attempt 478
	fileName = "misc/pipeMap"
)

type PipePosition struct {
	x int
	y int
}

type PipeDetail struct {
	piece     string
	isMapTile bool
	//eastWest   bool
}

func main() {
	now := time.Now()
	file, scanner := util.GetFile(fileName)
	defer file.Close()

	var rowIndex, sX, sY int
	pipeMap := make(map[int]map[int]*PipeDetail)
	for scanner.Scan() {
		row, newX := buildMapRow(scanner.Text())
		pipeMap[rowIndex] = row

		if newX > -1 {
			sX, sY = newX, rowIndex
		}
		rowIndex++
	}

	forMap := getDistanceForMap(sX, sY, pipeMap)

	for y := 0; y < len(pipeMap); y++ {
		for x := 0; x < len(pipeMap[y]); x++ {
			if _, ok := forMap[y][x]; !ok {
				pipeMap[y][x].piece = "."
			} else {
				pipeMap[y][x].isMapTile = true
			}
		}
	}

	var included int
	fmt.Println("")
	for y := 0; y < len(pipeMap); y++ {
		rowMap := pipeMap[y]
		for x := 0; x < len(rowMap); x++ {
			if pipeMap[y][x].piece == "." && isContained(x, y, pipeMap) {
				pipeMap[y][x].piece = "X"
				included++
			}
		}
	}
	//isContained(7, 4, pipeMap)

	fmt.Println("UpdatedMap")
	for y := 0; y < len(pipeMap); y++ {
		fmt.Println("")
		rowMap := pipeMap[y]
		for x := 0; x < len(rowMap); x++ {
			pipePiece := rowMap[x]
			fmt.Printf("%s", pipePiece.piece)
		}

	}

	fmt.Printf("\n======\n"+
		"included: %d\n", included)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Processing time: %d\n", time.Now().UnixMilli()-now.UnixMilli())
}

func getDistanceForMap(startX, startY int, pipeMap map[int]map[int]*PipeDetail) map[int]map[int]struct{} {
	var err error
	var distance []PipePosition

	var firstPipe PipePosition
	if distance, err = delvePipe(startX, startY+1, startX, startY, 0, pipeMap); err == nil {
		firstPipe = PipePosition{
			x: startX,
			y: startY + 1,
		}
		fmt.Printf("found going down\n")
	} else if distance, err = delvePipe(startX+1, startY, startX, startY, 0, pipeMap); err == nil {
		firstPipe = PipePosition{
			x: startX + 1,
			y: startY,
		}
		fmt.Printf("found going right\n")
	} else if distance, err = delvePipe(startX, startY-1, startX, startY, 0, pipeMap); err == nil {
		firstPipe = PipePosition{
			x: startX,
			y: startY - 1,
		}
		fmt.Printf("found going up\n")
	} else if distance, err = delvePipe(startX-1, startY, startX, startY, 0, pipeMap); err == nil {
		firstPipe = PipePosition{
			x: startX - 1,
			y: startY,
		}
		fmt.Printf("found going left\n")
	}

	positions := append(distance, firstPipe)

	consumedSpaces := make(map[int]map[int]struct{})
	for _, position := range positions {
		rowMap, ok := consumedSpaces[position.y]
		if !ok {
			rowMap = make(map[int]struct{})
		}

		rowMap[position.x] = struct{}{}
		consumedSpaces[position.y] = rowMap
	}

	return consumedSpaces
}

func isContained(x, y int, pipeMap map[int]map[int]*PipeDetail) bool {
	north := checkNorth(x, y, pipeMap)
	east := checkEast(x, y, pipeMap)
	south := checkSouth(x, y, pipeMap)
	west := checkWest(x, y, pipeMap)

	return north && east && south && west
}

func delvePipe(curX, curY, oldX, oldY, i int, theMap map[int]map[int]*PipeDetail) ([]PipePosition, error) {
	pipe, err := getPipePiece(curX, curY, theMap)
	if err != nil {
		return nil, err
	}

	if pipe.piece == "S" {
		return []PipePosition{{
			x: curX,
			y: curY,
		}}, nil
	}

	if pipe.piece == "." {
		return nil, fmt.Errorf("ground at %d, %d\n", curX, curY)
	}

	newX, newY := findNewCoordinate(curX, curY, oldX, oldY, pipe)

	if pipePositions, err := delvePipe(newX, newY, curX, curY, i+1, theMap); err != nil {
		return nil, err
	} else {
		return append(pipePositions, PipePosition{
			x: newX,
			y: newY,
		}), nil
	}
}

func getPipePiece(curX int, curY int, theMap map[int]map[int]*PipeDetail) (*PipeDetail, error) {
	columnMap, ok := theMap[curY]
	if !ok {
		return nil, fmt.Errorf("row out of bounds")
	}
	pipePiece, ok := columnMap[curX]
	if !ok {
		return pipePiece, fmt.Errorf("column out of bounds")
	}
	return pipePiece, nil
}

func findNewCoordinate(curX, curY, oldX, oldY int, pipeDetail *PipeDetail) (int, int) {
	east := curX > oldX
	south := curY > oldY
	west := curX < oldX

	switch pipeDetail.piece {
	case "|":
		if south {
			return curX, curY + 1
		} else {
			return curX, curY - 1
		}
	case "-":
		if east {
			return curX + 1, curY
		} else {
			return curX - 1, curY
		}
	case "L":
		if south {
			return curX + 1, curY
		} else {
			return curX, curY - 1
		}
	case "7":
		if east {
			return curX, curY + 1
		} else {
			return curX - 1, curY
		}
	case "J":
		if south {
			return curX - 1, curY
		} else {
			return curX, curY - 1
		}
	case "F":
		if west {
			return curX, curY + 1
		} else {
			return curX + 1, curY
		}
	}

	panic("no coord")
}

func buildMapRow(lineText string) (map[int]*PipeDetail, int) {
	x := -1

	charArray := strings.Split(lineText, "")
	rowMap := make(map[int]*PipeDetail)
	for index, val := range charArray {
		rowMap[index] = &PipeDetail{
			piece: val,
		}

		if val == "S" {
			x = index
		}
	}
	return rowMap, x
}

type checkFunc func(x, y int, pipeMap map[int]map[int]*PipeDetail) bool

func checkEastWest(checkX, checkY int, topCorner bool, pipeMap map[int]map[int]*PipeDetail) bool {
	var steps int
	var happyLeft, happyRight bool
	for {
		if happyLeft && happyRight {
			return true
		}

		var checkFn checkFunc
		if topCorner {
			checkFn = checkNorth
		} else {
			checkFn = checkSouth
		}

		if !happyLeft {
			detail, ok := pipeMap[checkY][checkX-steps]
			if !ok {
				return false
			}

			var happyPiece, sadPiece string
			if topCorner {
				happyPiece, sadPiece = "F", "L"
			} else {
				happyPiece, sadPiece = "L", "F"
			}

			if detail.piece == "S" || detail.piece == happyPiece {
				happyLeft = true
			} else if detail.piece == sadPiece && steps > 0 {
				fn := checkFn(checkX-steps, checkY, pipeMap)
				if !fn {
					return false
				}
				happyLeft = fn
			}
		}

		if !happyRight {
			detail, ok := pipeMap[checkY][checkX+steps]
			if !ok {
				return false
			}

			var happyPiece, sadPiece string
			if topCorner {
				happyPiece, sadPiece = "7", "J"
			} else {
				happyPiece, sadPiece = "J", "7"
			}

			if detail.piece == "S" || detail.piece == happyPiece {
				happyRight = true
			} else if detail.piece == sadPiece && steps > 0 {
				fn := checkFn(checkX+steps, checkY, pipeMap)
				if !fn {
					return false
				}
				happyRight = fn
			}
		}
		steps++
	}
}

// consider odd encounter #
func checkNorthSouth(checkX, checkY int, leftCorners bool, pipeMap map[int]map[int]*PipeDetail) bool {
	var steps int
	var happyUp, happyDown bool
	for {
		if happyUp && happyDown {
			return true
		}
		var checkFn checkFunc
		if leftCorners {
			checkFn = checkWest
		} else {
			checkFn = checkEast
		}
		if !happyUp {
			detail, ok := pipeMap[checkY-steps][checkX]
			if !ok {
				return false
			}

			var happyPiece, sadPiece string
			if leftCorners {
				happyPiece, sadPiece = "F", "7"
			} else {
				happyPiece, sadPiece = "7", "F" //FROM J TO F TO FIX 6,6 ON ex2
			}

			if detail.piece == "S" || detail.piece == happyPiece {
				happyUp = true
			} else if detail.piece == sadPiece && steps > 0 {
				fn := checkFn(checkX, checkY-steps, pipeMap)
				if !fn {
					return false
				}
				happyUp = fn
			}
		}

		if !happyDown {
			detail, ok := pipeMap[checkY+steps][checkX]
			if !ok {
				return false
			}

			var happyPiece, sadPiece string
			if leftCorners {
				happyPiece, sadPiece = "L", "J" //FROM F TO J TO FIX 11,5 ON ex3
			} else {
				happyPiece, sadPiece = "J", "L"
			}

			if detail.piece == "S" || detail.piece == happyPiece {
				happyDown = true
			} else if detail.piece == sadPiece && steps > 0 {
				fn := checkFn(checkX, checkY+steps, pipeMap)
				if !fn {
					return false
				}
				happyDown = fn
			}
		}
		steps++
	}
}

var northList = []string{"F", "7", "S", "-"}

func checkNorth(x, y int, pipeMap map[int]map[int]*PipeDetail) bool {
	var blockingPiece string
	for i := y; i >= 0; i-- {
		detail := pipeMap[i][x]
		northInclusive := slices.Contains(northList, detail.piece)
		if len(blockingPiece) > 0 {
			if blockingPiece == "J" && detail.piece == "7" {
				blockingPiece = ""
				continue
			} else if blockingPiece == "L" && detail.piece == "F" {
				blockingPiece = ""
				continue
			} else if northInclusive {
				blockingPiece = ""
			}
		}

		if detail.piece == "J" || detail.piece == "L" {
			blockingPiece = detail.piece
			continue
		}
		if detail.isMapTile && northInclusive {
			return checkEastWest(x, i, true, pipeMap)
		}
	}
	return false
}

var eastList = []string{"|", "J", "7", "S"}

func checkEast(x, y int, pipeMap map[int]map[int]*PipeDetail) bool {
	var blockingPiece string
	for i := x; i < len(pipeMap[y]); i++ {
		detail := pipeMap[y][i]
		eastInclusive := slices.Contains(eastList, detail.piece)
		if len(blockingPiece) > 0 {
			if blockingPiece == "L" && detail.piece == "J" {
				blockingPiece = ""
				continue
			} else if blockingPiece == "F" && detail.piece == "7" {
				blockingPiece = ""
				continue
			} else if eastInclusive {
				blockingPiece = ""
			}
		}

		if detail.piece == "L" || detail.piece == "F" {
			blockingPiece = detail.piece
			continue
		}

		if detail.isMapTile && eastInclusive {
			return checkNorthSouth(i, y, false, pipeMap)
		}
	}
	return false
}

var southList = []string{"L", "J", "S", "-"}

func checkSouth(x, y int, pipeMap map[int]map[int]*PipeDetail) bool {
	var blockingPiece string
	for i := y; i < len(pipeMap); i++ {
		detail := pipeMap[i][x]
		southInclusive := slices.Contains(southList, detail.piece)
		if len(blockingPiece) > 0 {
			if blockingPiece == "F" && detail.piece == "L" {
				blockingPiece = ""
				continue
			} else if blockingPiece == "7" && detail.piece == "J" {
				blockingPiece = ""
				continue
			} else if southInclusive {
				blockingPiece = ""
			}
		}

		if detail.piece == "7" || detail.piece == "F" {
			blockingPiece = detail.piece
			continue
		}
		if detail.isMapTile && southInclusive {
			return checkEastWest(x, i, false, pipeMap)
		}
	}
	return false
}

var westList = []string{"L", "F", "S", "|"}

func checkWest(x, y int, pipeMap map[int]map[int]*PipeDetail) bool {
	var blockingPiece string
	for i := x; i > 0; i-- {
		detail := pipeMap[y][i]
		westInclusive := slices.Contains(westList, detail.piece)
		if len(blockingPiece) > 0 {
			if blockingPiece == "7" && detail.piece == "F" {
				blockingPiece = ""
				continue
			} else if blockingPiece == "J" && detail.piece == "L" {
				blockingPiece = ""
				continue
			} else if westInclusive {
				blockingPiece = ""
			}
		}

		if detail.piece == "7" || detail.piece == "J" {
			blockingPiece = detail.piece
			continue
		}
		if detail.isMapTile && westInclusive {
			return checkNorthSouth(i, y, true, pipeMap)
		}
	}
	return false
}
