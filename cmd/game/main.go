package main

import (
	"fmt"
	"time"

	"github.com/alfiehiscox/submarines/pkg/board"
	"github.com/alfiehiscox/submarines/pkg/player"
)

func main() {
	p1 := player.NewPlayer("player 1")
	p1.RandomizePlacement()
	p2 := player.NewPlayer("player 2")
	p2.RandomizePlacement()

	var winner *player.Player

	turn_player := p1
	enemy_player := p2

	for {

		coord := turn_player.GetGuess()

		hit := enemy_player.CheckHit(coord)
		turn_player.MarkTargetAttempt(coord, hit)
		enemy_player.MarkPlayerAttempt(coord, hit)

		if board.CheckWinner(turn_player.TargetBoard, enemy_player.PlayerBoard) {
			winner = turn_player
			break
		}

		turn_player, enemy_player = enemy_player, turn_player
		time.Sleep(time.Second)
	}

	if winner != nil {
		fmt.Printf("The winner is %s!\n", winner.Name)
	}
}
