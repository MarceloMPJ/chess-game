package board

import "github.com/MarceloMPJ/chess-game/libs/values"

func (b *Board) IsCheckMate() bool {
	if !b.isKingAttacked(b.currentColor) {
		return false
	}

	return !b.canMoveKing() && !b.canMoveAnyPiece()
}

func (b *Board) canMoveKing() bool {
	kingCoord := b.searchKing(b.currentColor)

	for _, move := range kingMoves {
		currentX, currentY := int(kingCoord.X)+move.X, int(kingCoord.Y)+move.Y
		dest := values.Coord{X: uint8(currentX), Y: uint8(currentY)}

		if currentX >= 0 && currentX < sizeOfBoard && currentY >= 0 && currentY < sizeOfBoard {
			if b.simulateMove(kingCoord, dest) {
				return true
			}
		}
	}

	return false
}

func (b *Board) canMoveAnyPiece() bool {
	for originY := 0; originY < sizeOfBoard; originY++ {
		for originX := 0; originX < sizeOfBoard; originX++ {
			origin := values.Coord{X: uint8(originX), Y: uint8(originY)}

			if b.pieceColor(origin) != b.currentColor {
				continue
			}

			for destY := 0; destY < sizeOfBoard; destY++ {
				for destX := 0; destX < sizeOfBoard; destX++ {
					dest := values.Coord{X: uint8(destX), Y: uint8(destY)}

					if b.simulateMove(origin, dest) {
						return true
					}
				}
			}
		}
	}

	return false
}
