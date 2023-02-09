package game

import "github.com/MarceloMPJ/chess-game/pkg/board"

func Debug() string {
	b := board.NewBoard()
	b.Start()

	return b.Debug(board.GraphicalMode)
}
