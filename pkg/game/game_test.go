package game_test

import (
	"testing"

	"github.com/MarceloMPJ/chess-game/pkg/game"
)

func TestGame_Debug(t *testing.T) {
	expected := `♖ ♘ ♗ ♔ ♕ ♗ ♘ ♖
♙ ♙ ♙ ♙ ♙ ♙ ♙ ♙
               
               
               
               
♟ ♟ ♟ ♟ ♟ ♟ ♟ ♟
♜ ♞ ♝ ♚ ♛ ♝ ♞ ♜
`
	result := game.Debug()

	if result != expected {
		t.Errorf("result: %s, expected: %s", result, expected)
	}
}
