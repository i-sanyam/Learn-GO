package main

func main() {
	// cards := newDeck()
	// cards.print()
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()
	// cards.print()
	// cards.saveToFile("cards1.txt")
	cards := readDeckFromFile("cards1.txt")
	cards.print()
}
