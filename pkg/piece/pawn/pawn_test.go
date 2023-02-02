package pawn_test

import (
	"testing"

	"github.com/MarceloMPJ/chess-game/libs/values"
	"github.com/MarceloMPJ/chess-game/pkg/piece/pawn"
)

func TestPawn_Show(t *testing.T) {
	t.Run("when color is white", func(t *testing.T) {
		p := pawn.NewPawn(values.White)

		expected := '♟'
		result := p.Show()

		if result != expected {
			t.Errorf("result: %c, expected: %c", result, expected)
		}
	})

	t.Run("when color is black", func(t *testing.T) {
		p := pawn.NewPawn(values.Black)

		expected := '♙'
		result := p.Show()

		if result != expected {
			t.Errorf("result: %c, expected: %c", result, expected)
		}
	})
}

func TestPawn_IsValidMove(t *testing.T) {
	t.Run("when color is white", func(t *testing.T) {
		p := pawn.NewPawn(values.White)

		t.Run("when the move is valid", func(t *testing.T) {
			expected := true

			result := p.IsValidMove(
				values.Coord{X: 7, Y: 7},
				values.Coord{X: 7, Y: 6},
			)
			if result != expected {
				t.Errorf("result: %t, expected: %t", result, expected)
			}

			result = p.IsValidMove(
				values.Coord{X: 7, Y: 7},
				values.Coord{X: 7, Y: 5},
			)

			if result != expected {
				t.Errorf("result: %t, expected: %t", result, expected)
			}
		})

		t.Run("when the move is invalid", func(t *testing.T) {
			expected := false

			result := p.IsValidMove(
				values.Coord{X: 7, Y: 7},
				values.Coord{X: 6, Y: 6},
			)
			if result != expected {
				t.Errorf("result: %t, expected: %t", result, expected)
			}

			result = p.IsValidMove(
				values.Coord{X: 7, Y: 7},
				values.Coord{X: 7, Y: 4},
			)
			if result != expected {
				t.Errorf("result: %t, expected: %t", result, expected)
			}

			result = p.IsValidMove(
				values.Coord{X: 6, Y: 6},
				values.Coord{X: 6, Y: 7},
			)
			if result != expected {
				t.Errorf("result: %t, expected: %t", result, expected)
			}
		})
	})

	t.Run("when color is black", func(t *testing.T) {
		p := pawn.NewPawn(values.Black)

		t.Run("when the move is valid", func(t *testing.T) {
			expected := true

			result := p.IsValidMove(
				values.Coord{X: 7, Y: 6},
				values.Coord{X: 7, Y: 7},
			)
			if result != expected {
				t.Errorf("result: %t, expected: %t", result, expected)
			}

			result = p.IsValidMove(
				values.Coord{X: 7, Y: 5},
				values.Coord{X: 7, Y: 7},
			)

			if result != expected {
				t.Errorf("result: %t, expected: %t", result, expected)
			}
		})

		t.Run("when the move is invalid", func(t *testing.T) {
			expected := false

			result := p.IsValidMove(
				values.Coord{X: 6, Y: 6},
				values.Coord{X: 7, Y: 7},
			)
			if result != expected {
				t.Errorf("result: %t, expected: %t", result, expected)
			}

			result = p.IsValidMove(
				values.Coord{X: 7, Y: 4},
				values.Coord{X: 7, Y: 7},
			)
			if result != expected {
				t.Errorf("result: %t, expected: %t", result, expected)
			}

			result = p.IsValidMove(
				values.Coord{X: 6, Y: 7},
				values.Coord{X: 6, Y: 6},
			)
			if result != expected {
				t.Errorf("result: %t, expected: %t", result, expected)
			}
		})
	})
}
