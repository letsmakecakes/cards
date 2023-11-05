# Go Playing Cards Deck

This Go program defines a simple deck of playing cards and provides various operations on it. You can create, manipulate, and save decks of cards, making it a useful tool for card games and simulations.

## Features

- **Create a New Deck**: Generate a deck of playing cards containing all the standard cards, such as Ace of Spades, Two of Diamonds, Three of Hearts, etc.

- **Display Cards**: Display all the cards in the deck.

- **Deal Cards**: Split the deck into two decks, simulating a card deal.

- **Convert to String**: Convert the deck into a comma-separated string.

- **Save to File**: Save the deck to a file on your local machine.

- **Load from File**: Read a deck from a file on your local machine.

- **Shuffle Cards**: Shuffle the cards in the deck, ensuring randomness.

## Installation

1. Clone this repository to your Go workspace:

   ```bash
   git clone https://github.com/yourusername/cards.git
   ```

2. Navigate to the project directory:

   ```bash
   cd goplayingcardsdeck
   ```

3. Build and run the application:

   ```bash
   go build
   ./goplayingcardsdeck
   ```

## Usage

Here's an example of how to use the Go Playing Cards Deck in your Go program:

```go
package main

import "fmt"

func main() {
    // Create a new deck
    cards := newDeck()

    // Shuffle the cards
    cards.shuffle()

    // Deal a hand of 5 cards
    hand, remainingDeck := deal(cards, 5)

    // Display the hand and the remaining cards
    fmt.Println("Hand:")
    hand.print()
    fmt.Println("Remaining Cards:")
    remainingDeck.print()

    // Save the remaining deck to a file
    remainingDeck.saveToFile("mydeck.txt")

    // Load a deck from a file
    loadedDeck := newDeckFromFile("mydeck.txt")

    // Display the loaded deck
    fmt.Println("Loaded Deck:")
    loadedDeck.print()
}
```
