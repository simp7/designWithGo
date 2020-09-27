package window

import (
	"github.com/simp7/designWithGo/window/mode"
)

type CLIAdapter struct {
	mode.CLI
	OS
}

func NewCLIAdapter(os OS) *CLIAdapter {
	ca := new(CLIAdapter)
	ca.OS = os
	return ca
}

func (ca *CLIAdapter) Show() {
	ca.Print()
}

type GUIAdapter struct {
	mode.GUI
	OS
}

func NewGUIAdapter(os OS) *GUIAdapter {
	ga := new(GUIAdapter)
	ga.OS = os
	return ga
}

func (ga *GUIAdapter) Show() {
	ga.SetPath(ga.GetFilePath())
}
