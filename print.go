package main

import (
	"fmt"
)

func printTeam(team Team) {

	printPlayer(team.Player1)
	printPlayer(team.Player2)
	fmt.Printf("Points: %d\n\n", team.Points)

}

func printHand(hand Hand) {

	for i := 0; i < len(hand.Cards); i++ {
		printCard(hand.Cards[i])
		fmt.Printf(", ")
	}
	fmt.Printf("\n")

}

func printPlayer(player Player) {

	fmt.Printf("Player %d has the hand: ", player.PId)
	fmt.Printf("Tricks: %d\n", player.Tricks)
	printHand(player.Hand)

}

func printDeck(deck Deck) {

	fmt.Printf("Deck has: ")
	for i := 0; i < len(deck.InDeckCards); i++ {
		printCard(deck.InDeckCards[i])
		fmt.Printf(", ")
	}
	fmt.Printf("\n\n")

}

func printCard(card Card) {

	tempValue := numToFace(card.Value)

	if card.Suit == 0 {
		fmt.Printf("%s of Spades", tempValue)
	} else if card.Suit == 1 {
		fmt.Printf("%s of Hearts", tempValue)
	} else if card.Suit == 2 {
		fmt.Printf("%s of Clubs", tempValue)
	} else if card.Suit == 3 {
		fmt.Printf("%s of Diamonds", tempValue)
	}

}
