package main

import (
  "fmt"
)

type Hand struct {
  cards []Card
  tricks int
}

type Player struct {
  hand Hand
  pId int
}

func getInputCard(hand Hand, played []Card, trump int) Card {
  success := false

  var suit int = -1
  var card int = -1

  var inSuit string = ""
  var inCard string = ""

  if len(hand.cards) < 1 {
    return Card{-1,-1}
  }

  for success == false {
    fmt.Scanf("%s of %s\n",&inCard,&inSuit)
    suit = suitToNum(inSuit)
    card = faceToNum(inCard)
    inputCard := Card{suit, card}
    exists := false
    hasSuitLed := false

    if len(played) > 0 {
      for i := 0; i < len(hand.cards); i++ {
        //check if exists in hand
        if hand.cards[i].suit == suit && hand.cards[i].value == card {
          exists = true
        }
        //check for suit led
        if adjustCard(hand.cards[i], trump).suit == adjustCard(played[0], trump).suit {
          hasSuitLed = true
        }
      }


      fmt.Printf("Card chosen: %+v, Card led: %+v\n",adjustCard(inputCard, trump),adjustCard(played[0], trump))
      if exists == true && hasSuitLed == true && adjustCard(inputCard, trump).suit == adjustCard(played[0], trump).suit {
        return inputCard
      } else if exists == true && hasSuitLed == false {
        return inputCard
      } else if exists == false {
        fmt.Printf("Card not in hand, example: ")
        printCard(hand.cards[0])
        fmt.Printf("\n")
      } else if hasSuitLed == true {
        fmt.Printf("Must follow lead suit of: ",)
        printCard(played[0])
        fmt.Printf("\n")
      }

    } else if len(played) == 0 {
      for i := 0; i < len(hand.cards); i++ {
        //check if exists in hand
        if hand.cards[i].suit == suit && hand.cards[i].value == card {
          return inputCard
        }
      }
      fmt.Printf("Card not in hand, example: ")
      printCard(hand.cards[0])
      fmt.Printf("\n")
    }


  }

  return Card{suit, card}
}

func playCard(player *Player, card Card) Card {

  for i := 0; i < len(player.hand.cards); i++ {
    if player.hand.cards[i].suit == card.suit && player.hand.cards[i].value == card.value {
      fmt.Printf("remove %+v\n",card)
      player.hand.cards = append(player.hand.cards[:i], player.hand.cards[i+1:]...)
      return card
    } else {
      fmt.Printf("dont remove %+v\n",player.hand.cards[i])
    }
  }

  return card

}

func getCard(player *Player, card Card) {
  player.hand.cards = append(player.hand.cards, card)
}
