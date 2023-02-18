package pawn

import "github.com/MarceloMPJ/chess-game/libs/values"

type Pawn struct {
	color int
}

const initialYBoardWhite = 6
const initialYBoardBlack = 1

func NewPawn(color int) Pawn {
	return Pawn{color}
}

func (p *Pawn) Show() rune {
	if p.color == values.White {
		return '♟'
	}

	return '♙'
}

func (p *Pawn) ShowFEN() rune {
	if p.color == values.White {
		return 'P'
	}

	return 'p'
}

func (p *Pawn) IsValidMove(origin, dest values.Coord) bool {
	if origin.X != dest.X {
		return false
	}

	return p.checkValidSteps(origin, p.steps(origin, dest))
}

func (p *Pawn) steps(origin, dest values.Coord) uint8 {
	var steps uint8

	if p.color == values.White && origin.Y > dest.Y {
		steps = origin.Y - dest.Y
	}

	if p.color == values.Black && dest.Y > origin.Y {
		steps = dest.Y - origin.Y
	}

	return steps
}

func (p *Pawn) checkValidSteps(origin values.Coord, steps uint8) bool {
	if steps == 1 {
		return true
	}

	if p.color == values.White && origin.Y == initialYBoardWhite && steps == 2 {
		return true
	}

	if p.color == values.Black && origin.Y == initialYBoardBlack && steps == 2 {
		return true
	}

	return false
}
