package main

import "fmt"

func main() {
	cards := []string{newCard(), newCard(), "Six of Clubs"}
	cards = append(cards, "Two of Spades")
	fmt.Println(cards)
}

func newCard() string {
	return "Seven of Hearts"
}
