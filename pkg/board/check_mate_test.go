package board_test

import (
	"testing"

	"github.com/MarceloMPJ/chess-game/pkg/board"
)

func TestBoard_IsCheckMate(t *testing.T) {
	tests := []struct {
		fen      string
		expected bool
	}{
		{"rnbqk1nr/pppp1ppp/4p3/8/6Pb/2N2P2/PPPPP2P/R1BQKBNR", true},
		{"rnb1k1nr/pppppp1p/8/q7/8/3P4/PP2PPnP/RN1QKNNR", true},
		{"rnbqk1nr/pppppp1p/8/8/8/8/PP1PPPnP/RN1QKNNR", true},
	}

	for _, test := range tests {
		b := board.Build(test.fen)
		result := b.IsCheckMate()

		if result != test.expected {
			t.Errorf("result: %t, expected: %t", result, test.expected)
		}
	}
}
