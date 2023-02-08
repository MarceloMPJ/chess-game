package game

import "github.com/MarceloMPJ/chess-game/pkg/board"

func Debug() string {
	board := board.NewBoard()

	return board.Debug()
}
