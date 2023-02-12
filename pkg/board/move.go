package board

import (
	"github.com/MarceloMPJ/chess-game/libs/basic"
	"github.com/MarceloMPJ/chess-game/libs/values"
	"github.com/MarceloMPJ/chess-game/pkg/piece"
)

func (b *Board) Move(origin, dest values.Coord) bool {
	p := b.rows[origin.Y][origin.X]

	if b.allowMove(origin, dest, p) {
		b.rows[origin.Y][origin.X] = nil
		b.rows[dest.Y][dest.X] = p
		b.nextTurn()

		return true
	}

	return false
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
				if origin.Y == i && origin.X == j {
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
			if origin.Y == i && origin.X == j {
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

func (b *Board) isKnight(origin values.Coord) bool {
	fen := b.rows[origin.Y][origin.X].ShowFEN()

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
	fen := b.rows[origin.Y][origin.X].ShowFEN()

	if fen < 'Z' {
		return values.White
	}

	return values.Black
}
