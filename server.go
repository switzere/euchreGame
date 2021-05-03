package main

import (
  "fmt"
  "net/http"
  // "io/ioutil"
  "github.com/gorilla/mux"
  "encoding/json"
)

var deck Deck

type Card1 struct {
  suit int
  value int
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

func main() {

  router := mux.NewRouter()


  deck = *makeDeck("euchre")

  router.HandleFunc("/", home)
  router.HandleFunc("/cards", cards)
  router.HandleFunc("/resetDeck", resetDeck)
  router.HandleFunc("/play", play).Methods("POST", "OPTIONS")
  http.ListenAndServe(":3001", router)

}
