package main

import (
	"os"
	"testing"
)

func TestNewDec(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Clubs" {
		t.Errorf("Expected first card to be Ace of Clubs, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Spades" {
		t.Errorf("Expected last card to be Four of Spades, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	const testFilename string = "_deckTesting.txt"

	os.Remove(testFilename)

	d := newDeck()
	err := d.saveToFile(testFilename)

	if err != nil {
		t.Error(err)
	}

	loadedDeck := newDeckFromFile(testFilename)

	if len(d) != len(loadedDeck) {
		t.Errorf("Expected saved length %v to be the same as loaded length %v", len(d), len(loadedDeck))
	}

	for i, card := range d {
		if card != loadedDeck[i] {
			t.Errorf("Expected saved %v to be the same as loaded %v", card, loadedDeck[i])
			t.FailNow()
		}
	}
}
