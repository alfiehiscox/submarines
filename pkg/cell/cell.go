package cell

import (
	"errors"
	"fmt"
	"math/rand/v2"
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

func (c Coordinate) ToIndex() int {
	return c[1]*BOARD_WIDTH + c[0]
}

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
	Occupied bool
	Chosen   bool
}

func GetRandomOrientation() Orientation {
	if rand.IntN(2) == 0 {
		return HORIZONTAL
	} else {
		return VERTICAL
	}
}

func GetRandomCoord(size int) Coordinate {
	x := rand.IntN(BOARD_WIDTH - size)
	y := rand.IntN(BOARD_HEIGHT - size)
	return Coordinate{x, y}
}

func VerifyCoordinate(size int, orientation Orientation, coord Coordinate) error {
	if coord[0] < 0 || coord[1] < 0 {
		msg := fmt.Sprintf("Carrier at %v [%s] is off the board", coord, orientation)
		return errors.New(msg)
	}

	if coord[0] > BOARD_WIDTH || coord[1] > BOARD_HEIGHT {
		msg := fmt.Sprintf("Carrier at %v [%s] is off the board", coord, orientation)
		return errors.New(msg)
	}

	if orientation == HORIZONTAL && coord[0] > BOARD_WIDTH-size {
		msg := fmt.Sprintf("Carrier at %v [%s] is off the board", coord, orientation)
		return errors.New(msg)
	}

	if orientation == VERTICAL && coord[1] > BOARD_HEIGHT-size {
		msg := fmt.Sprintf("Carrier at %v [%s] is off the board", coord, orientation)
		return errors.New(msg)
	}

	return nil
}
