package gamestate

import "github.com/wwi22ama-prog/cardgames/cards"

type GameState struct {
	Players       []cards.Hand
	Deck          cards.Deck
	UsedDeck      cards.Deck
	CurrentPlayer int
}
