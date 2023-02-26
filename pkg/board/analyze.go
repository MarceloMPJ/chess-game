package board

import (
	"github.com/MarceloMPJ/chess-game/libs/basic"
	"github.com/MarceloMPJ/chess-game/libs/values"
	"github.com/MarceloMPJ/chess-game/pkg/piece"
)

func (b *Board) isValidEnPassant(origin, dest values.Coord) bool {
	return b.isPawn(origin) && b.enPassant != nil && (dest.Y == b.enPassant.Y && dest.X == b.enPassant.X)
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
