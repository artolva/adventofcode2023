package main

import (
	"adventofcode2023/util"
	"cmp"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

// With Numbers: 54450
// Without: 	 54265

type CardRank int32
type CardValue int32

//go:generate stringer -type=Pill
const (
	FIVE_OF_A_KIND  CardRank = 7
	FOUR_OF_A_KIND  CardRank = 6
	FULL_HOUSE      CardRank = 5
	THREE_OF_A_KIND CardRank = 4
	TWO_PAIR        CardRank = 3
	ONE_PAIR        CardRank = 2
	HIGH_CARD       CardRank = 1
)

var cardRankValues = map[int]CardRank{
	7: FIVE_OF_A_KIND,
	6: FOUR_OF_A_KIND,
	5: FULL_HOUSE,
	4: THREE_OF_A_KIND,
	3: TWO_PAIR,
	2: ONE_PAIR,
	1: HIGH_CARD,
}

//go:generate stringer -type=Pill
const (
	CARD_A CardValue = 13
	CARD_K CardValue = 12
	CARD_Q CardValue = 11
	CARD_J CardValue = 10
	CARD_T CardValue = 9
	CARD_9 CardValue = 8
	CARD_8 CardValue = 7
	CARD_7 CardValue = 6
	CARD_6 CardValue = 5
	CARD_5 CardValue = 4
	CARD_4 CardValue = 3
	CARD_3 CardValue = 2
	CARD_2 CardValue = 1
)

const (
	fileName = "misc/camelCards"
)

type CardHand struct {
	hand     []CardValue
	bid      int
	handRank CardRank
}

func main() {
	//now := time.Now()
	file, scanner := util.GetFile(fileName)
	defer file.Close()

	handMap := make(map[CardRank][]CardHand)
	for scanner.Scan() {
		cardHand := extractHandSummary(scanner.Text())

		set, _ := handMap[cardHand.handRank]
		handMap[cardHand.handRank] = append(set, cardHand)
	}

	var total int
	rankCounter := 1
	for i := 1; i <= 7; i++ {
		cardHands := handMap[cardRankValues[i]]

		slices.SortFunc(cardHands, func(a, b CardHand) int {
			if compare := cmp.Compare(a.hand[0], b.hand[0]); compare != 0 {
				return compare
			}
			if compare := cmp.Compare(a.hand[1], b.hand[1]); compare != 0 {
				return compare
			}

			if compare := cmp.Compare(a.hand[2], b.hand[2]); compare != 0 {
				return compare
			}

			if compare := cmp.Compare(a.hand[3], b.hand[3]); compare != 0 {
				return compare
			}

			return cmp.Compare(a.hand[4], b.hand[4])
		})

		for i := 0; i < len(cardHands); i++ {
			hand := cardHands[i]
			fmt.Printf("Rank %d for hand: %+v\n", rankCounter, hand)
			total = total + (rankCounter * cardHands[i].bid)
			rankCounter++
		}

	}

	fmt.Printf("Total: %d", total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func extractHandSummary(hand string) CardHand {
	split := strings.Split(hand, " ")

	cards := strings.Split(split[0], "")
	bid, _ := strconv.Atoi(split[1])
	var cardSet []CardValue
	cardMap := make(map[CardValue]int)

	for _, char := range cards {
		cardValue := convertToCardValue(char)

		cardMap[cardValue]++
		cardSet = append(cardSet, cardValue)
	}

	return CardHand{
		bid:      bid,
		hand:     cardSet,
		handRank: getCardRank(cardMap),
	}
}

func getCardRank(cards map[CardValue]int) CardRank {
	uniqueCards := len(cards)

	switch uniqueCards {
	case 1:
		return FIVE_OF_A_KIND
	case 2:
		for _, i := range cards {
			if i == 4 || i == 1 {
				return FOUR_OF_A_KIND
			}
		}
		return FULL_HOUSE
	case 3:
		for _, i := range cards {
			if i == 3 {
				return THREE_OF_A_KIND
			}
		}
		return TWO_PAIR
	case 4:
		return ONE_PAIR
	case 5:
		return HIGH_CARD
	}

	panic("we don do2")
}

func convertToCardValue(char string) CardValue {
	switch char {
	case "A":
		return CARD_A
	}
	switch char {
	case "K":
		return CARD_K
	}
	switch char {
	case "Q":
		return CARD_Q
	}
	switch char {
	case "J":
		return CARD_J
	}
	switch char {
	case "T":
		return CARD_T
	}
	switch char {
	case "9":
		return CARD_9
	}
	switch char {
	case "8":
		return CARD_8
	}
	switch char {
	case "7":
		return CARD_7
	}
	switch char {
	case "6":
		return CARD_6
	}
	switch char {
	case "5":
		return CARD_5
	}
	switch char {
	case "4":
		return CARD_4
	}
	switch char {
	case "3":
		return CARD_3
	}
	switch char {
	case "2":
		return CARD_2
	}

	panic("we don do")
}
