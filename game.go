package main

import (
  "fmt"
)

type Team struct {
  player1 Player
  player2 Player
  points int
}

//Using a slice instead of list here. List is probably better but for familiarity slice should work
type gameState struct {
  order []Player
}

type roundState struct {
  start Player
}

func winRound(cards [4]Card) int {
//temporary calculator
  if cards[0].value > cards[1].value && cards[0].value > cards[2].value && cards[0].value > cards[3].value {
    return 0
  } else if cards[1].value > cards[0].value && cards[1].value > cards[2].value && cards[1].value > cards[3].value {
    return 1
  } else if cards[2].value > cards[0].value && cards[2].value > cards[1].value && cards[2].value > cards[3].value {
    return 2
  } else if cards[3].value > cards[0].value && cards[3].value > cards[1].value && cards[3].value > cards[2].value {
    return 3
  }

  return -1
}

func playRound(team1 *Team, team2 *Team, gS gameState, trump int) Player {

  var suit int = -1
  var card int = -1

  var cardsPlayed [4]Card

  for i := 0; i < 4; i++ {
    if team1.player1.pId == gS.order[i].pId {
      fmt.Printf("\nTeam 1 Player 1:\n")
      fmt.Printf("Hand: %+v\n\n",team1.player1.hand)
      fmt.Printf("Choose a card to play: ")
      fmt.Scanf("%d,%d", &suit, &card)
      fmt.Println(suit)
      fmt.Println(card)
      removeCard := Card{suit, card}
      cardsPlayed[i] = playCard(&team1.player1, removeCard)
    } else if team1.player2.pId == gS.order[i].pId {
      fmt.Printf("\nTeam 1 Player 2:\n")
      fmt.Printf("Hand: %+v\n\n",team1.player2.hand)
      fmt.Printf("Choose a card to play: ")
      fmt.Scanf("%d,%d", &suit, &card)
      fmt.Println(suit)
      fmt.Println(card)
      removeCard := Card{suit, card}
      cardsPlayed[i] = playCard(&team1.player2, removeCard)
    } else if team2.player1.pId == gS.order[i].pId {
      fmt.Printf("\nTeam 2 Player 1:\n")
      fmt.Printf("Hand: %+v\n\n",team2.player1.hand)
      fmt.Printf("Choose a card to play: ")
      fmt.Scanf("%d,%d", &suit, &card)
      fmt.Println(suit)
      fmt.Println(card)
      removeCard := Card{suit, card}
      cardsPlayed[i] = playCard(&team2.player1, removeCard)
    } else if team2.player2.pId == gS.order[i].pId {
      fmt.Printf("\nTeam 2 Player 2:\n")
      fmt.Printf("Hand: %+v\n\n",team2.player2.hand)
      fmt.Printf("Choose a card to play: ")
      fmt.Scanf("%d,%d", &suit, &card)
      fmt.Println(suit)
      fmt.Println(card)
      removeCard := Card{suit, card}
      cardsPlayed[i] = playCard(&team2.player2, removeCard)
    }
  }
  t := winRound(cardsPlayed)

  if team1.player1.pId == gS.order[t].pId {
    team1.player1.hand.tricks++
    fmt.Printf("\nteam1:\n%+v\n",team1)
    fmt.Printf("\nteam2:\n%+v\n",team2)
    return team1.player1
  } else if team1.player2.pId == gS.order[t].pId {
    team1.player2.hand.tricks++
    fmt.Printf("\nteam1:\n%+v\n",team1)
    fmt.Printf("\nteam2:\n%+v\n",team2)
    return team1.player2
  } else if team2.player1.pId == gS.order[t].pId {
    team2.player1.hand.tricks++
    fmt.Printf("\nteam1:\n%+v\n",team1)
    fmt.Printf("\nteam2:\n%+v\n",team2)
    return team2.player1
  } else if team2.player2.pId == gS.order[t].pId {
    team2.player2.hand.tricks++
    fmt.Printf("\nteam1:\n%+v\n",team1)
    fmt.Printf("\nteam2:\n%+v\n",team2)
    return team2.player2
  }

  return team1.player1
}

func playHand(team1 *Team, team2 *Team, gS gameState, firstPlayer Player) {

  var trump int = -1
  trump = 0

  for i := 0; i < 5; i++ {
    firstPlayerFound := false

    for firstPlayerFound == false {
      if gS.order[0].pId == firstPlayer.pId {
        firstPlayerFound = true
      } else {
        x := gS.order[0]
        gS.order = gS.order[1:]
        gS.order = append(gS.order, x)
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

  playerA.pId = 0
  playerB.pId = 1
  playerC.pId = 2
  playerD.pId = 3


  team1 := Team{player1: *playerA, player2: *playerB, points: 0}
  team2 := Team{player1: *playerC, player2: *playerD, points: 0}

  gState := gameState{order: []Player{team1.player1, team1.player2, team2.player1, team2.player2}}

  deck := *makeDeck("euchre")

  fmt.Printf("%+v\n",deck)

  drawHand(&deck, &team1.player1.hand, 5)
  drawHand(&deck, &team1.player2.hand, 5)
  drawHand(&deck, &team2.player1.hand, 5)
  drawHand(&deck, &team2.player2.hand, 5)


  fmt.Printf("\nplayer1 from team1:\n%+v\n",team1.player1)
  fmt.Printf("player2 from team1:\n%+v\n\n",team1.player2)
  fmt.Printf("player1 from team2:\n%+v\n\n",team2.player1)
  fmt.Printf("player2 from team2:\n%+v\n\n",team2.player2)

  fmt.Printf("deck:\n%+v\n\n",deck)

  playHand(&team1, &team2, gState, team1.player1)

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
    } else {
      fmt.Println("Invalid command\n")
    }

    newCard := drawCard(&deck)
    fmt.Println(newCard)
    fmt.Println(deck.inDeckCards)

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
