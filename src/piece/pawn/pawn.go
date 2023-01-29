package pawn

type Pawn struct {
	color int
}

func NewPawn(color int) Pawn {
	return Pawn{color}
}

func (p *Pawn) Show() rune {
	if p.color == 0 {
		return '♟'
	}

	return '♙'
}
