package poker

import (
	"errors"
	"strconv"
	"strings"
)

func ParseHand(hand string) ([]Card, error) {
	cardsStr := strings.Split(hand, ",")
	if len(cardsStr) != 5 {
		return nil, errors.New(ErrInvalidHandSize)
	}

	cards := make([]Card, 0, 5)
	for _, cardStr := range cardsStr {
		if len(cardStr) < 2 {
			return nil, errors.New(ErrInvalidCardFormat)
		}

		suit := string(cardStr[0])
		value := cardStr[1:]

		if !isValidSuit(suit) {
			return nil, errors.New(ErrInvalidSuit)
		}

		if !isValidValue(value) {
			return nil, errors.New(ErrInvalidValue)
		}

		cards = append(cards, Card{Suit: suit, Value: value})
	}

	return cards, nil
}

func isValidSuit(suit string) bool {
	switch suit {
	case Heart, Spade, Club, Diamond:
		return true
	default:
		return false
	}
}

func isValidValue(value string) bool {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return false
	}
	return intValue >= MinCardValue && intValue <= MaxCardValue
}
