package core

type Pos struct {
	x int
	y int
}

func NewPos(x int, y int) *Pos {
	p := new(Pos)
	p.x = x
	p.y = y
	return p
}

func (p *Pos) GetX() int {
	return p.x
}

func (p *Pos) GetY() int {
	return p.y
}

func (p *Pos) Clone() *Pos {
	return NewPos(p.x, p.y)
}
