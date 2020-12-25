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
