package mode

type GUI struct {
	assetPath string
}

func NewGUI(path string) *GUI {
	g := new(GUI)
	g.assetPath = ""
	return g
}

func (g *GUI) SetPath(path string) {
	g.assetPath = path
}

func (g *GUI) ProcessImg() {
	println("process img from ", g.assetPath)
}
