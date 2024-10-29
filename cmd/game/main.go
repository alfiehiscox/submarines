package main

import (
	"fmt"
	"math/rand"
)

const (
	BOARD_WIDTH  = 10
	BOARD_HEIGHT = 10

	MAX_RANDOM_LIMIT = 100

	HORIZONTAL Orientation = "HORIZONTAL"
	VERTICAL   Orientation = "VERTICAL"
)

type Orientation string
type Coordinate [2]int

type Cell struct {
	occupied bool
	chosen   bool
}

type Board []Cell

func NewBoard() Board {
	return make(Board, BOARD_HEIGHT*BOARD_WIDTH)
}

type Player struct {
	name string

	// Holds no information about turns, only ships
	player_board Board

	// Holds no information about ships, only turns
	target_board Board
}

func NewPlayer(name string) *Player {
	return &Player{
		name:         name,
		player_board: NewBoard(),
		target_board: NewBoard(),
	}
}

func get_random_orientation() Orientation {
	if rand.Intn(2) == 0 {
		return HORIZONTAL
	} else {
		return VERTICAL
	}
}

// Places ships on player_board in random fashion
func (p *Player) randomize_placement() {}

// Places 5 square ship on player_board. Errors if invalid placement.
func (p *Player) place_carrier(orientation Orientation, coord Coordinate) error { return nil }

// Places 4 square ship on player_board. Errors if invalid placement.
func (p *Player) place_battleship(orientation Orientation, coord Coordinate) error { return nil }

// Places 3 square ship on player_board. Errors if invalid placement.
func (p *Player) place_cruiser(orientation Orientation, coord Coordinate) error { return nil }

// Places 3 square ship on player_board. Errors if invalid placement.
func (p *Player) place_submarine(orientation Orientation, coord Coordinate) error { return nil }

// Places 2 square ship on player_board. Errors if invalid placement.
func (p *Player) place_destroyer(orientation Orientation, coord Coordinate) error { return nil }

// Get's a player's coordinate guess checking against their
// previous turns (i.e. target_board)
func (p *Player) get_guess() Coordinate {
	return Coordinate{0, 0}
}

// Check's if coordinate hits a ship in player_board
func (p *Player) check_hit(coordinate Coordinate) bool {
	return false
}

// Mark an attempt on target_board
func (p *Player) mark_target_attempt(coordiant Coordinate, hit bool) {}

// Mark an attempt on player_board
func (p *Player) mark_player_attempt(coordiant Coordinate, hit bool) {}

func main() {
	p1 := Player{name: "player 1"}
	p1.randomize_placement()
	p2 := Player{name: "player 2"}
	p2.randomize_placement()

	var winner *Player

	turn_player := &p1
	enemy_player := &p2

	for {

		coord := turn_player.get_guess()

		hit := enemy_player.check_hit(coord)
		turn_player.mark_target_attempt(coord, hit)
		enemy_player.mark_player_attempt(coord, hit)

		if check_winner(turn_player.target_board, enemy_player.player_board) {
			winner = turn_player
			break
		}

		turn_player, enemy_player = enemy_player, turn_player
	}

	if winner != nil {
		fmt.Printf("The winner is %d!\n", winner.name)
	}
}

// checks if the target_board has hit all ships on enemy_board
func check_winner(target_board, enemy_board Board) bool { return false }
