package bishop

import (
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

func (b *Bishop) IsValidMove(origin, dest values.Coord) bool {
	return abs(int(origin.X)-int(dest.X)) == abs(int(origin.Y)-int(dest.Y))
}

func abs(num int) int {
	if num < 0 {
		return -num
	}

	return num
}
