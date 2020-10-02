package game

type Game interface {
	start()
}

type game struct {
	player int
}

func newGame(player int) *game {
	g := new(game)
	g.player = player
	return g
}
