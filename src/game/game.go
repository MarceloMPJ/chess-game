package game

import "github.com/MarceloMPJ/chess-game/src/table"

func Debug() string {
	table := table.NewTable()

	return table.Debug()
}
