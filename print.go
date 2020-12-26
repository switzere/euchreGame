package main

import (
  "fmt"
)

func printTeam(team Team) {

  printPlayer(team.player1)
  printPlayer(team.player2)
  fmt.Printf("Points: %d\n\n",team.points)

}

func printHand(hand Hand) {

  for i := 0; i < len(hand.cards); i++ {
    printCard(hand.cards[i])
    fmt.Printf(", ")
  }
  fmt.Printf("\n")

  fmt.Printf("Tricks: %d\n", hand.tricks)

}

func printPlayer(player Player) {

  fmt.Printf("Player %d has the hand: ", player.pId)
  printHand(player.hand)

}

func printDeck(deck Deck) {

  fmt.Printf("Deck has: ")
  for i := 0; i < len(deck.inDeckCards); i++ {
    printCard(deck.inDeckCards[i])
    fmt.Printf(", ")
  }
  fmt.Printf("\n\n")

}

func printCard(card Card) {

  tempValue := numToFace(card.value)

  if card.suit == 0 {
    fmt.Printf("%s of Spades", tempValue)
  } else if card.suit == 1 {
    fmt.Printf("%s of Hearts", tempValue)
  } else if card.suit == 2 {
    fmt.Printf("%s of Clubs", tempValue)
  } else if card.suit == 3 {
    fmt.Printf("%s of Diamonds", tempValue)
  }

}
