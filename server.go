package main

import (
  "fmt"
  "net/http"
  //"io/ioutil"
  "github.com/gorilla/mux"
)

var deck Deck

func home(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html")
  w.Header().Set("Access-Control-Allow-Origin", "*")

  fmt.Fprint(w, "<h1>Welcome to my awesome site!<h1>")
}

func cards(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "*")

  card := drawCard(&deck)
  printCard(card)
  fmt.Printf("\n")

  //printDeck(deck)
  fmt.Fprint(w, "%+v", deck)
}

func play(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html; charset=utf-8")
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Access-Control-Allow-Headers", "*")

  //reqBody, _ := ioutil.ReadAll(r.Body)

  // for k, v := range r.Header {
  //   fmt.Fprintf(w, "Header field %q, Value %q\n", k, v)
  // }

  reqHeader := r.Header.Get("playerID")
  fmt.Fprintf(w, "%+v", string(reqHeader))
}

func main() {

  router := mux.NewRouter()

  deck = *makeDeck("euchre")

  router.HandleFunc("/", home)
  router.HandleFunc("/cards", cards)
  router.HandleFunc("/play", play).Methods("POST", "OPTIONS")
  http.ListenAndServe(":3001", router)

}
