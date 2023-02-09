package bishop

import (
	"github.com/MarceloMPJ/chess-game/libs/basic"
	"github.com/MarceloMPJ/chess-game/libs/values"
)

type Bishop struct {
	color int
}

func NewBishop(color int) Bishop {
	return Bishop{color}
}

func (b *Bishop) Show() rune {
	if b.color == values.White {
		return '♝'
	}

	return '♗'
}

func (b *Bishop) ShowFEN() rune {
	if b.color == values.White {
		return 'B'
	}

	return 'b'
}

func (b *Bishop) IsValidMove(origin, dest values.Coord) bool {
	return basic.Abs(int(origin.X)-int(dest.X)) == basic.Abs(int(origin.Y)-int(dest.Y))
}
