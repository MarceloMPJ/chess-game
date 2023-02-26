package board

import (
	"github.com/MarceloMPJ/chess-game/libs/basic"
	"github.com/MarceloMPJ/chess-game/libs/values"
	"github.com/MarceloMPJ/chess-game/pkg/piece"
	"github.com/MarceloMPJ/chess-game/pkg/piece/pawn"
)

func (b *Board) Move(origin, dest values.Coord) bool {
	return b.execMove(origin, dest)
}

func (b *Board) execMove(origin, dest values.Coord) bool {
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

	if b.isCastlingKingSide(origin, dest) && b.allowCastlingKingSide(origin, dest) {
		return b.castleKingSide(origin, dest)
	}

	if b.isCastlingQueenSide(origin, dest) && b.allowCastlingQueenSide(origin, dest) {
		return b.castleQueenSide(origin, dest)
	}

	return b.move(origin, dest, p)
}

func (b *Board) move(origin, dest values.Coord, p piece.PieceContract) bool {
	if !b.allowMove(origin, dest, p) {
		return false
	}

	b.setEnPassant(origin, dest)
	b.setCastling(origin, dest)
	b.moveTo(origin, dest, p)

	return true
}

func (b *Board) capture(origin, dest values.Coord, p piece.PieceContract) bool {
	if b.isPawn(origin) && !b.allowCaptureWithPawn(origin, dest, p.(*pawn.Pawn)) {
		return false
	}

	if !b.isPawn(origin) && !b.allowMove(origin, dest, p) {
		return false
	}

	b.setCastling(origin, dest)
	b.moveTo(origin, dest, p)

	b.resetEnPassant()

	return true
}

func (b *Board) castleKingSide(origin, dest values.Coord) bool {
	b.resetEnPassant()

	if b.currentColor == values.White {
		b.rows[dest.Y][dest.X] = b.rows[initialPositionKingWhiteY][initialPositionKingWhiteX]
		b.rows[dest.Y][dest.X-1] = b.rows[initialPositionRookKingWhiteY][initialPositionRookKingWhiteX]

		b.castlingKingWhite = false
		b.castlingQueenWhite = false
		b.rows[initialPositionKingWhiteY][initialPositionKingWhiteX] = nil
		b.rows[initialPositionRookKingWhiteY][initialPositionRookKingWhiteX] = nil

		b.setCheck(origin, dest)
		b.nextTurn()

		return true
	}

	b.rows[dest.Y][dest.X] = b.rows[initialPositionKingBlackY][initialPositionKingBlackX]
	b.rows[dest.Y][dest.X-1] = b.rows[initialPositionRookKingBlackY][initialPositionRookKingBlackX]

	b.castlingKingBlack = false
	b.castlingQueenBlack = false
	b.rows[initialPositionKingBlackY][initialPositionKingBlackX] = nil
	b.rows[initialPositionRookKingBlackY][initialPositionRookKingBlackX] = nil

	b.setCheck(origin, dest)
	b.nextTurn()

	return true
}

func (b *Board) castleQueenSide(origin, dest values.Coord) bool {
	b.resetEnPassant()

	if b.currentColor == values.White {
		b.rows[dest.Y][dest.X] = b.rows[initialPositionKingWhiteY][initialPositionKingWhiteX]
		b.rows[dest.Y][dest.X+1] = b.rows[initialPositionRookQueenWhiteY][initialPositionRookQueenWhiteX]

		b.castlingKingWhite = false
		b.castlingQueenWhite = false
		b.rows[initialPositionKingWhiteY][initialPositionKingWhiteX] = nil
		b.rows[initialPositionRookQueenWhiteY][initialPositionRookQueenWhiteX] = nil

		b.setCheck(origin, dest)
		b.nextTurn()

		return true
	}

	b.rows[dest.Y][dest.X] = b.rows[initialPositionKingBlackY][initialPositionKingBlackX]
	b.rows[dest.Y][dest.X+1] = b.rows[initialPositionRookQueenBlackY][initialPositionRookQueenBlackX]

	b.castlingKingBlack = false
	b.castlingQueenBlack = false
	b.rows[initialPositionKingBlackY][initialPositionKingBlackX] = nil
	b.rows[initialPositionRookQueenBlackY][initialPositionRookQueenBlackX] = nil

	b.setCheck(origin, dest)
	b.nextTurn()

	return true
}

func (b *Board) moveTo(origin, dest values.Coord, p piece.PieceContract) {
	b.rows[dest.Y][dest.X] = p

	if b.isValidEnPassant(origin, dest) {
		b.rows[origin.Y][dest.X] = nil
	}
	b.rows[origin.Y][origin.X] = nil
	b.setCheck(origin, dest)
	b.nextTurn()
}

func (b *Board) setCheck(origin, dest values.Coord) {
	targetColor := values.White
	if b.currentColor == values.White {
		targetColor = values.Black
	}

	if b.isKingAttacked(targetColor) {
		if targetColor == values.White {
			b.checkWhite = true

			return
		}

		b.checkBlack = true
	}
}

func (b *Board) setCastling(origin, dest values.Coord) {
	if b.currentColor == values.White {
		// Verify when moves a king white
		if b.isKing(origin) {
			b.castlingKingWhite = false
			b.castlingQueenWhite = false
			return
		}

		// Verify when moves a rook white of king
		if b.isRook(origin) && origin.X == initialPositionRookKingWhiteX && origin.Y == initialPositionRookKingWhiteY {
			b.castlingKingWhite = false
			return
		}

		// Verify when captures rook black of king
		if b.isRook(dest) && dest.X == initialPositionRookKingBlackX && dest.Y == initialPositionRookKingBlackY {
			b.castlingKingBlack = false
			return
		}

		// Verify when moves a rook white of queen
		if b.isRook(origin) && origin.X == initialPositionRookQueenWhiteX && origin.Y == initialPositionRookQueenWhiteY {
			b.castlingQueenWhite = false
			return
		}

		// Verify when captures rook black of queen
		if b.isRook(dest) && dest.X == initialPositionRookQueenBlackX && dest.Y == initialPositionRookQueenBlackY {
			b.castlingQueenBlack = false
			return
		}
	} else {
		// Verify when moves a king black
		if b.isKing(origin) {
			b.castlingKingBlack = false
			b.castlingQueenBlack = false
			return
		}

		// Verify when moves a rook black of king
		if origin.X == initialPositionRookKingBlackX && origin.Y == initialPositionRookKingBlackY && b.isRook(origin) {
			b.castlingKingBlack = false
			return
		}

		// Verify when captures rook white of king
		if b.isRook(dest) && dest.X == initialPositionRookKingWhiteX && dest.Y == initialPositionRookKingWhiteY {
			b.castlingKingBlack = false
			return
		}

		// Verify when moves a rook black of queen
		if origin.X == initialPositionRookQueenBlackX && origin.Y == initialPositionRookQueenBlackY && b.isRook(origin) {
			b.castlingQueenBlack = false
			return
		}

		// Verify when captures rook white of queen
		if b.isRook(dest) && dest.X == initialPositionRookQueenWhiteX && dest.Y == initialPositionRookQueenWhiteY {
			b.castlingQueenBlack = false
			return
		}
	}
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
