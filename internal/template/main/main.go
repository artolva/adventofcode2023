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
	allowStringNumbers = false
	fileName           = "misc/calibrationDocument"
)

func main() {
	now := time.Now()
	file, scanner := util.GetFile(fileName)
	defer file.Close()
	for scanner.Scan() {
		//text := scanner.Text()
	}

	fmt.Printf("Processing time: %d\n", time.Now().UnixMilli()-now.UnixMilli())
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
