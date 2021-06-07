package main

import (
  //"fmt"
  "math/rand"
  "time"
  "strings"
  "strconv"
)


type Deck struct {
  Cards []Card
  DrawnCards []Card
  InDeckCards []Card
}

type Card struct {
  Suit int
  Value int
}

type playedCard struct {
  Suit int
  Value int
  Owner Player
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
        deck.Cards = append(deck.Cards, newCard)
        deck.InDeckCards = append(deck.InDeckCards, newCard)

      }
    }
  } else if deckType != "euchre" {

    for i := 0; i < 4; i++ {
      for j := 0; j < 13; j++ {

        newCard := Card{i, j}
        deck.Cards = append(deck.Cards, newCard)
        deck.InDeckCards = append(deck.InDeckCards, newCard)

      }
    }
  }

  return deck

}

func drawCard(deck *Deck) Card {
  if(len(deck.InDeckCards) == 0) {
    return Card{-1, -1}
  }
  seed := rand.NewSource(time.Now().UnixNano())
  random := rand.New(seed)
  r := random.Intn(len(deck.InDeckCards))
  retCard := deck.InDeckCards[r]
  deck.InDeckCards = append(deck.InDeckCards[:r], deck.InDeckCards[r+1:]...)
  return retCard
}

func drawSpecificCard(deck *Deck, card Card) Card {

  for i := 0; i < len(deck.InDeckCards); i++ {
    if deck.InDeckCards[i].Suit == card.Suit && deck.InDeckCards[i].Value == card.Value {
      deck.InDeckCards = append(deck.InDeckCards[:i], deck.InDeckCards[i+1:]...)
      return card
    }
  }

  return Card{-1, -1}

}

func drawHand(deck *Deck, hand *Hand, numCards int) {
  if numCards <= 0 {
    return
  }

  hand.Cards = nil

  for i := 0; i < numCards; i++ {
    hand.Cards = append(hand.Cards, drawCard(deck))
  }
}

func drawHandUnique(deck *Deck, player *Player, numCards int) {
  drawHand(deck, &player.Hand, numCards)
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
    if card.Suit == Diamond && card.Value == 11 {
      card.Value = 15
      card.Suit = Heart
    } else if card.Suit == Heart && card.Value == 11 {
      card.Value = 16
    }
  } else if trump == Diamond {
    if card.Suit == Heart && card.Value == 11 {
      card.Value = 15
      card.Suit = Diamond
    } else if card.Suit == Diamond && card.Value == 11 {
      card.Value = 16
    }
  } else if trump == Spade {
    if card.Suit == Club && card.Value == 11 {
      card.Value = 15
      card.Suit = Spade
    } else if card.Suit == Spade && card.Value == 11 {
      card.Value = 16
    }
  } else if trump == Club {
    if card.Suit == Spade && card.Value == 11 {
      card.Value = 15
      card.Suit = Club
    } else if card.Suit == Club && card.Value == 11 {
      card.Value = 16
    }
  }

  return card

}
