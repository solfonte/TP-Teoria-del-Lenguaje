package server

type Round struct {
	players       map[int]*Player
	moves         int
	watingPlayer  *Player
	currentPlayer *Player
	championId    int
	envido        bool
	cardsPlayed   []Card
}

func (Round *Round) initialize(players map[int](*Player)) {
	Round.players = players
	Round.moves = 0
	Round.championId = -1
	Round.envido = false

}

func (round *Round) startRound() int {
	completeRound := 1
	finish := false
	round.decide_hand_players()
	for completeRound <= 3 || !finish {
		var move = Move{typeMove: completeRound}
		finish = move.start_move(round.currentPlayer, round.watingPlayer)
		completeRound += 1
		round.currentPlayer = round.players[move.winner]
		round.watingPlayer = round.players[move.loser]
	}
	return 0
}
func (round *Round) decide_hand_players() {
	//Mañana veo
	round.watingPlayer = round.players[1]
	round.currentPlayer = round.players[2]
}

// func (round *Round) waitingPlayerId() int {
// 	fmt.Println("jugador actual id", round.currentPlayerId)
// 	if round.currentPlayerId == 0 {
// 		fmt.Println("entre a  caso donde cambio a 1")
// 		return 1
// 	} else {
// 		fmt.Println("entre a  caso donde cambio a 0")
// 		return 0
// 	}
// }
