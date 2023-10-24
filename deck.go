package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string // deck is a slice of strings

// Create and return a list of playing cards
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit) // Ace of Spades,Two of Spades,Three of Spades,Four of Spades,Ace of Diamonds,Two of Diamonds,Three of Diamonds,Four of Diamonds,Ace of Hearts,Two of Hearts,Three of Hearts,Four of Hearts,Ace of Clubs,Two of Clubs,Three of Clubs,Four of Clubs
		}
	}

	return cards
}

// Display all the cards in the deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Return two values of type deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[:handSize]
}

// Covert a deck into a string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Save a deck to a file on the local machine
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666) // 0666 is read and write permission
}

// Read a deck from a file on the local machine
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - log the error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ",")) // Ace of Spades,Two of Spades,Three of Spades,Four of Spades,Ace of Diamonds,Two of Diamonds,Three of Diamonds,Four of Diamonds,Ace of Hearts,Two of Hearts,Three of Hearts,Four of Hearts,Ace of Clubs,Two of Clubs,Three of Clubs,Four of Clubs
}

// Shuffle the cards in a deck
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)           // Generate a random number between 0 and 15
		d[i], d[newPosition] = d[newPosition], d[i] // Swap the values
	}
}
