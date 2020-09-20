package main

type Pos struct {
	X int
	Y int
}

func MakePos(x int, y int) *Pos {
	p := new(Pos)
	p.X = x
	p.Y = y
	return p
}

type Cell interface {
	CanMove() bool
}

type Wall struct {
	pos *Pos
	tileImg string
	wallImg string
}

func MakeWall(x int, y int, tileUrl string, wallUrl string) *Wall {
	w := new(Wall)
	w.pos = MakePos(x, y)
	w.tileImg = tileUrl
	w.wallImg = wallUrl
	return w
}

type GameMap struct {
	height int
	width int
	a [][]Cell
}

func (w *Wall) CanMove() bool {
	return false
}