package player

import (
	"errors"
	"fmt"

	"github.com/alfiehiscox/submarines/pkg/board"
	"github.com/alfiehiscox/submarines/pkg/cell"
)

type Player struct {
	Name string

	// Holds no information about turns, only ships
	PlayerBoard board.Board

	// Holds no information about ships, only turns
	TargetBoard board.Board
}

func NewPlayer(name string) *Player {
	return &Player{
		Name:        name,
		PlayerBoard: board.NewBoard(),
		TargetBoard: board.NewBoard(),
	}
}

func (p *Player) place_ship(size int, orientation cell.Orientation, coord cell.Coordinate) error {

	if err := cell.VerifyCoordinate(size, orientation, coord); err != nil {
		return err
	}

	idx := coord.ToIndex()

	switch orientation {
	case cell.HORIZONTAL:
		for i := 0; i < size; i++ {
			cell := p.PlayerBoard[idx+i]
			if cell.Occupied {
				msg := fmt.Sprintf("Cell at %v already occupied", coord)
				return errors.New(msg)
			}
		}
	case cell.VERTICAL:
		for i := 0; i < size; i++ {
			cell := p.PlayerBoard[idx+(i*cell.BOARD_WIDTH)]
			if cell.Occupied {
				msg := fmt.Sprintf("Cell at %v already occupied", coord)
				return errors.New(msg)
			}
		}
	default:
		msg := fmt.Sprintf("Unknown orientation: %s", orientation)
		return errors.New(msg)
	}

	switch orientation {
	case cell.HORIZONTAL:
		for i := 0; i < size; i++ {
			p.PlayerBoard[idx+i].Occupied = true
		}
	case cell.VERTICAL:
		for i := 0; i < size; i++ {
			p.PlayerBoard[idx+(i*cell.BOARD_WIDTH)].Occupied = true
		}
	}

	return nil
}

func (p *Player) RandomizeShipPlacement(size int) error {

	orientation := cell.GetRandomOrientation()
	coord := cell.GetRandomCoord(size)

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

		orientation := cell.GetRandomOrientation()
		coord := cell.GetRandomCoord(size)

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

		if attempt == cell.MAX_RANDOM_LIMIT {
			return errors.New("Max random limit reached. Cannot place ship!")
		} else {
			attempt += 1
		}
	}

	return nil
}

// Places ships on player_board in random fashion
func (p *Player) RandomizePlacement() error {

	// Carrier
	if err := p.RandomizeShipPlacement(5); err != nil {
		return err
	}

	// Battleship
	if err := p.RandomizeShipPlacement(4); err != nil {
		return err
	}

	// Cruiser
	if err := p.RandomizeShipPlacement(3); err != nil {
		return err
	}

	// Submarine
	if err := p.RandomizeShipPlacement(3); err != nil {
		return err
	}

	// Destroyer
	if err := p.RandomizeShipPlacement(2); err != nil {
		return err
	}

	return nil
}

// Places 5 square ship on player_board. Errors if invalid placement.
func (p *Player) place_carrier(orientation cell.Orientation, coord cell.Coordinate) error {
	return p.place_ship(5, orientation, coord)
}

// Places 4 square ship on player_board. Errors if invalid placement.
func (p *Player) place_battleship(orientation cell.Orientation, coord cell.Coordinate) error {
	return p.place_ship(4, orientation, coord)
}

// Places 3 square ship on player_board. Errors if invalid placement.
func (p *Player) place_cruiser_or_submarine(orientation cell.Orientation, coord cell.Coordinate) error {
	return p.place_ship(3, orientation, coord)
}

// Places 2 square ship on player_board. Errors if invalid placement.
func (p *Player) place_destroyer(orientation cell.Orientation, coord cell.Coordinate) error {
	return p.place_ship(2, orientation, coord)
}

// Get's a player's coordinate guess checking against their
// previous turns (i.e. target_board)
func (p *Player) GetGuess() cell.Coordinate {
	return cell.GetRandomCoord(0)
}

// Check's if coordinate hits a ship in player_board
func (p *Player) CheckHit(coordinate cell.Coordinate) bool {
	return p.PlayerBoard[coordinate.ToIndex()].Occupied
}

// Mark an attempt on target_board.
func (p *Player) MarkTargetAttempt(coordinate cell.Coordinate, hit bool) {
	p.TargetBoard[coordinate.ToIndex()].Chosen = hit
}

// Mark an attempt on player_board
func (p *Player) MarkPlayerAttempt(coordinate cell.Coordinate, hit bool) {
	p.PlayerBoard[coordinate.ToIndex()].Chosen = hit
}
