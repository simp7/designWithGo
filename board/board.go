package board

import (
	"github.com/simp7/designWithGo/board/cell"
	"github.com/simp7/designWithGo/core"
)

type Board struct {
	xSize int
	ySize int
	board [][]cell.Cell
}

var instance *Board = nil

func Get() *Board {
	if instance == nil {
		instance = new(Board)
	}
	return instance
}

func (b *Board) Add(object core.Object) {
}

func (b *Board) GetChild(idx int) core.Object {
	if idx > b.xSize*b.ySize {
		return nil
	}
	return b.board[idx/b.xSize][idx%b.xSize]
}

func (b *Board) Create(builder Builder) {
	builder.buildSize()
	builder.buildCell()
	builder.buildPiece()
}

func (b *Board) Clone(pos *core.Pos) core.Object {
	return instance
}

func (b *Board) Highlight() {
}

func (b *Board) OnClick() {

}

func (b *Board) Draw() {
	for i := range b.board {
		for _, v := range b.board[i] {
			v.Draw()
		}
	}
}

func (b *Board) Remove(object core.Object) {
}
