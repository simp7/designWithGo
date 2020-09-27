package board

import (
	"github.com/simp7/designWithGo/board/cell"
	"github.com/simp7/designWithGo/board/piece"
)

type ChessBuilder struct {
}

func NewChessBuilder() *ChessBuilder {
	return new(ChessBuilder)
}

func (b *ChessBuilder) buildSize() {
	instance.xSize = 8
	instance.ySize = 8
	instance.board = make([][]cell.Cell, instance.ySize)
	for i := range instance.board {
		instance.board[i] = make([]cell.Cell, instance.xSize)
	}
}

func (b *ChessBuilder) buildCell() {
	prototype := cell.NewChessCell(nil)
	buildCell(prototype)
}

func (b *ChessBuilder) buildPiece() {
	buildBlackPiece()
	buildWhitePiece()
}

func buildBlackPiece() {
	creator := piece.NewChessPieceCreator(piece.BLACK)
	pawn := creator.Create(piece.PAWN, nil)
	for i := range instance.board[6] {
		instance.board[6][i] = pawn
	}
}

func buildWhitePiece() {
	creator := piece.NewChessPieceCreator(piece.WHITE)
	pawn := creator.Create(piece.PAWN, nil)
	for i := range instance.board[1] {
		instance.board[1][i] = pawn
	}
}
