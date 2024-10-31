package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"strings"
)

const (
	BOARD_WIDTH  = 10
	BOARD_HEIGHT = 10

	MAX_RANDOM_LIMIT = 1000

	HORIZONTAL Orientation = "HORIZONTAL"
	VERTICAL   Orientation = "VERTICAL"
)

type Orientation string

// Coordinates are zero based, and therefore
// - x  is between 0 and BOARD_WIDHT - 1
// - y  is between 0 and BOARD_HEIGHT - 1
type Coordinate [2]int

func NewCoordinate(x, y int) (Coordinate, error) {
	if x < 0 || x >= BOARD_WIDTH {
		return Coordinate{}, errors.New(fmt.Sprintf("x value %d out of bounds", x))
	}

	if y < 0 || y >= BOARD_HEIGHT {
		return Coordinate{}, errors.New(fmt.Sprintf("y value %d out of bounds", y))
	}

	return Coordinate{x, y}, nil
}

type Cell struct {
	occupied bool
	chosen   bool
}

type Board []Cell

func NewBoard() Board {
	board := make(Board, BOARD_HEIGHT*BOARD_WIDTH)
	for i := range board {
		board[i].chosen = false
		board[i].occupied = false
	}
	return board
}

func (b Board) String() string {
	builder := strings.Builder{}
	builder.WriteString("\n")

	count := 1
	for i := range b {
		if b[i].occupied {
			builder.WriteString(" X ")
		} else {
			builder.WriteString(" O ")
		}

		if count == BOARD_WIDTH {
			builder.WriteString("\n")
			count = 1
		} else {
			count += 1
		}
	}

	return builder.String()
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
	if rand.IntN(2) == 0 {
		return HORIZONTAL
	} else {
		return VERTICAL
	}
}

func get_random_coord(size int) Coordinate {
	x := rand.IntN(BOARD_WIDTH - size)
	y := rand.IntN(BOARD_HEIGHT - size)
	//fmt.Printf("Coord{%d,%d} - ", x, y)
	return Coordinate{x, y}
}

func (p *Player) place_ship(size int, orientation Orientation, coord Coordinate) error {
	x, y := coord[0], coord[1]
	idx := y*BOARD_WIDTH + x

	switch orientation {
	case HORIZONTAL:
		for i := 0; i < size; i++ {
			cell := p.player_board[idx+i]
			// fmt.Printf("index=%d - cell=%v - ", idx+i, cell)
			if cell.occupied {
				msg := fmt.Sprintf("Cell at %v already occupied", coord)
				//fmt.Print(msg + "\n")
				return errors.New(msg)
			}
		}
	case VERTICAL:
		for i := 0; i < size; i++ {
			cell := p.player_board[idx+(i*BOARD_WIDTH)]
			// fmt.Printf("index=%d - cell=%v - ", idx+(i*BOARD_WIDTH), cell)
			if cell.occupied {
				msg := fmt.Sprintf("Cell at %v already occupied", coord)
				//fmt.Print(msg + "\n")
				return errors.New(msg)
			}
		}
	default:
		msg := fmt.Sprintf("Unknown orientation: %s", orientation)
		//fmt.Print(msg + "\n")
		return errors.New(msg)
	}

	switch orientation {
	case HORIZONTAL:
		for i := 0; i < size; i++ {
			p.player_board[idx+i].occupied = true
		}
	case VERTICAL:
		for i := 0; i < size; i++ {
			p.player_board[idx+(i*BOARD_WIDTH)].occupied = true
		}
	}

	// msg := fmt.Sprintf("Placed Ship: %v [%s]", coord, orientation)
	//fmt.Println(msg)

	return nil
}

func verify_coordinate(size int, orientation Orientation, coord Coordinate) error {
	if coord[0] < 0 || coord[1] < 0 {
		msg := fmt.Sprintf("Carrier at %v [%s] is off the board", coord, orientation)
		//fmt.Print(msg + "\n")
		return errors.New(msg)
	}

	if coord[0] > BOARD_WIDTH || coord[1] > BOARD_HEIGHT {
		msg := fmt.Sprintf("Carrier at %v [%s] is off the board", coord, orientation)
		//fmt.Print(msg + "\n")
		return errors.New(msg)
	}

	if orientation == HORIZONTAL && coord[0] > BOARD_WIDTH-size {
		msg := fmt.Sprintf("Carrier at %v [%s] is off the board", coord, orientation)
		//fmt.Print(msg + "\n")
		return errors.New(msg)
	}

	if orientation == VERTICAL && coord[1] > BOARD_HEIGHT-size {
		msg := fmt.Sprintf("Carrier at %v [%s] is off the board", coord, orientation)
		//fmt.Print(msg + "\n")
		return errors.New(msg)
	}

	return nil
}

func (p *Player) randomize_ship_placement(size int) error {

	orientation := get_random_orientation()
	coord := get_random_coord(size)

	var err error
	switch size {
	case 2:
		err = p.place_destroyer(orientation, coord)
	case 3:
		err = p.place_cruiser_or_submarine(orientation, coord)
	case 4:
		err = p.place_battleship(orientation, coord)
	case 5:
		err = p.place_carrier(orientation, coord)
	default:
		return errors.New("Unknown ship with unknown size")
	}

	attempt := 0
	for err != nil {

		orientation := get_random_orientation()
		coord := get_random_coord(size)

		switch size {
		case 2:
			err = p.place_destroyer(orientation, coord)
		case 3:
			err = p.place_cruiser_or_submarine(orientation, coord)
		case 4:
			err = p.place_battleship(orientation, coord)
		case 5:
			err = p.place_carrier(orientation, coord)
		default:
			return errors.New("Unknown ship with unknown size")
		}

		if attempt == MAX_RANDOM_LIMIT {
			return errors.New("Max random limit reached. Cannot place ship!")
		} else {
			attempt += 1
		}
	}

	return nil
}

// Places ships on player_board in random fashion
func (p *Player) randomize_placement() error {

	if err := p.randomize_ship_placement(5); err != nil {
		return err
	}

	if err := p.randomize_ship_placement(4); err != nil {
		return err
	}

	if err := p.randomize_ship_placement(3); err != nil {
		return err
	}

	if err := p.randomize_ship_placement(3); err != nil {
		return err
	}

	if err := p.randomize_ship_placement(2); err != nil {
		return err
	}

	return nil
}

// Places 5 square ship on player_board. Errors if invalid placement.
func (p *Player) place_carrier(orientation Orientation, coord Coordinate) error {
	if err := verify_coordinate(5, orientation, coord); err != nil {
		return err
	}

	if err := p.place_ship(5, orientation, coord); err != nil {
		return err
	}

	return nil
}

// Places 4 square ship on player_board. Errors if invalid placement.
func (p *Player) place_battleship(orientation Orientation, coord Coordinate) error {
	if err := verify_coordinate(4, orientation, coord); err != nil {
		return err
	}

	if err := p.place_ship(4, orientation, coord); err != nil {
		return err
	}

	return nil
}

// Places 3 square ship on player_board. Errors if invalid placement.
func (p *Player) place_cruiser_or_submarine(orientation Orientation, coord Coordinate) error {
	if err := verify_coordinate(3, orientation, coord); err != nil {
		return err
	}

	if err := p.place_ship(3, orientation, coord); err != nil {
		return err
	}

	return nil
}

// Places 2 square ship on player_board. Errors if invalid placement.
func (p *Player) place_destroyer(orientation Orientation, coord Coordinate) error {
	if err := verify_coordinate(2, orientation, coord); err != nil {
		return err
	}

	if err := p.place_ship(2, orientation, coord); err != nil {
		return err
	}

	return nil
}

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
	p1 := NewPlayer("player 1")
	p1.randomize_placement()
	p2 := NewPlayer("player 2")
	p2.randomize_placement()

	var winner *Player

	turn_player := p1
	enemy_player := p2

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
		fmt.Printf("The winner is %s!\n", winner.name)
	}
}

// checks if the target_board has hit all ships on enemy_board
func check_winner(target_board, enemy_board Board) bool { return false }
