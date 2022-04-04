package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 12 {
		t.Errorf("Expected a deck length of 20, but got %v", len(d))
	}
	if d[0] != "Ace of Clubs" {
		t.Errorf("Expected Ace of Clubs, but got %v", d[0])
	}
	if d[len(d)-1] != "Three of Spades" {
		t.Errorf("Expected Three of Spades, but got %v", d[0])
	}
}
