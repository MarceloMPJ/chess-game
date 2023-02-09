package board_test

import (
	"testing"

	"github.com/MarceloMPJ/chess-game/libs/values"
	"github.com/MarceloMPJ/chess-game/pkg/board"
)

func TestBoard_Debug(t *testing.T) {

	b := board.NewBoard()
	b.Start()

	t.Run("when mode is GraficalMode", func(t *testing.T) {
		expected := `♖ ♘ ♗ ♕ ♔ ♗ ♘ ♖
♙ ♙ ♙ ♙ ♙ ♙ ♙ ♙
               
               
               
               
♟ ♟ ♟ ♟ ♟ ♟ ♟ ♟
♜ ♞ ♝ ♛ ♚ ♝ ♞ ♜
`
		result := b.Debug(board.GraphicalMode)

		if result != expected {
			t.Errorf("result: %s, expected: %s", result, expected)
		}
	})

	t.Run("when mode is FenMode", func(t *testing.T) {
		expected := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"
		result := b.Debug(board.FenMode)

		if result != expected {
			t.Errorf("result: %s, expected: %s", result, expected)
		}
	})
}

func TestBoard_Move(t *testing.T) {
	t.Run("when is valid move", func(t *testing.T) {
		b := board.NewBoard()
		b.Start()

		expectedBool := b.Move(values.Coord{X: 0, Y: 6}, values.Coord{X: 0, Y: 5})

		if !expectedBool {
			t.Errorf("result: %t, expected: %t", true, expectedBool)
		}

		expected := "rnbqkbnr/pppppppp/8/8/8/P7/1PPPPPPP/RNBQKBNR"
		result := b.Debug(board.FenMode)

		if result != expected {
			t.Errorf("result: %s, expected: %s", result, expected)
		}
	})

	t.Run("when is invalid move", func(t *testing.T) {
		t.Run("when has pieces between path and piece selected isn't knight", func(t *testing.T) {
		})
	})
}
