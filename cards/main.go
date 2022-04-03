package main

func main() {
	// create a new deck and print it
	// cards := newDeck()
	// cards.print()


	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()
	// cards.print()

	// save deck to a file
	// cards.saveToFile("cards1.txt")

	// read deck from a file
	// cards := readDeckFromFile("cards1.txt") // test error
	cards := readDeckFromFile("cards1.txt")
	cards.print()
}
