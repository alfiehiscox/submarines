package board

import (
	"strings"

	"github.com/alfiehiscox/submarines/pkg/cell"
)

type Board []cell.Cell

func NewBoard() Board {
	board := make(Board, cell.BOARD_HEIGHT*cell.BOARD_WIDTH)
	for i := range board {
		board[i].Chosen = false
		board[i].Occupied = false
	}
	return board
}

func (b Board) String() string {
	builder := strings.Builder{}
	builder.WriteString("\n")

	count := 1
	for i := range b {
		if b[i].Occupied {
			builder.WriteString(" X ")
		} else {
			builder.WriteString(" O ")
		}

		if count == cell.BOARD_WIDTH {
			builder.WriteString("\n")
			count = 1
		} else {
			count += 1
		}
	}

	return builder.String()
}

// checks if the target_board has hit all ships on enemy_board
func CheckWinner(target_board, enemy_board Board) bool {
	for i := range enemy_board {
		if enemy_board[i].Occupied != target_board[i].Chosen {
			return false
		}
	}

	return true
}
