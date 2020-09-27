package cell

import (
	"fmt"
	"github.com/simp7/designWithGo/core"
)

type ChessCell struct {
	pos   *core.Pos
	child core.Object
}

func NewChessCell(pos *core.Pos) *ChessCell {
	c := new(ChessCell)
	c.pos = pos
	return c
}

func (c *ChessCell) Add(piece core.Object) {
	c.child = piece
}

func (c *ChessCell) Clone(pos *core.Pos) core.Object {
	return NewChessCell(pos)
}

func (c *ChessCell) Highlight() {
}

func (c *ChessCell) OnClick() {
}

func (c *ChessCell) Draw() {
	if c.child != nil {
		c.child.Draw()
	}
	sum := c.pos.GetX() + c.pos.GetY()
	fmt.Printf("(%d, %d) = ", c.pos.GetX(), c.pos.GetY())
	if sum%2 == 0 {
		fmt.Println("Black")
	} else {
		fmt.Println("White")
	}
}

func (c *ChessCell) GetChild(idx int) core.Object {
	return c.child
}
func (c *ChessCell) Remove(obj core.Object) {
	c.child = nil
}
