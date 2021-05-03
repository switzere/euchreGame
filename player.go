package main

import (
  "fmt"
)

type Team struct {
  Player1 Player
  Player2 Player
  Points int
}

type Hand struct {
  Cards []Card
  Tricks int
}

type Player struct {
  Hand Hand
  PId int
}

func getInputCard(hand Hand, played []Card, trump int) Card {
  success := false

  var suit int = -1
  var card int = -1

  var inSuit string = ""
  var inCard string = ""

  if len(hand.Cards) < 1 {
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
      for i := 0; i < len(hand.Cards); i++ {
        //check if exists in hand
        if hand.Cards[i].Suit == suit && hand.Cards[i].Value == card {
          exists = true
        }
        //check for suit led
        if adjustCard(hand.Cards[i], trump).Suit == adjustCard(played[0], trump).Suit {
          hasSuitLed = true
        }
      }


      fmt.Printf("Card chosen: %+v, Card led: %+v\n",adjustCard(inputCard, trump),adjustCard(played[0], trump))
      if exists == true && hasSuitLed == true && adjustCard(inputCard, trump).Suit == adjustCard(played[0], trump).Suit {
        return inputCard
      } else if exists == true && hasSuitLed == false {
        return inputCard
      } else if exists == false {
        fmt.Printf("Card not in hand, example: ")
        printCard(hand.Cards[0])
        fmt.Printf("\n")
      } else if hasSuitLed == true {
        fmt.Printf("Must follow lead suit of: ",)
        printCard(played[0])
        fmt.Printf("\n")
      }

    } else if len(played) == 0 {
      for i := 0; i < len(hand.Cards); i++ {
        //check if exists in hand
        if hand.Cards[i].Suit == suit && hand.Cards[i].Value == card {
          return inputCard
        }
      }
      fmt.Printf("Card not in hand, example: ")
      printCard(hand.Cards[0])
      fmt.Printf("\n")
    }


  }

  return Card{suit, card}
}

func playCard(player *Player, card Card) Card {

  for i := 0; i < len(player.Hand.Cards); i++ {
    if player.Hand.Cards[i].Suit == card.Suit && player.Hand.Cards[i].Value == card.Value {
      fmt.Printf("remove %+v\n",card)
      player.Hand.Cards = append(player.Hand.Cards[:i], player.Hand.Cards[i+1:]...)
      return card
    } else {
      fmt.Printf("dont remove %+v\n",player.Hand.Cards[i])
    }
  }

  return card

}

func getCard(player *Player, card Card) {
  player.Hand.Cards = append(player.Hand.Cards, card)
}
