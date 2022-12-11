package server

import (
	
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
		matchManager.waitingPlayers = append(matchManager.waitingPlayers, WaitingPlayer{player: player, duration: requestedmatch["duration"]})
	}
}

func (matchManager *MatchManager) processWaitingPlayers(finishChannel chan bool) {
	
	var finish bool = false
	for !finish {
		for index, waitingPlayer := range matchManager.waitingPlayers {
			matchManager.mutexMatches.Lock()
			for _, match := range matchManager.matches {
				if waitingPlayer.duration == match.duration && !match.readyToStart {
					match.addPlayerToMatch(waitingPlayer.player)
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
	var finish bool = false
	for !finish {
		matchManager.mutexMatches.Lock()
		for _, match := range matchManager.matches {
			if match.readyToStart && !match.started {
				match.started = true
				go match.beginGame()
			}
			if len(finishChannel) > 0 {
				finish = <-finishChannel
			}
		}
		matchManager.mutexMatches.Unlock()
	}
}

func (matchManager *MatchManager) delete_finish_matches() {

	temp := matchManager.matches[:0]
	matchManager.mutexMatches.Lock()
	for _, match := range matchManager.matches {
		if !match.finish {
			temp = append(temp, match)
		}
	}
	matchManager.mutexMatches.Unlock()
	matchManager.matches = temp

}
