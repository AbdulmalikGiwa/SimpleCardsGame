package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	testDeck := newDeck()

	if len(testDeck) != 52 {
		t.Errorf("Expected length of deck is 52, but got %v", len(testDeck))
	}

	if testDeck[0] != "Ace of Spades" {
		t.Errorf("Expected first card Ace of Spades, but got %v", testDeck[0])
	}

	if testDeck[len(testDeck)-1] != "Ten of Diamonds" {
		t.Errorf("Expected last card Ten of Diamonds, but got %v", testDeck[len(testDeck)-1])
	}

}
func TestSaveToFileAndNewFileFromDeck(t *testing.T) {
	os.Remove("_deckTesting")
	d := newDeck()
	d.saveToFile("_deckTesting")

	deckFile := newDeckFromFile("_deckTesting")

	if len(deckFile) != 52 {
		t.Errorf("Deck not loaded completely, 52 expected but got %v ", len(deckFile))
	}

}
