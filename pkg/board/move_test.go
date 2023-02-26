package board_test

import (
	"testing"

	"github.com/MarceloMPJ/chess-game/libs/values"
	"github.com/MarceloMPJ/chess-game/pkg/board"
)

func TestBoard_Move(t *testing.T) {
	t.Run("when is valid move", func(t *testing.T) {
		expectedBool := true

		b := board.NewBoard()
		b.Start()

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

		t.Run("when moves knight white", func(t *testing.T) {
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

		t.Run("when bishop white captures knight black", func(t *testing.T) {
			expectedFen := "rn1qkb1r/ppp1pppp/3p1B2/8/6b1/3P1N2/PPP1PPPP/RN1QKB1R"

			resultBool := b.Move(values.Coord{X: 6, Y: 3}, values.Coord{X: 5, Y: 2})
			resultFen := b.Debug(board.FenMode)

			checkMove(t, resultBool, expectedBool, resultFen, expectedFen)
		})

		t.Run("when bishop black captures knight white", func(t *testing.T) {
			expectedFen := "rn1qkb1r/ppp1pppp/3p1B2/8/8/3P1b2/PPP1PPPP/RN1QKB1R"

			resultBool := b.Move(values.Coord{X: 6, Y: 4}, values.Coord{X: 5, Y: 5})
			resultFen := b.Debug(board.FenMode)

			checkMove(t, resultBool, expectedBool, resultFen, expectedFen)
		})

		t.Run("when pawn white captures bishop black", func(t *testing.T) {
			expectedFen := "rn1qkb1r/ppp1pppp/3p1B2/8/8/3P1P2/PPP1PP1P/RN1QKB1R"

			resultBool := b.Move(values.Coord{X: 6, Y: 6}, values.Coord{X: 5, Y: 5})
			resultFen := b.Debug(board.FenMode)

			checkMove(t, resultBool, expectedBool, resultFen, expectedFen)
		})

		t.Run("when pawn black captures bishop white", func(t *testing.T) {
			expectedFen := "rn1qkb1r/ppp2ppp/3p1p2/8/8/3P1P2/PPP1PP1P/RN1QKB1R"

			resultBool := b.Move(values.Coord{X: 4, Y: 1}, values.Coord{X: 5, Y: 2})
			resultFen := b.Debug(board.FenMode)

			checkMove(t, resultBool, expectedBool, resultFen, expectedFen)
		})

		t.Run("when pawn black captures pawn white en passant", func(t *testing.T) {
			expectedFen := "rn1qkb1r/1pp2ppp/3p1p2/8/7P/1p1P1P2/P1P1PP2/RN1QKB1R"

			b.Move(values.Coord{X: 7, Y: 6}, values.Coord{X: 7, Y: 5})
			b.Move(values.Coord{X: 0, Y: 1}, values.Coord{X: 0, Y: 3})
			b.Move(values.Coord{X: 7, Y: 5}, values.Coord{X: 7, Y: 4})
			b.Move(values.Coord{X: 0, Y: 3}, values.Coord{X: 0, Y: 4})
			b.Move(values.Coord{X: 1, Y: 6}, values.Coord{X: 1, Y: 4})

			resultBool := b.Move(values.Coord{X: 0, Y: 4}, values.Coord{X: 1, Y: 5})
			resultFen := b.Debug(board.FenMode)

			checkMove(t, resultBool, expectedBool, resultFen, expectedFen)
		})

		t.Run("when white has castled kingside and black has castled queenside", func(t *testing.T) {
			expectedFen := "2kr1b1r/1ppq1ppp/2np1p2/8/7P/1P1P1P2/2P1PPB1/RN1Q1RK1"

			b.Move(values.Coord{X: 5, Y: 7}, values.Coord{X: 6, Y: 6})
			b.Move(values.Coord{X: 3, Y: 0}, values.Coord{X: 3, Y: 1})
			b.Move(values.Coord{X: 4, Y: 7}, values.Coord{X: 6, Y: 7}) // White Castle
			b.Move(values.Coord{X: 1, Y: 0}, values.Coord{X: 2, Y: 2})
			b.Move(values.Coord{X: 0, Y: 6}, values.Coord{X: 1, Y: 5})

			resultBool := b.Move(values.Coord{X: 4, Y: 0}, values.Coord{X: 2, Y: 0}) // Black Castle
			resultFen := b.Debug(board.FenMode)

			checkMove(t, resultBool, expectedBool, resultFen, expectedFen)
		})

		t.Run("when white check the king black and black defends", func(t *testing.T) {
			expectedFen := "Rnkr1b1r/1ppq1ppp/3p1p2/8/7P/1P1P1P2/2P1PPB1/1N1Q1RK1"

			b.Move(values.Coord{X: 0, Y: 7}, values.Coord{X: 0, Y: 0}) // Check!
			resultBool := b.Move(values.Coord{X: 2, Y: 2}, values.Coord{X: 1, Y: 0})
			resultFen := b.Debug(board.FenMode)

			checkMove(t, resultBool, expectedBool, resultFen, expectedFen)
		})
	})

	t.Run("when is invalid move", func(t *testing.T) {
		expectedBool := false
		expectedFen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"

		b := board.NewBoard()
		b.Start()

		t.Run("when has pieces between path and piece selected isn't knight", func(t *testing.T) {
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

		t.Run("when origin is empty", func(t *testing.T) {
			resultBool := b.Move(values.Coord{X: 3, Y: 3}, values.Coord{X: 6, Y: 3})
			resultFen := b.Debug(board.FenMode)

			checkMove(t, resultBool, expectedBool, resultFen, expectedFen)
		})

		t.Run("when dest is out of board", func(t *testing.T) {
			t.Run("when move rook on initial board", func(t *testing.T) {
				resultBool := b.Move(values.Coord{X: 7, Y: 7}, values.Coord{X: 8, Y: 7})
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
