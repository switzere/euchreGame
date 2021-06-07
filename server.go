package main

import (
  "fmt"
  "net/http"
  // "io/ioutil"
  "github.com/gorilla/mux"
  "encoding/json"
)

var deck Deck
var hand Hand

var playerA Player
var playerB Player
var playerC Player
var playerD Player

type Card1 struct {
  suit int
  value int
}

type PID struct {
  PId int
}



func home(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html")
  w.Header().Set("Access-Control-Allow-Origin", "*")

  fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func cards(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "*")

  card := drawCard(&deck)
  printCard(card)
  fmt.Printf("\n")

  printDeck(deck)
  fmt.Fprint(w, "%+v", deck)
}

func play(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html; charset=utf-8")
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Access-Control-Allow-Headers", "*")

  fmt.Fprintf(w, "Hi")

  var card Card

  err := json.NewDecoder(r.Body).Decode(&card)
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }
  fmt.Fprintf(w, "Card: %+v", card)

  _ = drawSpecificCard(&deck, card)
  printCard(card)
  fmt.Printf("\n")
  printDeck(deck)



  reqHeader := r.Header.Get("playerID")
  fmt.Fprintf(w, "playerID: %+v\n", string(reqHeader))
  n := r.Header.Get("name")
  fmt.Fprintf(w, "name: %+v\n", string(n))

}

func resetDeck(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "*")

  deck = *makeDeck("euchre")
}

func drawPlayerHand(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "*")


  drawHand(&deck, &hand, 5)
  printHand(hand)

  fmt.Fprintf(w, "Hand: %+v", hand)
}

func drawPlayerSpecificHand(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html; charset=utf-8")
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Access-Control-Allow-Headers", "*")

  var PId PID


  err := json.NewDecoder(r.Body).Decode(&PId)
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

  fmt.Printf("In dPSH")

  if playerA.PId == PId.PId {
    drawHandUnique(&deck, &playerA, 5)
    fmt.Fprintf(w, "Player: %+v", playerA)
  } else if playerB.PId == PId.PId {
    drawHandUnique(&deck, &playerB, 5)
    fmt.Fprintf(w, "Player: %+v", playerB)
  } else if playerC.PId == PId.PId {
    drawHandUnique(&deck, &playerC, 5)
    fmt.Fprintf(w, "Player: %+v", playerC)
  } else if playerD.PId == PId.PId {
    drawHandUnique(&deck, &playerD, 5)
    fmt.Fprintf(w, "Player: %+v", playerD)
  } else {
    fmt.Fprintf(w, "Error")
    fmt.Fprintf(w, "%+v", PId.PId)
    fmt.Fprintf(w, "%+v", playerB.PId)
  }


}

func main() {

  router := mux.NewRouter()


  deck = *makeDeck("euchre")

  //Should probably be replaced with a call but works for now if only 4 people are playing
  // playerA := new(Player)
  // playerB := new(Player)
  // playerC := new(Player)
  // playerD := new(Player)

  playerA.PId = 0
  playerB.PId = 1
  playerC.PId = 2
  playerD.PId = 3


  //team1 := Team{Player1: *playerA, Player2: *playerB, Points: 0}
  //team2 := Team{Player1: *playerC, Player2: *playerD, Points: 0}

  //gState := gameState{Order: []Player{team1.Player1, team2.Player1, team1.Player2, team2.Player2}}


  router.HandleFunc("/", home)
  router.HandleFunc("/cards", cards)
  router.HandleFunc("/resetDeck", resetDeck)
  router.HandleFunc("/drawPlayerHand", drawPlayerHand)
  router.HandleFunc("/play", play).Methods("POST", "OPTIONS")
  router.HandleFunc("/drawPlayerSpecificHand", drawPlayerSpecificHand).Methods("POST", "OPTIONS")
  http.ListenAndServe(":3001", router)

}
