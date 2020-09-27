package mode

type CLI struct {
	background int
	foreground int
	font       string
}

func NewCLI(bg int, fg int, font string) *CLI {
	c := new(CLI)
	c.background = bg
	c.foreground = fg
	c.font = font
	return c
}

func (c *CLI) Print() {
	println("BackgroundColor: ", c.background, "ForegroundColor: ", c.foreground, "font: ", c.font)
}
