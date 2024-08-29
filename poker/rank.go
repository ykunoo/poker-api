package poker

import (
	"sort"
	"strconv"
)

func EvaluateHand(cards []Card) (string, int) {

	if isStraightFlush(cards) {
		if isRoyalStraightFlush(cards) {
			return YakuRoyalStraightFlush, RoyalStraightFlushRank
		}
		return YakuStraightFlush, StraightFlushRank
	}

	if isFourOfAKind(cards) {
		return YakuFourOfAKind, FourOfAKindRank
	}

	if isFullHouse(cards) {
		return YakuFullHouse, FullHouseRank
	}

	if isFlush(cards) {
		return YakuFlush, FlushRank
	}

	if isStraight(cards) {
		return YakuStraight, StraightRank
	}

	if isThreeOfAKind(cards) {
		return YakuThreeOfAKind, ThreeOfAKindRank
	}

	if isTwoPair(cards) {
		return YakuTwoPair, TwoPairRank
	}

	if isOnePair(cards) {
		return YakuOnePair, OnePairRank
	}

	return YakuHighCard, HighCardRank
}

func isRoyalStraightFlush(cards []Card) bool {
	// スペードの10, J, Q, K, Aの場合
	values := getCardValues(cards)
	sort.Ints(values)
	return isFlush(cards) && values[0] == 1 && values[1] == 10 && values[2] == 11 && values[3] == 12 && values[4] == 13
}

func isStraightFlush(cards []Card) bool {
	return isFlush(cards) && isStraight(cards)
}

func isFourOfAKind(cards []Card) bool {
	valueCounts := getValueCounts(cards)
	for _, count := range valueCounts {
		if count == 4 {
			return true
		}
	}
	return false
}

func isFullHouse(cards []Card) bool {
	valueCounts := getValueCounts(cards)
	hasThree := false
	hasTwo := false
	for _, count := range valueCounts {
		if count == 3 {
			hasThree = true
		} else if count == 2 {
			hasTwo = true
		}
	}
	return hasThree && hasTwo
}

func isFlush(cards []Card) bool {
	suit := cards[0].Suit
	for _, card := range cards {
		if card.Suit != suit {
			return false
		}
	}
	return true
}

func isStraight(cards []Card) bool {
	values := getCardValues(cards)
	sort.Ints(values)
	for i := 1; i < len(values); i++ {
		if values[i] != values[i-1]+1 {
			return false
		}
	}
	return true
}

func isThreeOfAKind(cards []Card) bool {
	valueCounts := getValueCounts(cards)
	for _, count := range valueCounts {
		if count == 3 {
			return true
		}
	}
	return false
}

func isTwoPair(cards []Card) bool {
	valueCounts := getValueCounts(cards)
	pairCount := 0
	for _, count := range valueCounts {
		if count == 2 {
			pairCount++
		}
	}
	return pairCount == 2
}

func isOnePair(cards []Card) bool {
	valueCounts := getValueCounts(cards)
	for _, count := range valueCounts {
		if count == 2 {
			return true
		}
	}
	return false
}

func getCardValues(cards []Card) []int {
	values := make([]int, len(cards))
	for i, card := range cards {
		intValue, _ := strconv.Atoi(card.Value)
		values[i] = intValue
	}
	return values
}

func getValueCounts(cards []Card) map[int]int {
	valueCounts := make(map[int]int)
	for _, card := range cards {
		intValue, _ := strconv.Atoi(card.Value)
		valueCounts[intValue]++
	}
	return valueCounts
}
