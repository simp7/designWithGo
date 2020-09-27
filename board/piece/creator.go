package piece

import "github.com/simp7/designWithGo/core"

type Creator interface {
	Create(string, *core.Pos)
}
