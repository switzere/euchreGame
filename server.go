package main

import (
	"fmt"
	"net/http"

	// "io/ioutil"
	"encoding/json"

	"github.com/gorilla/mux"
)

var deck Deck

var team1 Team
var team2 Team

var turnedUpCard Card
var trump int
var playerOrder []Player
var roundPlayerOrder []Player
var playedCards []Card
var dealer Player

var playerA Player
var playerB Player
var playerC Player
var playerD Player

type PID struct {
	PId int
}

type Data struct {
	Card   Card
	PId    int
	Choice bool
}

type RoundStack struct {
	PlayedCards   []Card
	WinningPlayer Player
	PlayerTurn    Player
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

	//drawHand(&deck, &hand, 5)
	//printHand(hand)

	//fmt.Fprintf(w, "Hand: %+v", hand)
}

func startGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	drawHand(&deck, &playerA.Hand, 5)
	drawHand(&deck, &playerB.Hand, 5)
	drawHand(&deck, &playerC.Hand, 5)
	drawHand(&deck, &playerD.Hand, 5)

	fmt.Fprintf(w, "PlayerA: %+v", playerA)
	fmt.Fprintf(w, "PlayerB: %+v", playerB)
	fmt.Fprintf(w, "PlayerC: %+v", playerC)
	fmt.Fprintf(w, "PlayerD: %+v", playerD)

}

func getPlayerHand(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	var PId PID

	err := json.NewDecoder(r.Body).Decode(&PId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//using pID get player and return hand
	if playerA.PId == PId.PId {
		fmt.Fprintf(w, "PlayerA: %+v", playerA)
	} else if playerB.PId == PId.PId {
		fmt.Fprintf(w, "PlayerB: %+v", playerB)
	} else if playerC.PId == PId.PId {
		fmt.Fprintf(w, "PlayerC: %+v", playerC)
	} else if playerD.PId == PId.PId {
		fmt.Fprintf(w, "PlayerD: %+v", playerD)
	} else {
		fmt.Fprintf(w, "Error")
		fmt.Fprintf(w, "%+v", PId.PId)
	}
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

func getHand(pId int) Hand {
	if playerA.PId == pId {
		return playerA.Hand
	}
	if playerB.PId == pId {
		return playerB.Hand
	}
	if playerC.PId == pId {
		return playerC.Hand
	}
	if playerD.PId == pId {
		return playerD.Hand
	}
	return Hand{}
}

//////////////////////////
//
//
//
//
//
//
//
//
//divider
//
//
//
//
//
//////////////////////////

func initGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	deck = *makeDeck("euchre")

	playerA.PId = 0
	playerB.PId = 1
	playerC.PId = 2
	playerD.PId = 3

	team1 = Team{Player1: playerA, Player2: playerB, Points: 0}
	team2 = Team{Player1: playerC, Player2: playerD, Points: 0}

}

func initRound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	drawHand(&deck, &playerA.Hand, 5)
	drawHand(&deck, &playerB.Hand, 5)
	drawHand(&deck, &playerC.Hand, 5)
	drawHand(&deck, &playerD.Hand, 5)

	//will need to cycle round order after every round
	roundPlayerOrder = []Player{playerA, playerB, playerC, playerD}
	playerOrder = []Player{playerA, playerB, playerC, playerD}
	dealer = roundPlayerOrder[0]

	turnedUpCard = drawCard(&deck)

}

func getRoundInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	var PId PID

	err := json.NewDecoder(r.Body).Decode(&PId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if playerA.PId == PId.PId {
		fmt.Fprintf(w, "PlayerA: %+v\n", playerA)
	} else if playerB.PId == PId.PId {
		fmt.Fprintf(w, "PlayerB: %+v\n", playerB)
	} else if playerC.PId == PId.PId {
		fmt.Fprintf(w, "PlayerC: %+v\n", playerC)
	} else if playerD.PId == PId.PId {
		fmt.Fprintf(w, "PlayerD: %+v\n", playerD)
	} else {
		fmt.Fprintf(w, "Error")
		fmt.Fprintf(w, "%+v", PId.PId)
	}
	fmt.Fprintf(w, "Turned Up Card: %+v\n", turnedUpCard)
	fmt.Fprintf(w, "Trump: %+v\n", trump)
	fmt.Fprintf(w, "Player Turn: %+v\n", playerOrder[0].PId)
	fmt.Fprintf(w, "Dealer: %+v\n", dealer.PId)

}

func chooseTrump(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	var data Data

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(data)
	if data.PId == playerOrder[0].PId {
		if data.Choice {
			trump = turnedUpCard.Suit
			fmt.Fprintf(w, "Trump set to: %+v\n", trump)
			playerOrder = roundPlayerOrder
		} else {
			var x Player
			x, playerOrder = playerOrder[0], playerOrder[1:]
			playerOrder = append(playerOrder, x)
		}
	} else {
		fmt.Fprint(w, "Error - Not your turn\n", data.PId)
	}
	fmt.Fprintf(w, "Player Turn: %+v\n", playerOrder[0].PId)

}

func dealerPickUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	var data Data

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if data.PId == dealer.PId {
		allowed := swapWithCardInHand(returnPlayerOnPId([]*Player{&playerA, &playerB, &playerC, &playerD}, data.PId), data.Card, turnedUpCard)
		if !allowed {
			fmt.Fprint(w, "Error - Card not in hand\n", data.PId)
		} else {
			fmt.Fprintf(w, "Put Down: %+v\n", data.Card)
		}
	}
}

func playCardFromHand(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var data Data

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if data.PId == playerOrder[0].PId {
		fmt.Fprintf(w, "Player: %+v\n", data.PId)
		fmt.Fprintf(w, "Card: %+v\n", data.Card)
		fmt.Fprintf(w, "Player Turn: %+v\n", playerOrder[1].PId)
		var x Player
		x, playerOrder = playerOrder[0], playerOrder[1:]
		playerOrder = append(playerOrder, x)
	} else {
		fmt.Fprint(w, "Error - Not your turn\n", data.PId)
	}
	fmt.Fprintf(w, "Player Turn: %+v\n", playerOrder[0].PId)

}

func main() {
	//go run server.go deck.go player.go print.go

	fmt.Printf("Launching...")
	router := mux.NewRouter()

	// deck = *makeDeck("euchre")

	//Should probably be replaced with a call but works for now if only 4 people are playing
	// playerA := new(Player)
	// playerB := new(Player)
	// playerC := new(Player)
	// playerD := new(Player)

	// playerA.PId = 0
	// playerB.PId = 1
	// playerC.PId = 2
	// playerD.PId = 3

	//team1 := Team{Player1: *playerA, Player2: *playerB, Points: 0}
	//team2 := Team{Player1: *playerC, Player2: *playerD, Points: 0}

	//gState := gameState{Order: []Player{team1.Player1, team2.Player1, team1.Player2, team2.Player2}}

	router.HandleFunc("/", home)
	router.HandleFunc("/cards", cards)
	router.HandleFunc("/resetDeck", resetDeck)
	router.HandleFunc("/startGame", startGame)
	router.HandleFunc("/play", play).Methods("POST", "OPTIONS")
	router.HandleFunc("/drawPlayerSpecificHand", drawPlayerSpecificHand).Methods("POST", "OPTIONS")

	router.HandleFunc("/drawPlayerHand", drawPlayerHand)
	router.HandleFunc("/getPlayerHand", getPlayerHand).Methods("POST", "OPTIONS")

	router.HandleFunc("/initGame", initGame)
	router.HandleFunc("/initRound", initRound)
	router.HandleFunc("/getRoundInfo", getRoundInfo).Methods("POST", "OPTIONS")
	router.HandleFunc("/chooseTrump", chooseTrump).Methods("POST", "OPTIONS")
	router.HandleFunc("/dealerPickUp", dealerPickUp).Methods("POST", "OPTIONS")

	http.ListenAndServe(":3001", router)

}
