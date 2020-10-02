package window

import "github.com/simp7/designWithGo/window/theme"

type state int

const (
	MINIMIZED state = 0
	MAXIMIZED
	CUSTOM
)

type Window interface {
	Maximize()
	Minimize()
	Reform(x int, y int, width int, height int)
	Close()
}

type OS interface {
	GetFilePath() string
	GetMinX() int
	GetMinY() int
	GetHeight() int
	GetWidth() int
}

type Mode interface {
	Show()
}

type window struct {
	Window
	theme.Theme
	OS
} //일치하는 인터페이스를 선언했다면 클래스 이름은 소문자로 시작하는 것이 안전

func NewWindow(os OS, theme theme.Theme) *window {
	w := new(window)
	w.OS = os
	w.Theme = theme
	return w
}
