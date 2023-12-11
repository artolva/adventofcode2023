package main

import (
	"adventofcode2023/util"
	"fmt"
	"log"
	"strings"
	"time"
)

const (
	fileName = "misc/pipeMap"
)

func main() {
	now := time.Now()
	file, scanner := util.GetFile(fileName)
	defer file.Close()

	var rowIndex, sX, sY int
	pipeMap := make(map[int]map[int]string)
	for scanner.Scan() {
		row, newX := buildMapRow(scanner.Text())
		pipeMap[rowIndex] = row

		if newX > -1 {
			sX, sY = newX, rowIndex
		}
		rowIndex++
	}

	fmt.Printf("Distance: %d\n", getDistanceForMap(sX, sY, pipeMap)/2)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Processing time: %d\n", time.Now().UnixMilli()-now.UnixMilli())
}

func getDistanceForMap(startX, startY int, pipeMap map[int]map[int]string) int {
	var err error
	var distance int

	fmt.Printf("Starting at %d, %d\n", startX, startY)
	for i := 0; i < len(pipeMap); i++ {
		fmt.Println("")
		rowMap := pipeMap[i]
		for j := 0; j < len(rowMap); j++ {
			fmt.Printf(" %d,%d:%s ", i, j, rowMap[j])
		}
	}

	if distance, err = delvePipe(startX, startY+1, startX, startY, 0, pipeMap); err == nil {
		fmt.Printf("found going down\n")
		return distance
	} else if distance, err = delvePipe(startX+1, startY, startX, startY, 0, pipeMap); err == nil {
		fmt.Printf("found going right\n")
		return distance
	} else if distance, err = delvePipe(startX, startY-1, startX, startY, 0, pipeMap); err == nil {
		fmt.Printf("found going up\n")
		return distance
	} else if distance, err = delvePipe(startX-1, startY, startX, startY, 0, pipeMap); err == nil {
		fmt.Printf("found going left\n")
		return distance
	} else {
		panic("not found")
	}

}

func delvePipe(curX, curY, oldX, oldY, i int, theMap map[int]map[int]string) (int, error) {
	columnMap, ok := theMap[curY]
	if !ok {
		return -1, fmt.Errorf("row out of bounds")
	}
	pipePiece, ok := columnMap[curX]
	if !ok {
		return -1, fmt.Errorf("column out of bounds")
	}

	if pipePiece == "S" {
		return 1, nil
	}

	//if i > 20 {
	//	return -1, errors.New("failed")
	//}

	if pipePiece == "." {
		return -1, fmt.Errorf("ground at %d, %d\n", curX, curY)
	}

	newX, newY := findNewCoordinate(curX, curY, oldX, oldY, pipePiece)

	fmt.Printf("\n=======\nDelving pipe %s, going to %d, %d...\n", pipePiece, newX, newY)

	if distance, err := delvePipe(newX, newY, curX, curY, i+1, theMap); err != nil {
		return -1, err
	} else {
		return distance + 1, err
	}
}

func findNewCoordinate(curX, curY, oldX, oldY int, char string) (int, int) {
	east := curX > oldX
	south := curY > oldY
	west := curX < oldX

	switch char {
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

func buildMapRow(lineText string) (map[int]string, int) {
	x := -1

	charArray := strings.Split(lineText, "")
	rowMap := make(map[int]string)
	for index, val := range charArray {
		rowMap[index] = val

		if val == "S" {
			x = index
		}
	}
	return rowMap, x
}
