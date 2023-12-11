package main

import (
	"adventofcode2023/util"
	"errors"
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

	topZig := []string{"F", "J"}
	bottomZig := []string{"7", "L"}

	fmt.Println("UpdatedMap")
	for y := 0; y < len(pipeMap); y++ {
		fmt.Println("")
		rowMap := pipeMap[y]
		for x := 0; x < len(rowMap); x++ {
			pipePiece := rowMap[x]
			fmt.Printf("%s", pipePiece.piece)
		}

	}

	pipeMap[sY][sX].piece = findReplacement(sX, sY, pipeMap)

	var included int
	fmt.Println("")
	for y := 0; y < len(pipeMap); y++ {
		rowMap := pipeMap[y]
		var pipeCount int
		var blockSkip string
		for x := 0; x < len(rowMap); x++ {
			piece := pipeMap[y][x].piece

			// count top zigs
			if piece == "F" {
				blockSkip = piece
			} else if piece == "J" {
				if blockSkip[0] == 'F' {
					pipeCount++
				}
				blockSkip = ""
				continue
			} else if piece == "L" {
				blockSkip = piece
			} else if piece == "7" {
				if blockSkip[0] == 'L' {
					pipeCount++
				}
				blockSkip = ""
				continue
			}

			if piece == "-" {
				blockSkip += piece
			}

			if len(blockSkip) > 0 {
				continue
			}

			if slices.Contains(topZig, piece) {
				blockSkip += piece
			} else if slices.Contains(bottomZig, piece) {
				blockSkip += piece
			} else if piece == "|" {
				pipeCount++
			} else if piece == "." && pipeCount%2 == 1 {
				pipeMap[y][x].piece = "X"
				included++
			}

			//if piece == "." && isContained(x, y, pipeMap) {
			//	piece = "X"
			//	included++
			//}
		}
	}
	//isContained(7, 4, pipeMap)

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

	newX, newY, err := findNewCoordinate(curX, curY, oldX, oldY, pipe)
	if err != nil {
		return nil, err
	}

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

func findNewCoordinate(curX, curY, oldX, oldY int, pipeDetail *PipeDetail) (int, int, error) {
	east := curX > oldX
	south := curY > oldY
	north := curY < oldY
	west := curX < oldX

	switch pipeDetail.piece {
	case "|":
		if east || west {
			return 0, 0, errors.New("invalid direction")
		}

		if south {
			return curX, curY + 1, nil
		} else {
			return curX, curY - 1, nil
		}
	case "-":
		if north || south {
			return 0, 0, errors.New("invalid direction")
		}

		if east {
			return curX + 1, curY, nil
		} else {
			return curX - 1, curY, nil
		}
	case "L":
		if east || north {
			return 0, 0, errors.New("invalid direction")
		}
		if south {
			return curX + 1, curY, nil
		} else {
			return curX, curY - 1, nil
		}
	case "7":
		if west || south {
			return 0, 0, errors.New("invalid direction")
		}

		if east {
			return curX, curY + 1, nil
		} else {
			return curX - 1, curY, nil
		}
	case "J":
		if west || north {
			return 0, 0, errors.New("invalid direction")
		}
		if south {
			return curX - 1, curY, nil
		} else {
			return curX, curY - 1, nil
		}
	case "F":
		if east || south {
			return 0, 0, errors.New("invalid direction")
		}
		if west {
			return curX, curY + 1, nil
		} else {
			return curX + 1, curY, nil
		}
	}

	panic("no coord")
}

func findReplacement(x, y int, pipeMap map[int]map[int]*PipeDetail) string {
	top, topOk := pipeMap[y-1][x]
	right, rightOk := pipeMap[y][x+1]
	bottom, bottomOk := pipeMap[y+1][x]
	left, leftOk := pipeMap[y][x-1]

	fromLeft := leftOk && slices.Contains([]string{"-", "F", "L"}, left.piece)
	fromRight := rightOk && slices.Contains([]string{"-", "J", "7"}, right.piece)
	fromTop := topOk && slices.Contains([]string{"|", "F", "7"}, top.piece)
	fromBottom := bottomOk && slices.Contains([]string{"|", "J", "L"}, bottom.piece)

	if fromLeft && fromRight {
		return "-"
	} else if fromLeft && fromTop {
		return "J"
	} else if fromLeft && fromBottom {
		return "7"
	} else if fromTop && fromRight {
		return "L"
	} else if fromBottom && fromTop {
		return "|"
	} else if fromRight && fromBottom {
		return "F"
	}

	panic("why")
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
