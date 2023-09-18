package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of 'deck' which is a slice of strings

type deck []string // create new type deck with type keyword.
// It is in type slice string -> []string which means new type uses behavior of class slice string (extends) barrows the behavior.
// deck yerine []string yazdigimiz yerleri degistirebiliriz. cunku deck = []string same thing oldu.

func (d deck) print() { // receiver kullandik yani (d deck). Bu sayede .print ile erisebiliyoruz. Ne zaman kullandigimizi aciklayacak.
	// func () this paranthesis is called a receiver.
	// d deck mean s any variable of type deck can access to the print method
	// reference to the actual deck is 'd'
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck { // deck is a return type.
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	// This function will create cards from cardSuit and cardValues slices and add this card to cards slice.
	// We will have all combinations

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) { // function call deck.deal deal fonksiyonunu bunlar disinda bir seyle cagirmamiza izin vermeyecek
	//deck tipinde iki referans donecek (deck, deck)
	return d[:handSize], d[handSize:]
	// startindexincluded:endIndexNotIncluded
}

func (d deck) toString() string { // d deck belirttigimiz icin ilerde cards.toString ile cagiracagiz.
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) suffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
