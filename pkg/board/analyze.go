package board

import (
	"github.com/MarceloMPJ/chess-game/libs/basic"
	"github.com/MarceloMPJ/chess-game/libs/values"
	"github.com/MarceloMPJ/chess-game/pkg/piece"
	"github.com/MarceloMPJ/chess-game/pkg/piece/pawn"
)

func (b *Board) allowMove(origin, dest values.Coord, p piece.PieceContract) bool {
	return b.isCorrectTurn(origin) && p.IsValidMove(origin, dest) &&
		b.isFreePath(origin, dest) && (!b.isKing(origin) || !b.isAttacked(dest)) && b.simulateMove(origin, dest)
}

func (b *Board) allowCaptureWithPawn(origin, dest values.Coord, p *pawn.Pawn) bool {
	return p.IsValidCapture(origin, dest) && b.isCorrectTurn(origin) && b.simulateMove(origin, dest)
}

func (b *Board) simulateMove(origin, dest values.Coord) bool {
	if b.isSimulatation {
		return true
	}

	bclone := *b

	bclone.isSimulatation = true
	bclone.execMove(origin, dest)
	bclone.nextTurn()

	return !bclone.isKingAttacked(b.currentColor)
}

func (b *Board) isValidEnPassant(origin, dest values.Coord) bool {
	return b.isPawn(origin) && b.enPassant != nil && (dest.Y == b.enPassant.Y && dest.X == b.enPassant.X)
}

func (b *Board) isCastlingKingSide(origin, dest values.Coord) bool {
	if !b.isKing(origin) {
		return false
	}

	if b.currentColor == values.White {
		return origin.X == initialPositionKingWhiteX && origin.Y == initialPositionKingWhiteY && dest.Y == initialPositionKingWhiteY &&
			dest.X == initialPositionKingWhiteX+2
	}

	return origin.X == initialPositionKingBlackX && origin.Y == initialPositionKingBlackY && dest.Y == initialPositionKingBlackY &&
		dest.X == initialPositionKingBlackX+2
}

func (b *Board) isCastlingQueenSide(origin, dest values.Coord) bool {
	if !b.isKing(origin) {
		return false
	}

	if b.currentColor == values.White {
		return origin.X == initialPositionKingWhiteX && origin.Y == initialPositionKingWhiteY && dest.Y == initialPositionKingWhiteY &&
			dest.X == initialPositionKingWhiteX-2
	}

	return origin.X == initialPositionKingBlackX && origin.Y == initialPositionKingBlackY && dest.Y == initialPositionKingBlackY &&
		dest.X == initialPositionKingBlackX-2
}

func (b *Board) allowCastlingKingSide(origin, dest values.Coord) bool {
	if !b.isCorrectTurn(origin) || !b.isFreePath(origin, dest) {
		return false
	}

	if b.currentColor == values.White {
		return b.castlingKingWhite && !b.isAttacked(values.Coord{Y: 7, X: 5}) && !b.isAttacked(values.Coord{Y: 7, X: 6})
	}

	return b.castlingKingBlack && !b.isAttacked(values.Coord{Y: 0, X: 5}) && !b.isAttacked(values.Coord{Y: 0, X: 6})
}

func (b *Board) allowCastlingQueenSide(origin, dest values.Coord) bool {
	if !b.isCorrectTurn(origin) || !b.isFreePath(origin, dest) {
		return false
	}

	if b.currentColor == values.White {
		return b.castlingQueenWhite && !b.isAttacked(values.Coord{Y: 7, X: 2}) && !b.isAttacked(values.Coord{Y: 7, X: 3})
	}

	return b.castlingQueenBlack && !b.isAttacked(values.Coord{Y: 0, X: 2}) && !b.isAttacked(values.Coord{Y: 0, X: 3})
}

func (b *Board) isFreePath(origin, dest values.Coord) bool {
	if b.isKnight(origin) {
		return true
	}

	// When the path is horizontal or vertical
	if origin.X == dest.X || origin.Y == dest.Y {
		startX, startY := basic.MinUint8(origin.X, dest.X), basic.MinUint8(origin.Y, dest.Y)
		finishX, finishY := basic.MaxUint8(origin.X, dest.X), basic.MaxUint8(origin.Y, dest.Y)

		for i := startY; i <= finishY; i++ {
			for j := startX; j <= finishX; j++ {
				if (origin.Y == i && origin.X == j) || (dest.Y == i && dest.X == j) {
					continue
				}

				if b.hasPiece(i, j) {
					return false
				}
			}
		}
	} else {
		// When the path is horizontal
		for i, j := origin.Y, origin.X; i != dest.Y && j != dest.X; i, j = nextStep(i, j, dest) {
			if (origin.Y == i && origin.X == j) || (dest.Y == i && dest.X == j) {
				continue
			}

			if b.hasPiece(i, j) {
				return false
			}
		}
	}

	return true
}

func (b *Board) isKingAttacked(targetColor int) bool {
	coord := b.searchKing(targetColor)
	return b.isAttacked(coord)
}

func (b *Board) searchKing(targetColor int) (currentCoord values.Coord) {
	for i := uint8(0); i < sizeOfBoard; i++ {
		for j := uint8(0); j < sizeOfBoard; j++ {
			currentCoord = values.Coord{X: j, Y: i}

			if b.pieceColor(currentCoord) == targetColor && b.isKing(currentCoord) {
				return
			}
		}
	}

	return
}

func (b *Board) isAttacked(origin values.Coord) bool {
	targetColor := values.White
	if b.currentColor == values.White {
		targetColor = values.Black
	}

	return b.isAttackedByRookOrQueen(origin, targetColor) || b.isAttackedByBishopOrQueen(origin, targetColor) ||
		b.isAttackedByKnight(origin, targetColor) || b.isAttackedByPawn(origin, targetColor) ||
		b.isAttackedByKing(origin, targetColor)
}

func (b *Board) isAttackedByRookOrQueen(origin values.Coord, targetColor int) bool {
	for i := int(origin.X) - 1; i >= 0; i-- {
		currentCoord := values.Coord{X: uint8(i), Y: origin.Y}

		if b.pieceColor(currentCoord) == targetColor && (b.isRook(currentCoord) || b.isQueen(currentCoord)) {
			return true
		}

		if b.rows[origin.Y][i] != nil {
			break
		}
	}

	for i := int(origin.X) + 1; i < sizeOfBoard; i++ {
		currentCoord := values.Coord{X: uint8(i), Y: origin.Y}

		if b.pieceColor(currentCoord) == targetColor && (b.isRook(currentCoord) || b.isQueen(currentCoord)) {
			return true
		}

		if b.rows[origin.Y][i] != nil {
			break
		}
	}

	for i := int(origin.Y) - 1; i >= 0; i-- {
		currentCoord := values.Coord{X: origin.X, Y: uint8(i)}

		if b.pieceColor(currentCoord) == targetColor && (b.isRook(currentCoord) || b.isQueen(currentCoord)) {
			return true
		}

		if b.rows[i][origin.X] != nil {
			break
		}
	}

	for i := int(origin.Y) + 1; i < sizeOfBoard; i++ {
		currentCoord := values.Coord{X: origin.X, Y: uint8(i)}

		if b.pieceColor(currentCoord) == targetColor && (b.isRook(currentCoord) || b.isQueen(currentCoord)) {
			return true
		}

		if b.rows[i][origin.X] != nil {
			break
		}
	}

	return false
}

func (b *Board) isAttackedByBishopOrQueen(origin values.Coord, targetColor int) bool {
	for i, j := int(origin.Y)-1, int(origin.X)-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		currentCoord := values.Coord{X: uint8(j), Y: uint8(i)}

		if b.pieceColor(currentCoord) == targetColor && (b.isBishop(currentCoord) || b.isQueen(currentCoord)) {
			return true
		}

		if b.rows[i][j] != nil {
			break
		}
	}

	for i, j := int(origin.Y)-1, int(origin.X)+1; i >= 0 && j < sizeOfBoard; i, j = i-1, j+1 {
		currentCoord := values.Coord{X: uint8(j), Y: uint8(i)}

		if b.pieceColor(currentCoord) == targetColor && (b.isBishop(currentCoord) || b.isQueen(currentCoord)) {
			return true
		}

		if b.rows[i][j] != nil {
			break
		}
	}

	for i, j := int(origin.Y)+1, int(origin.X)-1; i < sizeOfBoard && j >= 0; i, j = i+1, j-1 {
		currentCoord := values.Coord{X: uint8(j), Y: uint8(i)}

		if b.pieceColor(currentCoord) == targetColor && (b.isBishop(currentCoord) || b.isQueen(currentCoord)) {
			return true
		}

		if b.rows[i][j] != nil {
			break
		}
	}

	for i, j := int(origin.Y)+1, int(origin.X)+1; i < sizeOfBoard && j < sizeOfBoard; i, j = i+1, j+1 {
		currentCoord := values.Coord{X: uint8(j), Y: uint8(i)}

		if b.pieceColor(currentCoord) == targetColor && (b.isBishop(currentCoord) || b.isQueen(currentCoord)) {
			return true
		}

		if b.rows[i][j] != nil {
			break
		}
	}

	return false
}

func (b *Board) isAttackedByKnight(origin values.Coord, targetColor int) bool {
	knightMoves := [8]values.Move{
		{X: -2, Y: -1},
		{X: -1, Y: -2},
		{X: -2, Y: 1},
		{X: 1, Y: -2},
		{X: -1, Y: 2},
		{X: 2, Y: -1},
		{X: 1, Y: 2},
		{X: 2, Y: 1},
	}

	for _, move := range knightMoves {
		currentX, currentY := int(origin.X)+move.X, int(origin.Y)+move.Y
		currentCoord := values.Coord{X: uint8(currentX), Y: uint8(currentY)}

		if currentX >= 0 && currentX < sizeOfBoard && currentY >= 0 && currentY < sizeOfBoard {
			if b.pieceColor(currentCoord) == targetColor && b.isKnight(currentCoord) {
				return true
			}
		}
	}

	return false
}

func (b *Board) isAttackedByPawn(origin values.Coord, targetColor int) bool {
	if targetColor == values.White && int(origin.Y)-1 >= 0 {
		if int(origin.X)-1 >= 0 {
			coord := values.Coord{X: origin.X - 1, Y: origin.Y - 1}

			if b.pieceColor(coord) == targetColor && b.isPawn(coord) {
				return true
			}
		}

		if int(origin.X)+1 < sizeOfBoard {
			coord := values.Coord{X: origin.X + 1, Y: origin.Y - 1}

			if b.pieceColor(coord) == targetColor && b.isPawn(coord) {
				return true
			}
		}
	}

	if targetColor == values.Black && int(origin.Y)+1 < sizeOfBoard {
		if int(origin.X)-1 >= 0 {
			coord := values.Coord{X: origin.X - 1, Y: origin.Y + 1}

			if b.pieceColor(coord) == targetColor && b.isPawn(coord) {
				return true
			}
		}

		if int(origin.X)+1 < sizeOfBoard {
			coord := values.Coord{X: origin.X + 1, Y: origin.Y + 1}

			if b.pieceColor(coord) == targetColor && b.isPawn(coord) {
				return true
			}
		}
	}

	return false
}

func (b *Board) isAttackedByKing(origin values.Coord, targetColor int) bool {
	kingMoves := [8]values.Move{
		{X: -1, Y: -1},
		{X: -1, Y: 0},
		{X: -1, Y: 1},
		{X: 0, Y: 1},
		{X: 1, Y: 1},
		{X: 1, Y: 0},
		{X: 1, Y: -1},
		{X: 0, Y: -1},
	}

	for _, move := range kingMoves {
		currentX, currentY := int(origin.X)+move.X, int(origin.Y)+move.Y
		currentCoord := values.Coord{X: uint8(currentX), Y: uint8(currentY)}

		if currentX >= 0 && currentX < sizeOfBoard && currentY >= 0 && currentY < sizeOfBoard {
			if b.pieceColor(currentCoord) == targetColor && b.isKing(currentCoord) {
				return true
			}
		}
	}

	return false
}

func (b *Board) isCorrectTurn(origin values.Coord) bool {
	return b.currentColor == b.pieceColor(origin)
}

func (b *Board) hasPiece(y, x uint8) bool {
	return b.rows[y][x] != nil
}

func isInsideBoard(coord values.Coord) bool {
	return coord.X >= sizeOfBoard || coord.Y >= sizeOfBoard
}

func (b *Board) pieceColor(origin values.Coord) int {
	p := b.rows[origin.Y][origin.X]

	if p == nil {
		return -1
	}

	fen := p.ShowFEN()

	if fen > 'A' && fen < 'Z' {
		return values.White
	}

	return values.Black
}

func (b *Board) isPawn(origin values.Coord) bool {
	p := b.rows[origin.Y][origin.X]

	if p == nil {
		return false
	}

	fen := p.ShowFEN()

	return fen == 'p' || fen == 'P'
}

func (b *Board) isKnight(origin values.Coord) bool {
	p := b.rows[origin.Y][origin.X]

	if p == nil {
		return false
	}

	fen := p.ShowFEN()

	return fen == 'n' || fen == 'N'
}

func (b *Board) isKing(origin values.Coord) bool {
	p := b.rows[origin.Y][origin.X]

	if p == nil {
		return false
	}

	fen := p.ShowFEN()

	return fen == 'k' || fen == 'K'
}

func (b *Board) isBishop(origin values.Coord) bool {
	p := b.rows[origin.Y][origin.X]

	if p == nil {
		return false
	}

	fen := p.ShowFEN()

	return fen == 'b' || fen == 'B'
}

func (b *Board) isRook(origin values.Coord) bool {
	p := b.rows[origin.Y][origin.X]

	if p == nil {
		return false
	}

	fen := p.ShowFEN()

	return fen == 'r' || fen == 'R'
}

func (b *Board) isQueen(origin values.Coord) bool {
	p := b.rows[origin.Y][origin.X]

	if p == nil {
		return false
	}

	fen := p.ShowFEN()

	return fen == 'q' || fen == 'Q'
}
