package queen

import (
	"github.com/MarceloMPJ/chess-game/libs/basic"
	"github.com/MarceloMPJ/chess-game/libs/values"
)

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

func (q *Queen) ShowFEN() rune {
	if q.color == values.White {
		return 'Q'
	}

	return 'q'
}

func (b *Queen) IsValidMove(origin, dest values.Coord) bool {
	return basic.Abs(int(origin.X)-int(dest.X)) == basic.Abs(int(origin.Y)-int(dest.Y)) ||
		(origin.X == dest.X || origin.Y == dest.Y)
}
