package main

// dynamic type languages javascript, ruby, python deveoopers dont care which type our varible is
// static type languages java, go, c++ type should be defined.
func main() {
	// var card string = "Ace of Spades" // to initialize a variable this line can also be used: card:="Ace of Spades". Go determines the
	// type itself.
	// cards := newDeck() // string card
	// cards := []string{"Ace of diamonds", newCard()} // create string slice now I wil change this []string deck.
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")
	// for i, card := range cards { // in each iteration it declares variables i and card from cards and with the use of range keyword.
	// fmt.Println(i, card) }
	// cards.print() // cards can access function print beacause it is in type deck.
	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()
	// 27. bolum joining slice of strings
	// part 28. Save cards to a file cards.saveToFile("newFile")
	// part 29. Read cards from the file
	cards := newDeckFromFile("newFile")
	cards.print()
}

// in go array is fixed length list, slice can grow.
// in slice every element must be of same type.
// Go is not oop.

// 26. Type conversion -> []byte("Hi there!") -> first one type we want. Second one value we have. Stringi byte fonksiyonuna almisin gibi.
// bizim deckimiz string slice tipinde zaten. []string -> tek bir string haline getirecegiz -> []byte
