package knight

type Knight struct {
	color int
}

func NewKnight(color int) Knight {
	return Knight{color}
}

func (n *Knight) Show() rune {
	if n.color == 0 {
		return '♞'
	}

	return '♘'
}
