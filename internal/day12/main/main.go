package main

import (
	"adventofcode2023/util"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"slices"
	"strconv"
	"strings"
	"time"
)

// With Numbers: 54450
// Without: 	 54265

const (
	fileName = "misc/springs"
)

type Block struct {
	id      string
	startAt int
	chars   string
	notEnd  bool
	opBlock bool
}

type BlockSet struct {
	blocks      []*Block
	combos      []int
	fullChars   string
	lineLen     int
	comboTotal  int
	combination map[string]struct{}
}

func main() {
	now := time.Now()
	lines := util.GetRowsFromFile(fileName)

	var blockSets []BlockSet
	for _, line := range lines {
		split := strings.Split(line, " ")

		var leftSide, rightSide string
		for i := 0; i < 5; i++ {
			if i > 0 {
				leftSide += "?"
				rightSide += ","
			}
			leftSide += split[0]
			rightSide += split[1]
		}

		fmt.Printf("Line: %s %s\n", leftSide, rightSide)

		lineLen := len(split[0])
		var opBlock bool
		var currentBlock string
		var blocks []*Block
		for i, s := range strings.Split(split[0], "") {
			isOp := s == "."

			if len(currentBlock) == 0 {
				currentBlock = s
				opBlock = isOp
			} else if isOp == opBlock {
				currentBlock = fmt.Sprintf("%s%s", currentBlock, s)
			} else if len(currentBlock) > 0 {
				blocks = append(blocks, &Block{
					startAt: i - len(currentBlock),
					id:      uuid.NewString(),
					chars:   currentBlock,
					opBlock: opBlock,
				})
				opBlock = !opBlock
				currentBlock = s
			}
		}

		blocks = append(blocks, &Block{
			id:      uuid.NewString(),
			opBlock: opBlock,
			startAt: lineLen - len(currentBlock),
			chars:   currentBlock,
		})

		var comboTotal int
		var combos []int
		for _, s := range strings.Split(split[1], ",") {
			atoi, _ := strconv.Atoi(s)
			comboTotal += atoi
			combos = append(combos, atoi)
		}

		//slices.Sort(combos)
		//slices.Reverse(combos)
		blockSets = append(blockSets, BlockSet{
			blocks:     blocks,
			combos:     combos,
			lineLen:    lineLen,
			fullChars:  split[0],
			comboTotal: comboTotal,
		})
	}

	var totalCombos int
	for _, blockSet := range blockSets {
		fmt.Printf("======       %s, %+v\n", blockSet.fullChars, blockSet.combos)
		blocksToCheck := slices.Clone(blockSet.blocks)
		for _, block := range blocksToCheck {
			if block.opBlock {
				blockSet.fullChars = util.ReplaceStartingAt(blockSet.fullChars, block.chars, block.startAt)
				blocksToCheck = setWithoutBlock(blocksToCheck, block)
			}
		}

		blockSet.combination = make(map[string]struct{})
		err := blockSet.findSetFromRemainder(&blockSet.fullChars, blocksToCheck, blockSet.combos)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err.Error())
		}

		for s, _ := range blockSet.combination {
			fmt.Printf("Combination: %s\n", s)
		}

		totalCombos += len(blockSet.combination)
	}
	fmt.Printf("Total Combos: %d\n", totalCombos)

	fmt.Printf("Processing time: %d\n", time.Now().UnixMilli()-now.UnixMilli())
}

func (bs *BlockSet) findSetFromRemainder(buildString *string, blocks []*Block, combos []int) error {
	comboLength := len(combos)
	blockLength := len(blocks)
	if comboLength == 0 {
		all := strings.ReplaceAll(*buildString, "?", ".")

		for _, block := range blocks {
			if strings.Contains(block.chars, "#") {
				fmt.Printf("left over required fields: %s\n\n", all)
				return nil
			}
		}

		if strings.Count(all, "#") != bs.comboTotal {
			fmt.Printf("too many #s: %s\n\n", all)
			//	return errors.New("too many #s")
		}

		bs.combination[all] = struct{}{}
		return nil
	} else if blockLength == 0 {
		return errors.New("ran out of blocks")
	}

	//var noValidCombo bool
	fitX := combos[0]
	for blockIndex := 0; blockIndex < len(blocks); blockIndex++ {
		block := blocks[blockIndex]
		mustSucceed := strings.Contains(block.chars, "#")
		//noValidCombo = true

		charLength := len(block.chars)
		//fmt.Printf("block chars: %s\n", block.chars)
		if fitX <= charLength {
			for i := 0; i <= charLength-fitX; i++ {
				clone := strings.Clone(*buildString)
				toIndex := i + fitX

				if toIndex > charLength {
					break
				}

				skipVal := 0
				toLastIndex := toIndex == charLength
				if block.notEnd || !toLastIndex {
					skipVal = 1

					if toIndex+skipVal > charLength {
						fmt.Println("requires too much space")
						continue
					}

					if block.chars[toIndex] == '#' {
						continue
					}
				}

				var useStr string
				for ind := 0; ind < fitX; ind++ {
					useStr += "#"
				}

				var newBlocks []*Block
				if charLength-(fitX+skipVal) > 0 {
					var leftChars, rightChars string
					for j, char := range block.chars {
						if j < i {
							leftChars = fmt.Sprintf("%s%c", leftChars, char)
						} else if j >= i+fitX+skipVal {
							rightChars = fmt.Sprintf("%s%c", rightChars, char)
						}
					}

					if strings.Contains(leftChars, "#") {
						//fmt.Println("bad option")
						continue
					}
					//if len(leftChars) > 1 {
					//withoutBlock = append(withoutBlock, &Block{
					//	id:      uuid.NewString(),
					//	chars:   leftChars,
					//	notEnd:  true,
					//	startAt: block.startAt,
					//})
					//}

					if len(rightChars) > 0 || (len(rightChars) == 1 && toLastIndex) {
						newBlocks = append(newBlocks, &Block{
							id:      uuid.NewString(),
							chars:   rightChars,
							notEnd:  false,
							startAt: block.startAt + i + fitX + skipVal,
						})
					}
				}

				newBlocks = append(newBlocks, blocks[blockIndex+1:]...)

				clone = util.ReplaceStartingAt(clone, useStr, block.startAt+i)
				if err := bs.findSetFromRemainder(&clone, newBlocks, combos[1:]); err == nil {
					//noValidCombo = false
				}
			}
		}

		if mustSucceed {
			break
		}
	}

	return nil
}

func setWithoutBlock(blocks []*Block, block *Block) []*Block {
	var newBlocks []*Block
	for _, b := range blocks {
		if b.id != block.id {
			newBlocks = append(newBlocks, b)
		}
	}
	return newBlocks
}
