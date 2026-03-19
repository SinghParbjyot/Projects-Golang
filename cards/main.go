package main

import (
	"strings"
)

func main() {
	cards := newDeck()
	hand, reamainingCards := deal(cards, 5)

	hand.print()
	reamainingCards.print()
}

func newCard() string {
	return "Five of Diamonds"
}

func replaceCard(card string, tipo string) string {
	split := strings.Split(card, " ")
	return split[0] + " " + split[1] + " " + tipo
}
