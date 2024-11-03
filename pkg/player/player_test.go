package player

import (
	"testing"

	"github.com/alfiehiscox/submarines/pkg/board"
	"github.com/alfiehiscox/submarines/pkg/cell"
)

func TestPlaceCarrierSuccess(t *testing.T) {
	player := NewPlayer("test_player")
	orientation := cell.HORIZONTAL
	coord := cell.Coordinate{2, 2}

	if err := player.place_carrier(orientation, coord); err != nil {
		t.Fatalf("err should be nil: %s", err)
	}

	for i := 0; i < 5; i++ {
		x := coord[0] + i
		y := coord[1]
		cell := player.PlayerBoard[y*cell.BOARD_WIDTH+x]
		if !cell.Occupied {
			t.Fatalf("cell{%d,%d} was meant to be occupied", x, y)
		}
	}
}

func TestPlaceCarrierFailure(t *testing.T) {
	player := NewPlayer("test_player")
	if err := player.place_carrier(cell.HORIZONTAL, cell.Coordinate{0, 1}); err != nil {
		t.Fatal("failed in set up")
	}

	tests := []struct {
		orientation cell.Orientation
		coord       cell.Coordinate
	}{
		// Off the right side of the grid
		{orientation: cell.HORIZONTAL, coord: cell.Coordinate{cell.BOARD_WIDTH + 2, 2}},
		{orientation: cell.HORIZONTAL, coord: cell.Coordinate{cell.BOARD_WIDTH - 3, 5}},
		// minus x
		{orientation: cell.HORIZONTAL, coord: cell.Coordinate{-2, 4}},

		// Off the bottom of the grid
		{orientation: cell.VERTICAL, coord: cell.Coordinate{2, cell.BOARD_HEIGHT + 2}},
		{orientation: cell.VERTICAL, coord: cell.Coordinate{5, cell.BOARD_HEIGHT - 3}},
		// minus y
		{orientation: cell.VERTICAL, coord: cell.Coordinate{-2, 8}},

		// overlapping existing ship
		{orientation: cell.VERTICAL, coord: cell.Coordinate{0, 0}},
		{orientation: cell.VERTICAL, coord: cell.Coordinate{3, 1}},
		{orientation: cell.HORIZONTAL, coord: cell.Coordinate{1, 1}},
		{orientation: cell.HORIZONTAL, coord: cell.Coordinate{4, 1}},
	}

	for _, test := range tests {
		if err := player.place_carrier(test.orientation, test.coord); err == nil {
			t.Fatalf("expected error, got nil: %v", test)
		}
	}
}

func TestPlaceBattleshipSuccess(t *testing.T) {
	player := NewPlayer("test_player")
	orientation := cell.HORIZONTAL
	coord := cell.Coordinate{2, 2}

	if err := player.place_battleship(orientation, coord); err != nil {
		t.Fatalf("err should be nil: %s", err)
	}

	for i := 0; i < 4; i++ {
		x := coord[0] + i
		y := coord[1]
		cell := player.PlayerBoard[y*cell.BOARD_WIDTH+x]
		if !cell.Occupied {
			t.Fatalf("cell{%d,%d} was meant to be occupied", x, y)
		}
	}
}

func TestPlaceBattleshipFailure(t *testing.T) {

	player := NewPlayer("test_player")
	if err := player.place_carrier(cell.HORIZONTAL, cell.Coordinate{0, 1}); err != nil {
		t.Fatal("failed in set up")
	}

	tests := []struct {
		orientation cell.Orientation
		coord       cell.Coordinate
	}{
		// Off the right side of the grid
		{orientation: cell.HORIZONTAL, coord: cell.Coordinate{cell.BOARD_WIDTH + 2, 2}},
		{orientation: cell.HORIZONTAL, coord: cell.Coordinate{cell.BOARD_WIDTH - 2, 5}},
		// minus x
		{orientation: cell.HORIZONTAL, coord: cell.Coordinate{-2, 4}},

		// Off the bottom of the grid
		{orientation: cell.VERTICAL, coord: cell.Coordinate{2, cell.BOARD_HEIGHT + 2}},
		{orientation: cell.VERTICAL, coord: cell.Coordinate{5, cell.BOARD_HEIGHT - 2}},
		// minus y
		{orientation: cell.VERTICAL, coord: cell.Coordinate{-2, 8}},

		// overlapping existing ship
		{orientation: cell.VERTICAL, coord: cell.Coordinate{0, 0}},
		{orientation: cell.VERTICAL, coord: cell.Coordinate{3, 1}},
		{orientation: cell.HORIZONTAL, coord: cell.Coordinate{1, 1}},
		{orientation: cell.HORIZONTAL, coord: cell.Coordinate{3, 1}},
	}

	for _, test := range tests {
		if err := player.place_battleship(test.orientation, test.coord); err == nil {
			t.Fatalf("expected error, got nil: %v", test)
		}
	}
}

func TestPlaceCruiserOrSubmarineSuccess(t *testing.T) {
	player := NewPlayer("test_player")
	orientation := cell.HORIZONTAL
	coord := cell.Coordinate{2, 2}

	if err := player.place_cruiser_or_submarine(orientation, coord); err != nil {
		t.Fatalf("err should be nil: %s", err)
	}

	for i := 0; i < 3; i++ {
		x := coord[0] + i
		y := coord[1]
		cell := player.PlayerBoard[y*cell.BOARD_WIDTH+x]
		if !cell.Occupied {
			t.Fatalf("cell{%d,%d} was meant to be occupied", x, y)
		}
	}
}

func TestPlaceSubmarineFailure(t *testing.T) {

	player := NewPlayer("test_player")
	if err := player.place_cruiser_or_submarine(cell.HORIZONTAL, cell.Coordinate{0, 1}); err != nil {
		t.Fatal("failed in set up")
	}

	tests := []struct {
		orientation cell.Orientation
		coord       cell.Coordinate
	}{
		// Off the right side of the grid
		{orientation: cell.HORIZONTAL, coord: cell.Coordinate{cell.BOARD_WIDTH + 2, 2}},
		{orientation: cell.HORIZONTAL, coord: cell.Coordinate{cell.BOARD_WIDTH - 1, 5}},
		// minus x
		{orientation: cell.HORIZONTAL, coord: cell.Coordinate{-2, 4}},

		// Off the bottom of the grid
		{orientation: cell.VERTICAL, coord: cell.Coordinate{2, cell.BOARD_HEIGHT + 2}},
		{orientation: cell.VERTICAL, coord: cell.Coordinate{5, cell.BOARD_HEIGHT - 1}},
		// minus y
		{orientation: cell.VERTICAL, coord: cell.Coordinate{-2, 8}},

		// overlapping existing ship
		{orientation: cell.VERTICAL, coord: cell.Coordinate{0, 0}},
		{orientation: cell.VERTICAL, coord: cell.Coordinate{2, 1}},
		{orientation: cell.HORIZONTAL, coord: cell.Coordinate{1, 1}},
		{orientation: cell.HORIZONTAL, coord: cell.Coordinate{2, 1}},
	}

	for _, test := range tests {
		if err := player.place_cruiser_or_submarine(test.orientation, test.coord); err == nil {
			t.Fatalf("expected error, got nil: %v", test)
		}
	}
}

func TestPlaceDestroyerSuccess(t *testing.T) {
	player := NewPlayer("test_player")
	orientation := cell.HORIZONTAL
	coord := cell.Coordinate{2, 2}

	if err := player.place_destroyer(orientation, coord); err != nil {
		t.Fatalf("err should be nil: %s", err)
	}

	for i := 0; i < 2; i++ {
		x := coord[0] + i
		y := coord[1]
		cell := player.PlayerBoard[y*cell.BOARD_WIDTH+x]
		if !cell.Occupied {
			t.Fatalf("cell{%d,%d} was meant to be occupied", x, y)
		}
	}
}

func TestPlaceDestroyerFailure(t *testing.T) {

	player := NewPlayer("test_player")
	if err := player.place_destroyer(cell.HORIZONTAL, cell.Coordinate{0, 1}); err != nil {
		t.Fatal("failed in set up")
	}

	tests := []struct {
		orientation cell.Orientation
		coord       cell.Coordinate
	}{
		// Off the right side of the grid
		{orientation: cell.HORIZONTAL, coord: cell.Coordinate{cell.BOARD_WIDTH + 2, 2}},
		{orientation: cell.HORIZONTAL, coord: cell.Coordinate{cell.BOARD_WIDTH, 5}},
		// minus x
		{orientation: cell.HORIZONTAL, coord: cell.Coordinate{-2, 4}},

		// Off the bottom of the grid
		{orientation: cell.VERTICAL, coord: cell.Coordinate{2, cell.BOARD_HEIGHT + 2}},
		{orientation: cell.VERTICAL, coord: cell.Coordinate{5, cell.BOARD_HEIGHT}},
		// minus y
		{orientation: cell.VERTICAL, coord: cell.Coordinate{-2, 8}},

		// overlapping existing ship
		{orientation: cell.VERTICAL, coord: cell.Coordinate{0, 0}},
		{orientation: cell.VERTICAL, coord: cell.Coordinate{1, 1}},
		{orientation: cell.HORIZONTAL, coord: cell.Coordinate{1, 1}},
	}

	for _, test := range tests {
		if err := player.place_destroyer(test.orientation, test.coord); err == nil {
			t.Fatalf("expected error, got nil: %v", test)
		}
	}
}

func TestRandomizePlacementSuccess(t *testing.T) {
	player := NewPlayer("test_player")

	t.Log(player.PlayerBoard)

	if err := player.RandomizePlacement(); err != nil {
		t.Fatalf("err should be nil: %s", err)
	}

	c := 0
	for _, cell := range player.PlayerBoard {
		if cell.Occupied {
			c += 1
		}
	}

	if c != 17 {
		t.Fatalf("Expected 17 squares to be filled, got=%d", c)
	}
}

func TestRandomizePlacementLimit(t *testing.T) {
	player := NewPlayer("test_player")

	paint_row := true
	for i := range player.PlayerBoard {

		if i%cell.BOARD_WIDTH-1 == 0 {
			paint_row = !paint_row
		}

		paint_row = !paint_row

		if paint_row {
			player.PlayerBoard[i].Occupied = true
		}

	}

	if err := player.RandomizePlacement(); err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestCheckHit(t *testing.T) {
	p := NewPlayer("test_player")

	if err := p.place_carrier(cell.HORIZONTAL, cell.Coordinate{0, 0}); err != nil {
		t.Fatal(err)
	}

	if err := p.place_battleship(cell.VERTICAL, cell.Coordinate{0, 1}); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		input    cell.Coordinate
		expected bool
	}{
		// True
		{cell.Coordinate{0, 0}, true},
		{cell.Coordinate{1, 0}, true},
		{cell.Coordinate{2, 0}, true},
		{cell.Coordinate{3, 0}, true},
		{cell.Coordinate{4, 0}, true},
		{cell.Coordinate{0, 1}, true},
		{cell.Coordinate{0, 2}, true},
		{cell.Coordinate{0, 3}, true},
		{cell.Coordinate{0, 4}, true},

		// False
		{cell.Coordinate{5, 0}, false},
		{cell.Coordinate{6, 0}, false},
		{cell.Coordinate{7, 0}, false},
		{cell.Coordinate{8, 0}, false},
		{cell.Coordinate{0, 5}, false},
		{cell.Coordinate{9, 6}, false},
		{cell.Coordinate{3, 3}, false},
		{cell.Coordinate{8, 3}, false},
		{cell.Coordinate{2, 9}, false},
	}

	for _, test := range tests {
		actual := p.CheckHit(test.input)
		if actual != test.expected {
			t.Fatalf("%v :: Exp=%v, Act=%v", test.input, test.expected, actual)
		}
	}
}

func TestMarkTargetAttempt(t *testing.T) {
	p := NewPlayer("test_player")
	coord, err := cell.NewCoordinate(2, 5)
	if err != nil {
		t.Fatal(err)
	}

	if p.TargetBoard[coord.ToIndex()].Chosen {
		t.Fatalf("Expected %v to start an unoccupied", coord)
	}

	p.MarkTargetAttempt(coord, true)
	if !p.TargetBoard[coord.ToIndex()].Chosen {
		t.Fatalf("Expected %v to then be occupied", coord)
	}

	p.MarkTargetAttempt(coord, false)
	if p.TargetBoard[coord.ToIndex()].Chosen {
		t.Fatalf("Expected %v to then be unoccupied", coord)
	}
}

func TestMarkPlayerAttempt(t *testing.T) {
	p := NewPlayer("test_player")
	coord, err := cell.NewCoordinate(1, 8)
	if err != nil {
		t.Fatal(err)
	}

	if p.PlayerBoard[coord.ToIndex()].Chosen {
		t.Fatalf("Expected %v to start an unoccupied", coord)
	}

	p.MarkPlayerAttempt(coord, true)
	if !p.PlayerBoard[coord.ToIndex()].Chosen {
		t.Fatalf("Expected %v to then be occupied", coord)
	}

	p.MarkPlayerAttempt(coord, false)
	if p.PlayerBoard[coord.ToIndex()].Chosen {
		t.Fatalf("Expected %v to then be unoccupied", coord)
	}
}

func TestCheckWinnerSuccess(t *testing.T) {
	p1 := NewPlayer("test_player_1")
	if err := p1.place_carrier(cell.HORIZONTAL, cell.Coordinate{0, 0}); err != nil {
		t.Fatal(err)
	}
	if err := p1.place_battleship(cell.VERTICAL, cell.Coordinate{0, 1}); err != nil {
		t.Fatal(err)
	}
	t.Log(p1.PlayerBoard.String())

	p2 := NewPlayer("test_player_1")
	p2.MarkTargetAttempt(cell.Coordinate{0, 0}, true)
	p2.MarkTargetAttempt(cell.Coordinate{1, 0}, true)
	p2.MarkTargetAttempt(cell.Coordinate{2, 0}, true)
	p2.MarkTargetAttempt(cell.Coordinate{3, 0}, true)
	p2.MarkTargetAttempt(cell.Coordinate{4, 0}, true)
	p2.MarkTargetAttempt(cell.Coordinate{0, 1}, true)
	p2.MarkTargetAttempt(cell.Coordinate{0, 2}, true)
	p2.MarkTargetAttempt(cell.Coordinate{0, 3}, true)
	p2.MarkTargetAttempt(cell.Coordinate{0, 4}, true)

	if !board.CheckWinner(p2.TargetBoard, p1.PlayerBoard) {
		t.Fatalf("Should have been won")
	}
}

func TestCheckWinnerFailure(t *testing.T) {
	p1 := NewPlayer("test_player_1")
	if err := p1.place_carrier(cell.HORIZONTAL, cell.Coordinate{0, 1}); err != nil {
		t.Fatal(err)
	}
	if err := p1.place_battleship(cell.VERTICAL, cell.Coordinate{3, 4}); err != nil {
		t.Fatal(err)
	}
	t.Log(p1.PlayerBoard.String())

	p2 := NewPlayer("test_player_1")
	p2.MarkTargetAttempt(cell.Coordinate{0, 0}, true)
	p2.MarkTargetAttempt(cell.Coordinate{1, 0}, true)
	p2.MarkTargetAttempt(cell.Coordinate{2, 0}, true)
	p2.MarkTargetAttempt(cell.Coordinate{3, 0}, true)
	p2.MarkTargetAttempt(cell.Coordinate{4, 0}, true)
	p2.MarkTargetAttempt(cell.Coordinate{0, 1}, true)
	p2.MarkTargetAttempt(cell.Coordinate{0, 2}, true)
	p2.MarkTargetAttempt(cell.Coordinate{0, 3}, true)
	p2.MarkTargetAttempt(cell.Coordinate{0, 4}, true)

	if board.CheckWinner(p2.TargetBoard, p1.PlayerBoard) {
		t.Fatalf("Should NOT have been won")
	}
}
