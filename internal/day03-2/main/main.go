package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// With Numbers: 54450
// Without: 	 54265

const (
	//allowStringNumbers = false
	fileName = "misc/engineSchematic"
)

type NumberDetail struct {
	rowNumber int
	colStart  int
	colEnd    int
	number    int
}

type SymbolDetail struct {
	symbol     string
	gearValues []int
	gearCheck  int
	x          int
	y          int
}

func main() {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	rowCount := 1
	var numberDetails []NumberDetail
	symbolDetails := make(map[int][]*SymbolDetail)
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), "")

		var numString string
		rowLength := len(text)
		for colCount := 0; colCount < rowLength; colCount++ {
			char := text[colCount]

			if char == "." {
				continue
			} else if isNumber(char) {
				numString = fmt.Sprintf("%s%s", numString, char)
			} else if char == "*" {
				symbolDetails[rowCount] = append(symbolDetails[rowCount], &SymbolDetail{
					symbol: char,
					x:      colCount,
					y:      rowCount,
				})
			}

			hasNum := len(numString) > 0
			endNum := (colCount+1) == rowLength || !isNumber(text[colCount+1])
			if hasNum && endNum {
				number, _ := strconv.Atoi(numString)
				numberDetails = append(numberDetails, NumberDetail{
					rowNumber: rowCount,
					colStart:  colCount - (len(numString) - 1),
					colEnd:    colCount,
					number:    number,
				})
				numString = ""
			}
		}

		rowCount++
	}

	for _, numberDetail := range numberDetails {
		checkGear(numberDetail, symbolDetails)
	}

	var total int
	for _, rowDetail := range symbolDetails {
		for _, symbolDetail := range rowDetail {
			fmt.Printf("Symbol %+v\n", symbolDetail)

			if symbolDetail.gearCheck == 2 {
				total = total + (symbolDetail.gearValues[0] * symbolDetail.gearValues[1])
			}
		}
	}

	fmt.Printf("Total: %d\n", total)
}

func isNumber(check string) bool {
	_, err := strconv.Atoi(check)

	return err == nil
}

func checkGear(numDetail NumberDetail, symbols map[int][]*SymbolDetail) {
	for i := -1; i < 2; i++ {
		checkRow := numDetail.rowNumber + i

		if symbolRow, ok := symbols[checkRow]; ok {
			leftAnchor := numDetail.colStart - 1
			rightAnchor := numDetail.colEnd + 1

			for _, symbolDetail := range symbolRow {
				if symbolDetail.x >= leftAnchor && symbolDetail.x <= rightAnchor {
					symbolDetail.gearCheck = symbolDetail.gearCheck + 1
					symbolDetail.gearValues = append(symbolDetail.gearValues, numDetail.number)
				}
			}
		}
	}
}
