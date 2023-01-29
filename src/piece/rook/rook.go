package rook

type Rook struct {
	color int
}

func NewRook(color int) Rook {
	return Rook{color}
}

func (r *Rook) Show() rune {
	if r.color == 0 {
		return '♜'
	}

	return '♖'
}
