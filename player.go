package main

import (
	"fmt"
)

type Team struct {
	Player1 Player
	Player2 Player
	Points  int
}

type Hand struct {
	Cards []Card
}

type Player struct {
	Hand   Hand
	PId    int
	Tricks int
}

func getInputCard(hand Hand, played []Card, trump int) Card {
	success := false

	var suit int = -1
	var card int = -1

	var inSuit string = ""
	var inCard string = ""

	if len(hand.Cards) < 1 {
		return Card{-1, -1}
	}

	for success == false {
		fmt.Scanf("%s of %s\n", &inCard, &inSuit)
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

			fmt.Printf("Card chosen: %+v, Card led: %+v\n", adjustCard(inputCard, trump), adjustCard(played[0], trump))
			if exists == true && hasSuitLed == true && adjustCard(inputCard, trump).Suit == adjustCard(played[0], trump).Suit {
				return inputCard
			} else if exists == true && hasSuitLed == false {
				return inputCard
			} else if exists == false {
				fmt.Printf("Card not in hand, example: ")
				printCard(hand.Cards[0])
				fmt.Printf("\n")
			} else if hasSuitLed == true {
				fmt.Printf("Must follow lead suit of: ")
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
			fmt.Printf("remove %+v\n", card)
			player.Hand.Cards = append(player.Hand.Cards[:i], player.Hand.Cards[i+1:]...)
			return card
		} else {
			fmt.Printf("dont remove %+v\n", player.Hand.Cards[i])
		}
	}

	return card

}

func validateCardInHand(hand Hand, card Card) bool {
	for i := 0; i < len(hand.Cards); i++ {
		if hand.Cards[i].Suit == card.Suit && hand.Cards[i].Value == card.Value {
			return true
		}
	}
	return false
}

func validateFollowSuitRules(hand Hand, card Card, lead Card, trump int) bool {
	if card.Suit == lead.Suit {
		return true
	} else {
		for i := 0; i < len(hand.Cards); i++ {
			if hand.Cards[i].Suit == lead.Suit {
				return false
			}
		}
		return true
	}
}

func getCard(player *Player, card Card) {
	player.Hand.Cards = append(player.Hand.Cards, card)
}

func swapWithCardInHand(player *Player, card Card, swap Card) bool {
	for i := 0; i < len(player.Hand.Cards); i++ {
		if player.Hand.Cards[i].Suit == card.Suit && player.Hand.Cards[i].Value == card.Value {
			player.Hand.Cards[i] = swap
			return true
		}
	}
	return false
}

func returnPlayerOnPId(players []*Player, pid int) *Player {
	for i := 0; i < len(players); i++ {
		if players[i].PId == pid {
			return players[i]
		}
	}
	return nil
}

func exampleUseofReturnPlayerOnPId() {
	players := []*Player{&Player{PId: 1}, &Player{PId: 2}, &Player{PId: 3}, &Player{PId: 4}}
	player := returnPlayerOnPId(players, 2)
	fmt.Printf("Player: %+v\n", player)
}
