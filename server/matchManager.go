package server

import (
	"fmt"
	"sync"
)

type WaitingPlayer struct {
	duration   int
	maxPlayers int
	player     *Player
}
type MatchManager struct {
	matches        []*Match
	waitingPlayers []WaitingPlayer
	mutexMatches   sync.Mutex
}

func (matchManager *MatchManager) process_player(player *Player) {
	player.welcomePlayer()
	messageClient, _ := sendMenu(*player)
	requestedmatch := processRequest(*player, messageClient)
	if requestedmatch["create"] == 0 {
		newMatch := Match{duration: requestedmatch["duration"], maxPlayers: requestedmatch["members"], started: false, players: make(map[int]*Player), readyToStart: false}
		newMatch.addPlayerToMatch(player)
		matchManager.mutexMatches.Lock()
		matchManager.matches = append(matchManager.matches, &newMatch)
		matchManager.mutexMatches.Unlock()
	} else {
		fmt.Println("Guardo al jugador que hizo join en la cola de jugadores esperando")
		matchManager.mutexMatches.Lock()
		matchManager.waitingPlayers = append(matchManager.waitingPlayers, WaitingPlayer{player: player, duration: requestedmatch["duration"], maxPlayers: requestedmatch["members"]})
		matchManager.mutexMatches.Unlock()
	}
}

func (matchManager *MatchManager) processWaitingPlayers(finishChannel chan bool) {
	fmt.Println("Entre en prosess waiting players")
	var finish bool = false
	for !finish {
		for index, waitingPlayer := range matchManager.waitingPlayers {
			matchManager.mutexMatches.Lock()
			for _, match := range matchManager.matches {
				if waitingPlayer.duration == match.duration && waitingPlayer.maxPlayers == match.maxPlayers && !match.readyToStart {
					fmt.Println("Jugador cumple con condiciones de match ", waitingPlayer.player.name)
					match.addPlayerToMatch(waitingPlayer.player)
					fmt.Println("match started: ", match.started)
					matchManager.waitingPlayers = append(matchManager.waitingPlayers[:index], matchManager.waitingPlayers[index+1:]...)
				}
			}
			matchManager.mutexMatches.Unlock()
		}
		if len(finishChannel) > 0 {
			finish = <-finishChannel
		}
	}
}

func (matchManager *MatchManager) startMatches(finishChannel chan bool) {
	fmt.Println("entre a start matches")
	var finish bool = false
	for !finish {
		matchManager.mutexMatches.Lock()
		for _, match := range matchManager.matches {
			if !match.started && matchManager.cancelMatch(match) {
				matchManager.addMatchPlayersToWaitingQueue(match)
				match.finish = true
			}
			if match.readyToStart && !match.started {
				match.started = true
				fmt.Println("arranco match")
				go match.beginGame()
			}
			if len(finishChannel) > 0 {
				finish = <-finishChannel
			}
		}
		matchManager.mutexMatches.Unlock()
	}
}

func (matchManager *MatchManager) cancelMatch(match *Match) bool{
	if match.finish {
		return false
	}
	cancel := false
	for _, p := range match.players {
		if !p.isReadyToPlay() {
			cancel = true
		}
	}
	return cancel
}

func (matchManager *MatchManager) addMatchPlayersToWaitingQueue(match *Match){
	for _, p := range match.players {
		if p.isConnected(){
			fmt.Print(p.name, " is connected when adding to waiting queue")

			matchManager.mutexMatches.Lock()
			matchManager.waitingPlayers = append(matchManager.waitingPlayers, WaitingPlayer{player: p, duration: match.duration, maxPlayers: match.maxPlayers})
			matchManager.mutexMatches.Unlock()
		}
	}
}

func (matchManager *MatchManager) delete_finish_matches() {
	fmt.Println("entro a delete matches")
	temp := matchManager.matches[:0]
	matchManager.mutexMatches.Lock()
	for _, match := range matchManager.matches {
		fmt.Println("Entro al for de delete matches ")
		
		if !match.finish {
			temp = append(temp, match)
		}
	}
	matchManager.mutexMatches.Unlock()
	matchManager.matches = temp
	fmt.Println("cantidad matches: ", len(matchManager.matches))
}
