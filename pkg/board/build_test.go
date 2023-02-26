package board_test

import (
	"testing"

	"github.com/MarceloMPJ/chess-game/pkg/board"
)

func TestBoard_Build(t *testing.T) {
	gamesFEN := []string{
		// Initial position
		"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR",

		// Bobby Fischer X Boris Spassky (1972)
		"rnbqkbnr/pppp1ppp/8/4p3/4P3/5N2/PPPP1PPP/RNBQKB1R",

		// Garry Kasparov X Deep Blue (1997)
		"r4rk1/1pp1qppp/p1n1pn2/8/2P5/1P1P1NP1/PBQ1PPBP/RN3RK1",

		// Adolf Anderssen X Lionel Kieseritzky (1851)
		"rnbqkbnr/pppppppp/8/8/8/1P6/P1PPPPPP/RNBQKBNR",

		// Adolf Anderssen X Jean Dufresne (1852)
		"r1bqkbnr/pppp1ppp/2n5/2b1p3/2B1P3/5N2/PPPP1PPP/RNBQK2R",

		// Anatoly Karpov X Garry Kasparov (1985)
		"r4rk1/pppb2pp/2n1p3/4Pp2/1PP1pP2/2N5/PB3KPP/R6R",
	}

	for _, fen := range gamesFEN {
		b := board.Build(fen)
		result := b.Debug(board.FenMode)

		if result != fen {
			t.Errorf("result: %s, expected: %s", result, fen)
		}
	}
}
