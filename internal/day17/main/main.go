package main

import (
	"adventofcode2023/util"
	"fmt"
	"time"
)

const (
	fileName = "misc/USE_FILE"
)

func main() {
	now := time.Now()
	lines := util.GetRowsFromFile(fileName)

	for _, line := range lines {

	}

	fmt.Printf("FILE: %+v", lines)

	fmt.Printf("Processing time: %d\n", time.Now().UnixMilli()-now.UnixMilli())
}
