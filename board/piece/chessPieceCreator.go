package piece

import "github.com/simp7/designWithGo/core"

const (
	PAWN   = "pawn"
	KNIGHT = "knight"
	ROOK   = "rook"
	BISHOP = "bishop"
	KING   = "king"
	QUEEN  = "queen"
)

type ChessPieceCreator struct {
	team team
}

func NewChessPieceCreator(t team) *ChessPieceCreator {
	c := new(ChessPieceCreator)
	c.team = t
	return c
}

func (c *ChessPieceCreator) Create(id string, pos *core.Pos) *chessPiece {
	return newChessPiece(pos, c.team, id)
}
