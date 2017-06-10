package blackjack

import (
	"math/rand"
	"time"
)

const (
	BLACKJACK   = 21
	INITIAL_POT = 500
)

var SUITS = []string{"Clubs", "Diamonds", "Hearts", "Spades"}
var RANKS = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King", "Ace"}

type Card struct {
	Suit, Rank string
}

type Deck interface {
	Deal() Card
	Shuffle()
}

type GenericDeck struct {
	cards []Card
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func (deck *GenericDeck) Shuffle() {
	for i := range deck.cards {
		j := rand.Intn(i + 1)
		deck.cards[i], deck.cards[j] = deck.cards[j], deck.cards[i]
	}
}

func (deck *GenericDeck) Deal() Card {
	if len(deck.cards) == 0 {
		return Card{}
	}
	card := deck.cards[0]
	deck.cards = deck.cards[1:]
	return card
}

type BlackJackDeck struct {
	*GenericDeck
}

func NewBlackJackDeck() *BlackJackDeck {
	deck := &BlackJackDeck{&GenericDeck{}}

	for _, suit := range SUITS {
		for _, rank := range RANKS {
			deck.cards = append(deck.cards, Card{suit, rank})
		}
	}

	return deck
}

func test() {
	deck := NewBlackJackDeck()
	deck.Shuffle()
}
