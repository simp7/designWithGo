package board

import (
	"github.com/simp7/designWithGo/board/cell"
	"github.com/simp7/designWithGo/core"
)

type Builder interface {
	buildSize()
	buildCell()
	buildPiece()
}

func buildCell(cell cell.Cell) {
	for i, v := range instance.board {
		for j := range v {
			pos := core.NewPos(j, i)
			instance.board[j][i] = cell.Clone(pos)
		}
	}
}
