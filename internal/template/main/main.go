package main

import (
	"adventofcode2023/util"
	"fmt"
	"log"
	"time"
)

// With Numbers: 54450
// Without: 	 54265

const (
	fileName = "misc/USE_A_REAL_FILE"
)

func main() {
	now := time.Now()
	file, scanner := util.GetRowsFromFile(fileName)
	defer file.Close()
	for scanner.Scan() {
		//text := scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Processing time: %d\n", time.Now().UnixMilli()-now.UnixMilli())
}
