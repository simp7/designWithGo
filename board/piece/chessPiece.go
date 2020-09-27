package piece

import (
	"fmt"
	"github.com/simp7/designWithGo/core"
)

type team string

const (
	BLACK team = "black"
	WHITE      = "white"
)

type ChessPiece interface {
	Piece
	Move()
}

type chessPiece struct {
	pos       *core.Pos
	pieceType string
	team      team
}

func newChessPiece(pos *core.Pos, team team, piece string) *chessPiece {
	p := new(chessPiece)
	p.pos = pos
	p.team = team
	p.pieceType = piece
	return p
}

func (p *chessPiece) Add(obj core.Object) {
}

func (p *chessPiece) Clone(pos *core.Pos) core.Object {
	return newChessPiece(pos, p.team, p.pieceType)
}

func (p *chessPiece) GetChild(idx int) core.Object {
	return nil
}

func (p *chessPiece) Highlight() {

}

func (p *chessPiece) OnClick() {

}

func (p *chessPiece) Draw() {
	fmt.Printf("I'm %s %s in ( %d , %d )\n", p.team, p.pieceType, p.pos.GetX(), p.pos.GetY())
}

func (p *chessPiece) Move() {

}

func (p *chessPiece) Remove(obj core.Object) {

}
