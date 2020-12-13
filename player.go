package main

type Hand struct {
  cards []Card
  tricks int
}

type Player struct {
  hand Hand
  pId int
}

func playCard(player *Player, card Card) {
  //player.hand.cards.remove(card)
}
