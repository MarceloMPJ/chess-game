package game

import "github.com/MarceloMPJ/chess-game/pkg/table"

func Debug() string {
	table := table.NewTable()

	return table.Debug()
}
