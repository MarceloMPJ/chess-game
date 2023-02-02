package queen

import "github.com/MarceloMPJ/chess-game/libs/values"

type Queen struct {
	color int
}

func NewQueen(color int) Queen {
	return Queen{color}
}

func (q *Queen) Show() rune {
	if q.color == values.White {
		return '♛'
	}

	return '♕'
}
