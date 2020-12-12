package main

import (
  //"fmt"
  "math/rand"
  "time"
)


type Deck struct {
  cards []Card
  drawnCards []Card
  inDeckCards []Card
}

type Card struct {
  suit int
  value int
}

const (
  Spade = iota
  Heart
  Diamond
  Club
)

func makeDeck(deckType string) *Deck {
  var deck = new(Deck)

  if deckType == "euchre" {

    for i := 0; i < 4; i++ {
      for j := 0; j < 6; j++ {

        newCard := Card{i, j + 9}
        deck.cards = append(deck.cards, newCard)
        deck.inDeckCards = append(deck.inDeckCards, newCard)

      }
    }
  } else if deckType != "euchre" {

    for i := 0; i < 4; i++ {
      for j := 0; j < 13; j++ {

        newCard := Card{i, j}
        deck.cards = append(deck.cards, newCard)
        deck.inDeckCards = append(deck.inDeckCards, newCard)

      }
    }
  }

  return deck

}

func drawCard(deck *Deck) Card {
  if(len(deck.inDeckCards) == 0) {
    return Card{-1, -1}
  }
  seed := rand.NewSource(time.Now().UnixNano())
  random := rand.New(seed)
  r := random.Intn(len(deck.inDeckCards))
  retCard := deck.inDeckCards[r]
  deck.inDeckCards = append(deck.inDeckCards[:r], deck.inDeckCards[r+1:]...)
  return retCard
}
