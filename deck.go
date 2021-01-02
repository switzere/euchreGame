package main

import (
  //"fmt"
  "math/rand"
  "time"
  "strings"
  "strconv"
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

type playedCard struct {
  suit int
  value int
  owner Player
}

const (
  Spade = iota
  Heart
  Club
  Diamond
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

func drawHand(deck *Deck, hand *Hand, numCards int) {
  if numCards <= 0 {
    return
  }

  for i := 0; i < numCards; i++ {
    hand.cards = append(hand.cards, drawCard(deck))
  }
}

func suitToNum(suit string) int {

  if strings.ToLower(suit) == "spade" || strings.ToLower(suit) == "spades" {
    return 0
  } else if strings.ToLower(suit) == "heart" || strings.ToLower(suit) == "hearts" {
    return 1
  } else if strings.ToLower(suit) == "club" || strings.ToLower(suit) == "clubs" {
    return 2
  } else if strings.ToLower(suit) == "diamond" || strings.ToLower(suit) == "diamonds" {
    return 3
  }

  return -1
}

func numToSuit(suit int) string {

  if suit == 0 {
    return "Spade"
  } else if suit == 1 {
    return "Heart"
  } else if suit == 2 {
    return "Club"
  } else if suit == 3 {
    return "Diamond"
  }

  return "error"
}

func faceToNum(value string) int {

  if strings.ToLower(value) == "jack" {
    return 11
  } else if strings.ToLower(value) == "queen" {
    return 12
  } else if strings.ToLower(value) == "king" {
    return 13
  } else if strings.ToLower(value) == "ace" {
    return 14
  }

  ret, _ := strconv.Atoi(value)
  return ret
}

func numToFace(value int) string {

  if value == 11 {
    return "Jack"
  } else if value == 12 {
    return "Queen"
  } else if value == 13 {
    return "King"
  } else if value == 14 {
    return "Ace"
  }

  return  strconv.Itoa(value)
}

func adjustCard(card Card, trump int) Card {

  if trump == Heart {
    if card.suit == Diamond && card.value == 11 {
      card.value = 15
      card.suit = Heart
    } else if card.suit == Heart && card.value == 11 {
      card.value = 16
    }
  } else if trump == Diamond {
    if card.suit == Heart && card.value == 11 {
      card.value = 15
      card.suit = Diamond
    } else if card.suit == Diamond && card.value == 11 {
      card.value = 16
    }
  } else if trump == Spade {
    if card.suit == Club && card.value == 11 {
      card.value = 15
      card.suit = Spade
    } else if card.suit == Spade && card.value == 11 {
      card.value = 16
    }
  } else if trump == Club {
    if card.suit == Spade && card.value == 11 {
      card.value = 15
      card.suit = Club
    } else if card.suit == Club && card.value == 11 {
      card.value = 16
    }
  }

  return card

}
