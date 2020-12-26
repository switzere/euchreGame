package main

import (
  "fmt"
)

//TODO fix order, I think teams are grouped currently
//TODO force following suit
//TODO go around letting people pass or play
//TODO choose trump after turned down
//TODO change start order after hands
//TODO validate input

//TODO Re-factor everything its a huge mess

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


func winRound(cards []Card, trump int) int {

  var numTrump int = 0
  var largestTrump int = -1
  var largestOfLed int = -1
  var memCards [4]Card

  led := cards[0].suit

  if cards[0].suit == Spade {
    fmt.Printf("Spade from const\n")
  }
  if cards[0].suit == 0 {
    fmt.Printf("Spade from num\n")
  }

  //adjusting for jacks
  for i := 0; i < 4; i++ {

    if trump == Heart {
      if cards[i].suit == Diamond && cards[i].value == 11 {
        cards[i].value = 15
        cards[i].suit = Heart
      } else if cards[i].suit == Heart && cards[i].value == 11 {
        cards[i].value = 16
      }
    } else if trump == Diamond {
      if cards[i].suit == Heart && cards[i].value == 11 {
        cards[i].value = 15
        cards[i].suit = Diamond
      } else if cards[i].suit == Diamond && cards[i].value == 11 {
        cards[i].value = 16
      }
    } else if trump == Spade {
      if cards[i].suit == Club && cards[i].value == 11 {
        cards[i].value = 15
        cards[i].suit = Spade
      } else if cards[i].suit == Spade && cards[i].value == 11 {
        cards[i].value = 16
      }
    } else if trump == Club {
      if cards[i].suit == Spade && cards[i].value == 11 {
        cards[i].value = 15
        cards[i].suit = Club
      } else if cards[i].suit == Club && cards[i].value == 11 {
        cards[i].value = 16
      }
    }

    if cards[i].suit == trump {
      numTrump++
    }
    memCards[i] = cards[i]
  }

  fmt.Printf("adjusted cards: %+v\n",memCards)


  //until 1 cards is determined the largest
  for len(cards) > 1 {

    //remove all cards that don't matter
    for i := 0; i < len(cards); i++ {

      if cards[i].suit != trump && numTrump > 0 {
        fmt.Printf("%+v not trump when trump was played\n",cards[i])
        cards = append(cards[:i], cards[i+1:]...)
        i = 0
      } else if cards[i].suit != led && numTrump == 0 {
        fmt.Printf("%+v not following suit when no trump was played\n",cards[i])
        cards = append(cards[:i], cards[i+1:]...)
        i = 0
      }

    }

    //find which of remaining trump is largest
    if numTrump == len(cards) {

      fmt.Printf("\n%+v\n\n",cards)

      for i := 0; i < len(cards); i++ {
        fmt.Printf("%d > %d\n",cards[i].value, largestTrump)
        if cards[i].value > largestTrump {
          fmt.Printf("%+v currently largest trump\n",cards[i])
          largestTrump = cards[i].value
        } else if cards[i].value < largestTrump {
          fmt.Printf("%+v trump but smaller\n",cards[i])
          cards = append(cards[:i], cards[i+1:]...)
          i = -1
        }
        fmt.Printf("\n%+v\n\n",cards)
      }

      //else do largest of led suit
    } else {
      for i := 0; i < len(cards); i++ {
        if cards[i].value > largestOfLed {
          fmt.Printf("%+v currently largest led\n",cards[i])
          largestOfLed = cards[i].value
        } else if cards[i].value < largestOfLed {
          fmt.Printf("%+v led but smaller\n",cards[i])
          cards = append(cards[:i], cards[i+1:]...)
          i = -1
        }
        fmt.Printf("\n%+v\n\n",cards)
      }
    }

  }

  fmt.Printf("memCards: %+v\n",memCards)
  for i := 0; i < 4; i++ {
    if memCards[i] == cards[0] {
      return i
    }
  }

  fmt.Printf("\n%+v\n\n",cards)
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

  var suit int = -1
  var card int = -1

  var inSuit string = ""
  var inCard string = ""

  var cardsPlayed []Card

  for i := 0; i < 4; i++ {
    if team1.player1.pId == gS.order[i].pId {
      fmt.Printf("\nTeam 1 Player 1:\n")
      printHand(team1.player1.hand)
      fmt.Printf("Choose a card to play: ")
      fmt.Scanf("%s of %s\n",&inCard,&inSuit)
      suit = suitToNum(inSuit)
      card = faceToNum(inCard)
      removeCard := Card{suit, card}
      cardsPlayed = append(cardsPlayed, playCard(&team1.player1, removeCard))

    } else if team1.player2.pId == gS.order[i].pId {
      fmt.Printf("\nTeam 1 Player 2:\n")
      printHand(team1.player2.hand)
      fmt.Printf("Choose a card to play: ")
      fmt.Scanf("%s of %s\n",&inCard,&inSuit)
      suit = suitToNum(inSuit)
      card = faceToNum(inCard)
      removeCard := Card{suit, card}
      cardsPlayed = append(cardsPlayed, playCard(&team1.player2, removeCard))

    } else if team2.player1.pId == gS.order[i].pId {
      fmt.Printf("\nTeam 2 Player 1:\n")
      printHand(team2.player1.hand)
      fmt.Printf("Choose a card to play: ")
      fmt.Scanf("%s of %s\n",&inCard,&inSuit)
      suit = suitToNum(inSuit)
      card = faceToNum(inCard)
      removeCard := Card{suit, card}
      cardsPlayed = append(cardsPlayed, playCard(&team2.player1, removeCard))

    } else if team2.player2.pId == gS.order[i].pId {
      fmt.Printf("\nTeam 2 Player 2:\n")
      printHand(team2.player2.hand)
      fmt.Printf("Choose a card to play: ")
      fmt.Scanf("%s of %s\n",&inCard,&inSuit)
      suit = suitToNum(inSuit)
      card = faceToNum(inCard)
      removeCard := Card{suit, card}
      cardsPlayed = append(cardsPlayed, playCard(&team2.player2, removeCard))
    }
  }
  t := winRound(cardsPlayed, trump)

  if team1.player1.pId == gS.order[t].pId {
    team1.player1.hand.tricks++
    fmt.Printf("\nTeam 1:\n")
    printTeam(*team1)
    fmt.Printf("\nTeam 2:\n")
    printTeam(*team2)
    return team1.player1
  } else if team1.player2.pId == gS.order[t].pId {
    team1.player2.hand.tricks++
    fmt.Printf("\nTeam 1:\n")
    printTeam(*team1)
    fmt.Printf("\nTeam 2:\n")
    printTeam(*team2)
    return team1.player2
  } else if team2.player1.pId == gS.order[t].pId {
    team2.player1.hand.tricks++
    fmt.Printf("\nTeam 1:\n")
    printTeam(*team1)
    fmt.Printf("\nTeam 2:\n")
    printTeam(*team2)
    return team2.player1
  } else if team2.player2.pId == gS.order[t].pId {
    team2.player2.hand.tricks++
    fmt.Printf("\nTeam 1:\n")
    printTeam(*team1)
    fmt.Printf("\nTeam 2:\n")
    printTeam(*team2)
    return team2.player2
  }

  return team1.player1
}

func playHand(team1 *Team, team2 *Team, deck *Deck, gS gameState, firstPlayer Player) {

  var trump int = -1
  trump = 0

  var suit int = -1
  var card int = -1

  var inSuit string = ""
  var inCard string = ""

  turnUp := drawCard(deck)
  //fmt.Printf("Turned up card is: %+v\n\n",turnUp)
  fmt.Printf("Turned up card is: ")
  printCard(turnUp)
  fmt.Printf("\n\n")


  if team1.player1.pId == firstPlayer.pId {
    fmt.Printf("Replace card from(t1p1): ")
    printHand(team1.player1.hand)
    fmt.Scanf("%s of %s\n",&inCard,&inSuit)
    suit = suitToNum(inSuit)
    card = faceToNum(inCard)
    removeCard := Card{suit, card}
    playCard(&team1.player1, removeCard)
    getCard(&team1.player1, turnUp)
    trump = turnUp.suit

  } else if team1.player2.pId == firstPlayer.pId {
    fmt.Printf("Replace card from(t1p2): ")
    printHand(team1.player2.hand)
    fmt.Scanf("%s of %s\n",&inCard,&inSuit)
    suit = suitToNum(inSuit)
    card = faceToNum(inCard)
    removeCard := Card{suit, card}
    playCard(&team1.player2, removeCard)
    getCard(&team1.player2, turnUp)
    trump = turnUp.suit

  } else if team2.player1.pId == firstPlayer.pId {
    fmt.Printf("Replace card from(t2p1): ")
    printHand(team2.player1.hand)
    fmt.Scanf("%s of %s\n",&inCard,&inSuit)
    suit = suitToNum(inSuit)
    card = faceToNum(inCard)
    removeCard := Card{suit, card}
    playCard(&team2.player1, removeCard)
    getCard(&team2.player1, turnUp)
    trump = turnUp.suit

  } else if team2.player2.pId == firstPlayer.pId {
    fmt.Printf("Replace card from(t2p2): ")
    printHand(team2.player2.hand)
    fmt.Scanf("%s of %s\n",&inCard,&inSuit)
    suit = suitToNum(inSuit)
    card = faceToNum(inCard)
    removeCard := Card{suit, card}
    playCard(&team2.player2, removeCard)
    getCard(&team2.player2, turnUp)
    trump = turnUp.suit
  }


  for i := 0; i < 5; i++ {
    firstPlayerFound := false

    //rotate through list to find first player
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

  gState := gameState{order: []Player{team1.player1, team2.player1, team1.player2, team2.player2}}

  deck := *makeDeck("euchre")

  //fmt.Printf("%+v\n",deck)
  printDeck(deck)

  drawHand(&deck, &team1.player1.hand, 5)
  drawHand(&deck, &team1.player2.hand, 5)
  drawHand(&deck, &team2.player1.hand, 5)
  drawHand(&deck, &team2.player2.hand, 5)


  //fmt.Printf("\nplayer1 from team1:\n%+v\n",team1.player1)
  //fmt.Printf("player2 from team1:\n%+v\n\n",team1.player2)
  //fmt.Printf("player1 from team2:\n%+v\n\n",team2.player1)
  //fmt.Printf("player2 from team2:\n%+v\n\n",team2.player2)
  printTeam(team1)
  printTeam(team2)

  //fmt.Printf("deck:\n%+v\n\n",deck)
  printDeck(deck)

  playHand(&team1, &team2, &deck, gState, team1.player1)

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
