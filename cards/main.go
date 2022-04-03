package main

func main() {
	cards := deck{newCard(), newCard(), "Six of Clubs"}
	cards = append(cards, "Two of Spades")

	cards.print()
}

func newCard() string {
	return "Seven of Hearts"
}
