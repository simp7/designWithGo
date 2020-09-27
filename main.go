package main

import (
	"github.com/simp7/designWithGo/board"
	"github.com/simp7/designWithGo/window"
	"github.com/simp7/designWithGo/window/os"
)

func main() {

	selectedOS := os.NewWindowOS()
	window.NewWindow(selectedOS, window.NewGUIAdapter(selectedOS))

	var game = board.Get()
	game.Create(board.NewChessBuilder())
	game.Draw()

	//var a board.Cell
	//
	//a = board.NewChessCell(core.NewPos(0, 0))
	//a.Draw()
	//
	//b := a.Clone(core.NewPos(0, 1))
	//b.Draw()

}
