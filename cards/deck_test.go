package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 12 {
		t.Errorf("Expected a deck length of 12, but got %v", len(d))
	}
	if d[0] != "Ace of Clubs" {
		t.Errorf("Expected Ace of Clubs, but got %v", d[0])
	}
	if d[len(d)-1] != "Three of Spades" {
		t.Errorf("Expected Three of Spades, but got %v", d[0])
	}
}

func TestSaveToFileAndreadDeckFromFile(t *testing.T) {
	os.Remove("_deck_testing_2231")

	deck := newDeck()
	deck.saveToFile("_deck_testing_2231")

	loadedDeck := readDeckFromFile("_deck_testing_2231")

	if len(loadedDeck) != len(deck) {
		t.Errorf("Expected a deck length of %v, but got %v", len(deck), len(loadedDeck))
	}

	for i := range deck {
		if deck[i] != loadedDeck[i] {
			t.Errorf("Cards do not match at index: %v, Expected Card Value of %v, but got %v", i, deck[i], loadedDeck[i])
		}
	}

	os.Remove("_deck_testing_2231")
}
