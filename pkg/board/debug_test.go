package board_test

import (
	"testing"

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
