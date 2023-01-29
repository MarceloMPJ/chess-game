package queen

type Queen struct {
	color int
}

func NewQueen(color int) Queen {
	return Queen{color}
}

func (q *Queen) Show() rune {
	if q.color == 0 {
		return '♛'
	}

	return '♕'
}
