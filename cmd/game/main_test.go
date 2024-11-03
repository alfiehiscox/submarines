package main

import (
	"testing"
)

func TestPlaceCarrierSuccess(t *testing.T) {
	player := NewPlayer("test_player")
	orientation := HORIZONTAL
	coord := Coordinate{2, 2}

	if err := player.place_carrier(orientation, coord); err != nil {
		t.Fatalf("err should be nil: %s", err)
	}

	for i := 0; i < 5; i++ {
		x := coord[0] + i
		y := coord[1]
		cell := player.player_board[y*BOARD_WIDTH+x]
		if !cell.occupied {
			t.Fatalf("cell{%d,%d} was meant to be occupied", x, y)
		}
	}
}

func TestPlaceCarrierFailure(t *testing.T) {
	player := NewPlayer("test_player")
	if err := player.place_carrier(HORIZONTAL, Coordinate{0, 1}); err != nil {
		t.Fatal("failed in set up")
	}

	tests := []struct {
		orientation Orientation
		coord       Coordinate
	}{
		// Off the right side of the grid
		{orientation: HORIZONTAL, coord: Coordinate{BOARD_WIDTH + 2, 2}},
		{orientation: HORIZONTAL, coord: Coordinate{BOARD_WIDTH - 3, 5}},
		// minus x
		{orientation: HORIZONTAL, coord: Coordinate{-2, 4}},

		// Off the bottom of the grid
		{orientation: VERTICAL, coord: Coordinate{2, BOARD_HEIGHT + 2}},
		{orientation: VERTICAL, coord: Coordinate{5, BOARD_HEIGHT - 3}},
		// minus y
		{orientation: VERTICAL, coord: Coordinate{-2, 8}},

		// overlapping existing ship
		{orientation: VERTICAL, coord: Coordinate{0, 0}},
		{orientation: VERTICAL, coord: Coordinate{3, 1}},
		{orientation: HORIZONTAL, coord: Coordinate{1, 1}},
		{orientation: HORIZONTAL, coord: Coordinate{4, 1}},
	}

	for _, test := range tests {
		if err := player.place_carrier(test.orientation, test.coord); err == nil {
			t.Fatalf("expected error, got nil: %v", test)
		}
	}
}

func TestPlaceBattleshipSuccess(t *testing.T) {
	player := NewPlayer("test_player")
	orientation := HORIZONTAL
	coord := Coordinate{2, 2}

	if err := player.place_battleship(orientation, coord); err != nil {
		t.Fatalf("err should be nil: %s", err)
	}

	for i := 0; i < 4; i++ {
		x := coord[0] + i
		y := coord[1]
		cell := player.player_board[y*BOARD_WIDTH+x]
		if !cell.occupied {
			t.Fatalf("cell{%d,%d} was meant to be occupied", x, y)
		}
	}
}

func TestPlaceBattleshipFailure(t *testing.T) {

	player := NewPlayer("test_player")
	if err := player.place_carrier(HORIZONTAL, Coordinate{0, 1}); err != nil {
		t.Fatal("failed in set up")
	}

	tests := []struct {
		orientation Orientation
		coord       Coordinate
	}{
		// Off the right side of the grid
		{orientation: HORIZONTAL, coord: Coordinate{BOARD_WIDTH + 2, 2}},
		{orientation: HORIZONTAL, coord: Coordinate{BOARD_WIDTH - 2, 5}},
		// minus x
		{orientation: HORIZONTAL, coord: Coordinate{-2, 4}},

		// Off the bottom of the grid
		{orientation: VERTICAL, coord: Coordinate{2, BOARD_HEIGHT + 2}},
		{orientation: VERTICAL, coord: Coordinate{5, BOARD_HEIGHT - 2}},
		// minus y
		{orientation: VERTICAL, coord: Coordinate{-2, 8}},

		// overlapping existing ship
		{orientation: VERTICAL, coord: Coordinate{0, 0}},
		{orientation: VERTICAL, coord: Coordinate{3, 1}},
		{orientation: HORIZONTAL, coord: Coordinate{1, 1}},
		{orientation: HORIZONTAL, coord: Coordinate{3, 1}},
	}

	for _, test := range tests {
		if err := player.place_battleship(test.orientation, test.coord); err == nil {
			t.Fatalf("expected error, got nil: %v", test)
		}
	}
}

func TestPlaceCruiserOrSubmarineSuccess(t *testing.T) {
	player := NewPlayer("test_player")
	orientation := HORIZONTAL
	coord := Coordinate{2, 2}

	if err := player.place_cruiser_or_submarine(orientation, coord); err != nil {
		t.Fatalf("err should be nil: %s", err)
	}

	for i := 0; i < 3; i++ {
		x := coord[0] + i
		y := coord[1]
		cell := player.player_board[y*BOARD_WIDTH+x]
		if !cell.occupied {
			t.Fatalf("cell{%d,%d} was meant to be occupied", x, y)
		}
	}
}

func TestPlaceSubmarineFailure(t *testing.T) {

	player := NewPlayer("test_player")
	if err := player.place_cruiser_or_submarine(HORIZONTAL, Coordinate{0, 1}); err != nil {
		t.Fatal("failed in set up")
	}

	tests := []struct {
		orientation Orientation
		coord       Coordinate
	}{
		// Off the right side of the grid
		{orientation: HORIZONTAL, coord: Coordinate{BOARD_WIDTH + 2, 2}},
		{orientation: HORIZONTAL, coord: Coordinate{BOARD_WIDTH - 1, 5}},
		// minus x
		{orientation: HORIZONTAL, coord: Coordinate{-2, 4}},

		// Off the bottom of the grid
		{orientation: VERTICAL, coord: Coordinate{2, BOARD_HEIGHT + 2}},
		{orientation: VERTICAL, coord: Coordinate{5, BOARD_HEIGHT - 1}},
		// minus y
		{orientation: VERTICAL, coord: Coordinate{-2, 8}},

		// overlapping existing ship
		{orientation: VERTICAL, coord: Coordinate{0, 0}},
		{orientation: VERTICAL, coord: Coordinate{2, 1}},
		{orientation: HORIZONTAL, coord: Coordinate{1, 1}},
		{orientation: HORIZONTAL, coord: Coordinate{2, 1}},
	}

	for _, test := range tests {
		if err := player.place_cruiser_or_submarine(test.orientation, test.coord); err == nil {
			t.Fatalf("expected error, got nil: %v", test)
		}
	}
}

func TestPlaceDestroyerSuccess(t *testing.T) {
	player := NewPlayer("test_player")
	orientation := HORIZONTAL
	coord := Coordinate{2, 2}

	if err := player.place_destroyer(orientation, coord); err != nil {
		t.Fatalf("err should be nil: %s", err)
	}

	for i := 0; i < 2; i++ {
		x := coord[0] + i
		y := coord[1]
		cell := player.player_board[y*BOARD_WIDTH+x]
		if !cell.occupied {
			t.Fatalf("cell{%d,%d} was meant to be occupied", x, y)
		}
	}
}

func TestPlaceDestroyerFailure(t *testing.T) {

	player := NewPlayer("test_player")
	if err := player.place_destroyer(HORIZONTAL, Coordinate{0, 1}); err != nil {
		t.Fatal("failed in set up")
	}

	tests := []struct {
		orientation Orientation
		coord       Coordinate
	}{
		// Off the right side of the grid
		{orientation: HORIZONTAL, coord: Coordinate{BOARD_WIDTH + 2, 2}},
		{orientation: HORIZONTAL, coord: Coordinate{BOARD_WIDTH, 5}},
		// minus x
		{orientation: HORIZONTAL, coord: Coordinate{-2, 4}},

		// Off the bottom of the grid
		{orientation: VERTICAL, coord: Coordinate{2, BOARD_HEIGHT + 2}},
		{orientation: VERTICAL, coord: Coordinate{5, BOARD_HEIGHT}},
		// minus y
		{orientation: VERTICAL, coord: Coordinate{-2, 8}},

		// overlapping existing ship
		{orientation: VERTICAL, coord: Coordinate{0, 0}},
		{orientation: VERTICAL, coord: Coordinate{1, 1}},
		{orientation: HORIZONTAL, coord: Coordinate{1, 1}},
	}

	for _, test := range tests {
		if err := player.place_destroyer(test.orientation, test.coord); err == nil {
			t.Fatalf("expected error, got nil: %v", test)
		}
	}
}

func TestRandomizePlacementSuccess(t *testing.T) {
	player := NewPlayer("test_player")

	t.Log(player.player_board)

	if err := player.randomize_placement(); err != nil {
		t.Fatalf("err should be nil: %s", err)
	}

	c := 0
	for _, cell := range player.player_board {
		if cell.occupied {
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
	for i := range player.player_board {

		if i%BOARD_WIDTH-1 == 0 {
			paint_row = !paint_row
		}

		paint_row = !paint_row

		if paint_row {
			player.player_board[i].occupied = true
		}

	}

	if err := player.randomize_placement(); err == nil {
		t.Fatalf("expected error, got nil")
	}
}

// TODO: Get Guess (probs needs DI)
func TestGetGuess(_ *testing.T) {}

func TestCheckHit(t *testing.T) {
	p := NewPlayer("test_player")

	if err := p.place_carrier(HORIZONTAL, Coordinate{0, 0}); err != nil {
		t.Fatal(err)
	}

	if err := p.place_battleship(VERTICAL, Coordinate{0, 1}); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		input    Coordinate
		expected bool
	}{
		// True
		{Coordinate{0, 0}, true},
		{Coordinate{1, 0}, true},
		{Coordinate{2, 0}, true},
		{Coordinate{3, 0}, true},
		{Coordinate{4, 0}, true},
		{Coordinate{0, 1}, true},
		{Coordinate{0, 2}, true},
		{Coordinate{0, 3}, true},
		{Coordinate{0, 4}, true},

		// False
		{Coordinate{5, 0}, false},
		{Coordinate{6, 0}, false},
		{Coordinate{7, 0}, false},
		{Coordinate{8, 0}, false},
		{Coordinate{0, 5}, false},
		{Coordinate{9, 6}, false},
		{Coordinate{3, 3}, false},
		{Coordinate{8, 3}, false},
		{Coordinate{2, 9}, false},
	}

	for _, test := range tests {
		actual := p.check_hit(test.input)
		if actual != test.expected {
			t.Fatalf("%v :: Exp=%v, Act=%v", test.input, test.expected, actual)
		}
	}
}

func TestMarkTargetAttempt(t *testing.T) {
	p := NewPlayer("test_player")
	coord, err := NewCoordinate(2, 5)
	if err != nil {
		t.Fatal(err)
	}

	if p.target_board[coord.ToIndex()].chosen {
		t.Fatalf("Expected %v to start an unoccupied", coord)
	}

	p.mark_target_attempt(coord, true)
	if !p.target_board[coord.ToIndex()].chosen {
		t.Fatalf("Expected %v to then be occupied", coord)
	}

	p.mark_target_attempt(coord, false)
	if p.target_board[coord.ToIndex()].chosen {
		t.Fatalf("Expected %v to then be unoccupied", coord)
	}
}

func TestMarkPlayerAttempt(t *testing.T) {
	p := NewPlayer("test_player")
	coord, err := NewCoordinate(1, 8)
	if err != nil {
		t.Fatal(err)
	}

	if p.player_board[coord.ToIndex()].chosen {
		t.Fatalf("Expected %v to start an unoccupied", coord)
	}

	p.mark_player_attempt(coord, true)
	if !p.player_board[coord.ToIndex()].chosen {
		t.Fatalf("Expected %v to then be occupied", coord)
	}

	p.mark_player_attempt(coord, false)
	if p.player_board[coord.ToIndex()].chosen {
		t.Fatalf("Expected %v to then be unoccupied", coord)
	}
}

func TestCheckWinnerSuccess(t *testing.T) {
	p1 := NewPlayer("test_player_1")
	if err := p1.place_carrier(HORIZONTAL, Coordinate{0, 0}); err != nil {
		t.Fatal(err)
	}
	if err := p1.place_battleship(VERTICAL, Coordinate{0, 1}); err != nil {
		t.Fatal(err)
	}
	t.Log(p1.player_board.String())

	p2 := NewPlayer("test_player_1")
	p2.mark_target_attempt(Coordinate{0, 0}, true)
	p2.mark_target_attempt(Coordinate{1, 0}, true)
	p2.mark_target_attempt(Coordinate{2, 0}, true)
	p2.mark_target_attempt(Coordinate{3, 0}, true)
	p2.mark_target_attempt(Coordinate{4, 0}, true)
	p2.mark_target_attempt(Coordinate{0, 1}, true)
	p2.mark_target_attempt(Coordinate{0, 2}, true)
	p2.mark_target_attempt(Coordinate{0, 3}, true)
	p2.mark_target_attempt(Coordinate{0, 4}, true)

	if !check_winner(p2.target_board, p1.player_board) {
		t.Fatalf("Should have been won")
	}
}

func TestCheckWinnerFailure(t *testing.T) {
	p1 := NewPlayer("test_player_1")
	if err := p1.place_carrier(HORIZONTAL, Coordinate{0, 1}); err != nil {
		t.Fatal(err)
	}
	if err := p1.place_battleship(VERTICAL, Coordinate{3, 4}); err != nil {
		t.Fatal(err)
	}
	t.Log(p1.player_board.String())

	p2 := NewPlayer("test_player_1")
	p2.mark_target_attempt(Coordinate{0, 0}, true)
	p2.mark_target_attempt(Coordinate{1, 0}, true)
	p2.mark_target_attempt(Coordinate{2, 0}, true)
	p2.mark_target_attempt(Coordinate{3, 0}, true)
	p2.mark_target_attempt(Coordinate{4, 0}, true)
	p2.mark_target_attempt(Coordinate{0, 1}, true)
	p2.mark_target_attempt(Coordinate{0, 2}, true)
	p2.mark_target_attempt(Coordinate{0, 3}, true)
	p2.mark_target_attempt(Coordinate{0, 4}, true)

	if check_winner(p2.target_board, p1.player_board) {
		t.Fatalf("Should NOT have been won")
	}
}
