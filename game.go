package main

import (
  "fmt"
)

//TODO go around letting people pass or play
//TODO choose trump after turned down
//TODO change start order after hands

//TODO Re-factor everything its a huge mess

//TODO create reactjs frontend - will need to remember how to do this


//Using a slice instead of list here. List is probably better but for familiarity slice should work
type gameState struct {
  Order []Player
}

type roundState struct {
  Start Player
}


func winRound(cards []Card, trump int) int {

  var numTrump int = 0
  var largestTrump int = -1
  var largestOfLed int = -1
  var memCards [4]Card

  led := cards[0].Suit


  //adjusting for jacks
  for i := 0; i < 4; i++ {
    cards[i] = adjustCard(cards[i], trump)

    if cards[i].Suit == trump {
      numTrump++
    }
    memCards[i] = cards[i]
  }


  fmt.Printf("adjusted cards: %+v\n",memCards)


  //until 1 cards is determined the largest
  for len(cards) > 1 {

    //remove all cards that don't matter
    for i := 0; i < len(cards); i++ {

      if cards[i].Suit != trump && numTrump > 0 {
        //fmt.Printf("%+v not trump when trump was played\n",cards[i])
        cards = append(cards[:i], cards[i+1:]...)
        i = 0
      } else if cards[i].Suit != led && numTrump == 0 {
        //fmt.Printf("%+v not following suit when no trump was played\n",cards[i])
        cards = append(cards[:i], cards[i+1:]...)
        i = 0
      }

    }

    //find which of remaining trump is largest
    if numTrump == len(cards) {

      //fmt.Printf("\n%+v\n\n",cards)

      for i := 0; i < len(cards); i++ {
        //fmt.Printf("%d > %d\n",cards[i].value, largestTrump)
        if cards[i].Value > largestTrump {
          //fmt.Printf("%+v currently largest trump\n",cards[i])
          largestTrump = cards[i].Value
        } else if cards[i].Value < largestTrump {
          //fmt.Printf("%+v trump but smaller\n",cards[i])
          cards = append(cards[:i], cards[i+1:]...)
          i = -1
        }
        //fmt.Printf("\n%+v\n\n",cards)
      }

      //else do largest of led suit
    } else {
      for i := 0; i < len(cards); i++ {
        if cards[i].Value > largestOfLed {
          //fmt.Printf("%+v currently largest led\n",cards[i])
          largestOfLed = cards[i].Value
        } else if cards[i].Value < largestOfLed {
          //fmt.Printf("%+v led but smaller\n",cards[i])
          cards = append(cards[:i], cards[i+1:]...)
          i = -1
        }
        //fmt.Printf("\n%+v\n\n",cards)
      }
    }

  }

  //fmt.Printf("memCards: %+v\n",memCards)
  for i := 0; i < 4; i++ {
    if memCards[i] == cards[0] {
      printCard(memCards[i])
      fmt.Printf(" won the trick!\n")
      return i
    }
  }

  //fmt.Printf("\n%+v\n\n",cards)
  return -1
}

func test() {
  var suit int = -1
  var card int = -1
  var cardsPlayed []Card

  for i := 0; i < 4; i++ {
    fmt.Printf("Card (suit,card):\n")
    fmt.Scanf("%d,%d", &suit, &card)
    //cardsPlayed[i] = Card{suit, card}
    cardsPlayed = append(cardsPlayed, Card{suit, card})
  }

  fmt.Printf("Result: %d\n\n",winRound(cardsPlayed, 0))



}

func playRound(team1 *Team, team2 *Team, gS gameState, trump int) Player {

  var cardsPlayed []Card

  for i := 0; i < 4; i++ {
    if team1.Player1.PId == gS.Order[i].PId {
      fmt.Printf("\nTeam 1 Player 1:\n")
      printHand(team1.Player1.Hand)
      fmt.Printf("Choose a card to play: ")
      removeCard := getInputCard(team1.Player1.Hand, cardsPlayed, trump)
      cardsPlayed = append(cardsPlayed, playCard(&team1.Player1, removeCard))

    } else if team1.Player2.PId == gS.Order[i].PId {
      fmt.Printf("\nTeam 1 Player 2:\n")
      printHand(team1.Player2.Hand)
      fmt.Printf("Choose a card to play: ")
      removeCard := getInputCard(team1.Player2.Hand, cardsPlayed, trump)
      cardsPlayed = append(cardsPlayed, playCard(&team1.Player2, removeCard))

    } else if team2.Player1.PId == gS.Order[i].PId {
      fmt.Printf("\nTeam 2 Player 1:\n")
      printHand(team2.Player1.Hand)
      fmt.Printf("Choose a card to play: ")
      removeCard := getInputCard(team2.Player1.Hand, cardsPlayed, trump)
      cardsPlayed = append(cardsPlayed, playCard(&team2.Player1, removeCard))

    } else if team2.Player2.PId == gS.Order[i].PId {
      fmt.Printf("\nTeam 2 Player 2:\n")
      printHand(team2.Player2.Hand)
      fmt.Printf("Choose a card to play: ")
      removeCard := getInputCard(team2.Player2.Hand, cardsPlayed, trump)
      cardsPlayed = append(cardsPlayed, playCard(&team2.Player2, removeCard))
    }
  }
  t := winRound(cardsPlayed, trump)

  if team1.Player1.PId == gS.Order[t].PId {
    team1.Player1.Hand.Tricks++
    fmt.Printf("\nTeam 1:\n")
    printTeam(*team1)
    fmt.Printf("\nTeam 2:\n")
    printTeam(*team2)
    return team1.Player1
  } else if team1.Player2.PId == gS.Order[t].PId {
    team1.Player2.Hand.Tricks++
    fmt.Printf("\nTeam 1:\n")
    printTeam(*team1)
    fmt.Printf("\nTeam 2:\n")
    printTeam(*team2)
    return team1.Player2
  } else if team2.Player1.PId == gS.Order[t].PId {
    team2.Player1.Hand.Tricks++
    fmt.Printf("\nTeam 1:\n")
    printTeam(*team1)
    fmt.Printf("\nTeam 2:\n")
    printTeam(*team2)
    return team2.Player1
  } else if team2.Player2.PId == gS.Order[t].PId {
    team2.Player2.Hand.Tricks++
    fmt.Printf("\nTeam 1:\n")
    printTeam(*team1)
    fmt.Printf("\nTeam 2:\n")
    printTeam(*team2)
    return team2.Player2
  }

  return team1.Player1
}

func playHand(team1 *Team, team2 *Team, deck *Deck, gS gameState, firstPlayer Player) {

  var trump int = -1
  trump = 0

  turnUp := drawCard(deck)
  //fmt.Printf("Turned up card is: %+v\n\n",turnUp)
  fmt.Printf("Turned up card is: ")
  printCard(turnUp)
  fmt.Printf("\n\n")

  var emptyCards []Card

  if team1.Player1.PId == firstPlayer.PId {
    fmt.Printf("Replace card from(t1p1): ")
    printHand(team1.Player1.Hand)
    removeCard := getInputCard(team1.Player1.Hand, emptyCards, -1)
    playCard(&team1.Player1, removeCard)
    getCard(&team1.Player1, turnUp)
    trump = turnUp.Suit

  } else if team1.Player2.PId == firstPlayer.PId {
    fmt.Printf("Replace card from(t1p2): ")
    printHand(team1.Player2.Hand)
    removeCard := getInputCard(team1.Player2.Hand, emptyCards, -1)
    playCard(&team1.Player2, removeCard)
    getCard(&team1.Player2, turnUp)
    trump = turnUp.Suit

  } else if team2.Player1.PId == firstPlayer.PId {
    fmt.Printf("Replace card from(t2p1): ")
    printHand(team2.Player1.Hand)
    removeCard := getInputCard(team2.Player1.Hand, emptyCards, -1)
    playCard(&team2.Player1, removeCard)
    getCard(&team2.Player1, turnUp)
    trump = turnUp.Suit

  } else if team2.Player2.PId == firstPlayer.PId {
    fmt.Printf("Replace card from(t2p2): ")
    printHand(team2.Player2.Hand)
    removeCard := getInputCard(team2.Player2.Hand, emptyCards, -99)
    playCard(&team2.Player2, removeCard)
    getCard(&team2.Player2, turnUp)
    trump = turnUp.Suit
  }


  for i := 0; i < 5; i++ {
    firstPlayerFound := false

    //rotate through list to find first player
    for firstPlayerFound == false {
      if gS.Order[0].PId == firstPlayer.PId {
        firstPlayerFound = true
      } else {
        x := gS.Order[0]
        gS.Order = gS.Order[1:]
        gS.Order = append(gS.Order, x)
      }
    }


    firstPlayer = playRound(team1, team2, gS, trump)
  }

}

func play() {
  playerA := new(Player)
  playerB := new(Player)
  playerC := new(Player)
  playerD := new(Player)

  playerA.PId = 0
  playerB.PId = 1
  playerC.PId = 2
  playerD.PId = 3


  team1 := Team{Player1: *playerA, Player2: *playerB, Points: 0}
  team2 := Team{Player1: *playerC, Player2: *playerD, Points: 0}

  gState := gameState{Order: []Player{team1.Player1, team2.Player1, team1.Player2, team2.Player2}}

  deck := *makeDeck("euchre")

  //fmt.Printf("%+v\n",deck)
  printDeck(deck)

  drawHand(&deck, &team1.Player1.Hand, 5)
  drawHand(&deck, &team1.Player2.Hand, 5)
  drawHand(&deck, &team2.Player1.Hand, 5)
  drawHand(&deck, &team2.Player2.Hand, 5)


  //fmt.Printf("\nplayer1 from team1:\n%+v\n",team1.Player1)
  //fmt.Printf("player2 from team1:\n%+v\n\n",team1.Player2)
  //fmt.Printf("player1 from team2:\n%+v\n\n",team2.Player1)
  //fmt.Printf("player2 from team2:\n%+v\n\n",team2.Player2)
  printTeam(team1)
  printTeam(team2)

  //fmt.Printf("deck:\n%+v\n\n",deck)
  printDeck(deck)

  playHand(&team1, &team2, &deck, gState, team1.Player1)

}

func main() {
  fmt.Println("hello world")

  var input string = ""

  deck := *makeDeck("euchre")

  for input != "exit" {
    fmt.Println("Enter a command: ")
    fmt.Scan(&input)

    if input == "help" {
      fmt.Println("Commands:\n play <card>\n help\n exit")
    } else if input == "exit" {
      fmt.Println("Exiting...\n")
    } else if input == "play" {
      play()
    } else if input == "test" {
      test()
    } else {
      fmt.Println("Invalid command\n")
    }

    newCard := drawCard(&deck)
    fmt.Println(newCard)
    fmt.Println(deck.InDeckCards)

  }



  //
  // fmt.Println(deck)
  //
  // for i := 0; i < 20; i++ {
  //   newCard := drawCard(&deck)
  //   fmt.Println(newCard)
  //
  // }
  //
  // //newCard := drawCard(&deck)
  //
  // //fmt.Println(newCard)
  // fmt.Println(deck)
}
