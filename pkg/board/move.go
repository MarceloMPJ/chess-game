package board

import (
	"github.com/MarceloMPJ/chess-game/libs/basic"
	"github.com/MarceloMPJ/chess-game/libs/values"
	"github.com/MarceloMPJ/chess-game/pkg/piece"
	"github.com/MarceloMPJ/chess-game/pkg/piece/pawn"
)

const sizeOfBoard = 8

func (b *Board) Move(origin, dest values.Coord) bool {
	if isInsideBoard(origin) || isInsideBoard(dest) {
		return false
	}

	p := b.rows[origin.Y][origin.X]
	pTarget := b.rows[dest.Y][dest.X]

	if p == nil {
		return false
	}

	if b.isValidEnPassant(origin, dest) {
		return b.capture(origin, dest, p)
	}

	if pTarget != nil {
		if b.pieceColor(origin) != b.pieceColor(dest) {
			return b.capture(origin, dest, p)
		}

		return false
	}

	return b.move(origin, dest, p)
}

func (b *Board) move(origin, dest values.Coord, p piece.PieceContract) bool {
	if !b.allowMove(origin, dest, p) {
		return false
	}

	b.setEnPassant(origin, dest)
	b.moveTo(origin, dest, p)

	return true
}

func (b *Board) capture(origin, dest values.Coord, p piece.PieceContract) bool {
	if b.isPawn(origin) && (!p.(*pawn.Pawn).IsValidCapture(origin, dest) || !b.isCorrectTurn(origin)) {
		return false
	}

	if !b.isPawn(origin) && !b.allowMove(origin, dest, p) {
		return false
	}

	b.moveTo(origin, dest, p)
	b.resetEnPassant()

	return true
}

func (b *Board) moveTo(origin, dest values.Coord, p piece.PieceContract) {
	b.rows[dest.Y][dest.X] = p

	if b.isValidEnPassant(origin, dest) {
		b.rows[origin.Y][dest.X] = nil
	}
	b.rows[origin.Y][origin.X] = nil
	b.nextTurn()
}

func (b *Board) isValidEnPassant(origin, dest values.Coord) bool {
	return b.isPawn(origin) && b.enPassant != nil && (dest.Y == b.enPassant.Y && dest.X == b.enPassant.X)
}

func (b *Board) setEnPassant(origin, dest values.Coord) {
	ymin, ymax := basic.MinUint8(origin.Y, dest.Y), basic.MaxUint8(origin.Y, dest.Y)

	var coord values.Coord

	if b.isPawn(origin) && ymax-ymin == 2 {
		coord.X = origin.X

		if origin.Y > dest.Y {
			coord.Y = origin.Y - 1
		} else {
			coord.Y = origin.Y + 1
		}

		b.enPassant = &coord

		return
	}

	b.enPassant = nil
}

func (b *Board) resetEnPassant() {
	b.enPassant = nil
}

func (b *Board) nextTurn() {
	if b.currentColor == values.White {
		b.currentColor = values.Black

		return
	}

	b.currentColor = values.White
}

func (b *Board) allowMove(origin, dest values.Coord, p piece.PieceContract) bool {
	return b.isCorrectTurn(origin) && p.IsValidMove(origin, dest) && b.isFreePath(origin, dest, p)
}

func (b *Board) isFreePath(origin, dest values.Coord, p piece.PieceContract) bool {
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

func (b *Board) isCorrectTurn(origin values.Coord) bool {
	return b.currentColor == b.pieceColor(origin)
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

func (b *Board) hasPiece(y, x uint8) bool {
	return b.rows[y][x] != nil
}

func nextStep(i, j uint8, dest values.Coord) (uint8, uint8) {
	var nextY, nextX uint8

	if i < dest.Y {
		nextY = i + 1
	} else {
		nextY = i - 1
	}

	if j < dest.X {
		nextX = j + 1
	} else {
		nextX = j - 1
	}

	return nextY, nextX
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

func isInsideBoard(coord values.Coord) bool {
	return coord.X >= sizeOfBoard || coord.Y >= sizeOfBoard
}
