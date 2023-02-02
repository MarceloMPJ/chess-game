package knight

import "github.com/MarceloMPJ/chess-game/libs/values"

type Knight struct {
	color int
}

func NewKnight(color int) Knight {
	return Knight{color}
}

func (n *Knight) Show() rune {
	if n.color == values.White {
		return '♞'
	}

	return '♘'
}
