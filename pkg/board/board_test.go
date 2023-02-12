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
		expectedBool := true

		b := board.NewBoard()
		b.Start()

		// Moves pawn on initial board
		t.Run("when moves pawn white on initial board", func(t *testing.T) {
			expectedFen := "rnbqkbnr/pppppppp/8/8/8/3P4/PPP1PPPP/RNBQKBNR"

			resultBool := b.Move(values.Coord{X: 3, Y: 6}, values.Coord{X: 3, Y: 5})
			resultFen := b.Debug(board.FenMode)

			checkMove(t, resultBool, expectedBool, resultFen, expectedFen)
		})

		t.Run("when moves pawn black", func(t *testing.T) {
			expectedFen := "rnbqkbnr/ppp1pppp/3p4/8/8/3P4/PPP1PPPP/RNBQKBNR"

			resultBool := b.Move(values.Coord{X: 3, Y: 1}, values.Coord{X: 3, Y: 2})
			resultFen := b.Debug(board.FenMode)

			checkMove(t, resultBool, expectedBool, resultFen, expectedFen)
		})

		t.Run("when moves bishop white", func(t *testing.T) {
			expectedFen := "rnbqkbnr/ppp1pppp/3p4/6B1/8/3P4/PPP1PPPP/RN1QKBNR"

			resultBool := b.Move(values.Coord{X: 2, Y: 7}, values.Coord{X: 6, Y: 3})
			resultFen := b.Debug(board.FenMode)

			checkMove(t, resultBool, expectedBool, resultFen, expectedFen)
		})

		t.Run("when moves bishop black", func(t *testing.T) {
			expectedFen := "rn1qkbnr/ppp1pppp/3p4/6B1/6b1/3P4/PPP1PPPP/RN1QKBNR"

			resultBool := b.Move(values.Coord{X: 2, Y: 0}, values.Coord{X: 6, Y: 4})
			resultFen := b.Debug(board.FenMode)

			checkMove(t, resultBool, expectedBool, resultFen, expectedFen)
		})

		t.Run("whem moves knight white", func(t *testing.T) {
			expectedFen := "rn1qkbnr/ppp1pppp/3p4/6B1/6b1/3P1N2/PPP1PPPP/RN1QKB1R"

			resultBool := b.Move(values.Coord{X: 6, Y: 7}, values.Coord{X: 5, Y: 5})
			resultFen := b.Debug(board.FenMode)

			checkMove(t, resultBool, expectedBool, resultFen, expectedFen)
		})

		t.Run("whem moves knight black", func(t *testing.T) {
			expectedFen := "rn1qkb1r/ppp1pppp/3p1n2/6B1/6b1/3P1N2/PPP1PPPP/RN1QKB1R"

			resultBool := b.Move(values.Coord{X: 6, Y: 0}, values.Coord{X: 5, Y: 2})
			resultFen := b.Debug(board.FenMode)

			checkMove(t, resultBool, expectedBool, resultFen, expectedFen)
		})
	})

	t.Run("when is invalid move", func(t *testing.T) {
		expectedBool := false
		expectedFen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"

		t.Run("when has pieces between path and piece selected isn't knight", func(t *testing.T) {
			b := board.NewBoard()
			b.Start()

			// Moves rook on initial board
			t.Run("when move rook on initial board", func(t *testing.T) {
				resultBool := b.Move(values.Coord{X: 0, Y: 7}, values.Coord{X: 0, Y: 5})
				resultFen := b.Debug(board.FenMode)

				checkMove(t, resultBool, expectedBool, resultFen, expectedFen)
			})

			t.Run("when move bishop on initial board", func(t *testing.T) {
				resultBool := b.Move(values.Coord{X: 2, Y: 7}, values.Coord{X: 6, Y: 3})
				resultFen := b.Debug(board.FenMode)

				checkMove(t, resultBool, expectedBool, resultFen, expectedFen)
			})
		})
	})
}

func checkMove(t *testing.T, resultBool, expectedBool bool, resultFen, expectedFen string) {
	t.Helper()

	if resultBool != expectedBool {
		t.Fatalf("result: %t, expected: %t", resultBool, expectedBool)
	}

	if resultFen != expectedFen {
		t.Errorf("result: %s, expected: %s", resultFen, expectedFen)
	}
}
