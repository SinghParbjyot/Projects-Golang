package main

import (
	"fmt"
	"strings"
)

func main() {
	//cards := newDeckFromFile("my_cards")
	//cards.shuffle()
	//cards.print()

	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, number := range numbers {
		if number%2 == 0 {
			fmt.Println(number, " is even")
		} else {
			fmt.Println(number, " is odd")

		}
	}

}

func newCard() string {
	return "Five of Diamonds"
}

func replaceCard(card string, tipo string) string {
	split := strings.Split(card, " ")
	return split[0] + " " + split[1] + " " + tipo
}
