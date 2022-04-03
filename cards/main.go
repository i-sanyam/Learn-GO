package main

import "fmt"

func main() {
	cards := []string{newCard(), newCard(), "Six of Clubs"}
	cards = append(cards, "Two of Spades")
	
	for i, card := range cards {
		fmt.Println(i, card);
	}
}

func newCard() string {
	return "Seven of Hearts"
}
