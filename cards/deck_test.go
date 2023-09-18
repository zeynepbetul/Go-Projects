package main

import (
	"os"
	"testing"
)

// Part 36.Writing Useful Tests

//8.00 -> t.Errorf("Expected deck length of 16, but got %v", len(d))

// we will create 2 test functions.
// First one will do 3 operations:  make sure that a deck is created with 4 number of cards, first card is ace of spades, last card is a four of clubs
// Second one will do an operation of save deck and new deck from file

func TestNumberofCards(t *testing.T) { // upper case because it is a test function // will be automatically called by test runner (t is test handler)
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck lenght of 16, but we got %v ", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but we got %v", d[0])
	}
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of clubs, but we got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) { // name of the test func includes which functions it tests and starts with capital Test
	// so we can easily find test function of function we are testing.

	os.Remove("_decktesting") // remove old test files created

	d := newDeck()

	d.saveToFile("_decktesting") // need to delete test files manually. Go does not delete automatically.
	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
