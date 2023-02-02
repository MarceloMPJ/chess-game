package pawn

import "github.com/MarceloMPJ/chess-game/libs/values"

type Pawn struct {
	color int
}

func NewPawn(color int) Pawn {
	return Pawn{color}
}

func (p *Pawn) Show() rune {
	if p.color == values.White {
		return '♟'
	}

	return '♙'
}

func (p *Pawn) IsValidMove(origin, dest values.Coord) bool {
	if origin.X != dest.X {
		return false
	}

	if p.color == values.White && origin.Y > dest.Y && (origin.Y-dest.Y == 1 || (origin.Y-dest.Y == 2)) {
		return true
	}

	if p.color == values.Black && dest.Y > origin.Y && (dest.Y-origin.Y == 1 || (dest.Y-origin.Y == 2)) {
		return true
	}

	return false
}
