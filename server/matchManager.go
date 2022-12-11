package server

import (
	"fmt"
	"sync"
)

type WaitingPlayer struct {
	duration int
	player   *Player
}
type MatchManager struct {
	matches        []*Match
	waitingPlayers []WaitingPlayer
	mutexMatches   sync.Mutex
	finishPlayers  []*Player
}

func (matchManager *MatchManager) process_player(player *Player) {
	player.welcomePlayer()
	messageClient, _ := sendMenu(*player)
	requestedmatch := processRequest(*player, messageClient)
	if requestedmatch["create"] == 0 {
		newMatch := Match{duration: requestedmatch["duration"], started: false, players: make(map[int]*Player), readyToStart: false}
		newMatch.addPlayerToMatch(player)
		matchManager.mutexMatches.Lock()
		matchManager.matches = append(matchManager.matches, &newMatch)
		matchManager.mutexMatches.Unlock()
	} else {
		fmt.Println("Guardo al jugador que hizo join en la cola de jugadores esperando")
		matchManager.mutexMatches.Lock()
		matchManager.waitingPlayers = append(matchManager.waitingPlayers, WaitingPlayer{player: player, duration: requestedmatch["duration"]})
		matchManager.mutexMatches.Unlock()
	}
}

func (matchManager *MatchManager) processWaitingPlayers(finishChannel chan bool) {

	var finish bool = false
	for !finish {
		select {
		case finish = <-finishChannel:
			return
		default:
			for index, waitingPlayer := range matchManager.waitingPlayers {
				matchManager.mutexMatches.Lock()
				for _, match := range matchManager.matches {
					if waitingPlayer.duration == match.duration && !match.readyToStart {
						if (waitingPlayer.player.isReadyToPlay()){
							match.addPlayerToMatch(waitingPlayer.player)
						}
						matchManager.waitingPlayers = append(matchManager.waitingPlayers[:index], matchManager.waitingPlayers[index+1:]...)
					}
				}
				matchManager.mutexMatches.Unlock()
			}
		}
	}
	fmt.Println("Sali de processwaitingPlayers")
}

func (matchManager *MatchManager) startMatches(finishChannel chan bool) {
	var finish bool = false
	for !finish {
		select {
		case finish = <-finishChannel:
			return
		default:
			matchManager.mutexMatches.Lock()
			for _, match := range matchManager.matches {
				if !match.started && matchManager.cancelMatch(match) {
					fmt.Println("esta para cancelar el match")
					matchManager.addMatchPlayersToWaitingQueue(match)
					match.finish = true
				}else if match.readyToStart && !match.started {
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
	fmt.Println("Sali de startMatches")
}

func (matchManager *MatchManager) deleteFinishMatches(finishChannel chan bool) {
	var finish bool = false
	for !finish {
		select {
		case finish = <-finishChannel:
			return
		default:
			temp := matchManager.matches[:0]
			matchManager.mutexMatches.Lock()
			for _, match := range matchManager.matches {
				if !match.finish {
					temp = append(temp, match)
				} else {
					player1, player2 := match.getPlayersMatch()
					matchManager.finishPlayers = append(matchManager.finishPlayers, player1)
					matchManager.finishPlayers = append(matchManager.finishPlayers, player2)
				}
			}
			matchManager.mutexMatches.Unlock()
			matchManager.matches = temp
		}
	}
	fmt.Println("sali de delete matches")
}

func (matchManager *MatchManager) playerFinish(player Player) bool {
	for _, playerFinish := range matchManager.finishPlayers {
		if playerFinish.id == player.id {
			return true
		}
	}
	return false
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
		if p.isReadyToPlay(){
			fmt.Print(p.name, " is connected when adding to waiting queue")
			matchManager.mutexMatches.Lock()
			matchManager.waitingPlayers = append(matchManager.waitingPlayers, WaitingPlayer{player: p, duration: match.duration})
			matchManager.mutexMatches.Unlock()
		}
	}
}

