package os

type WindowOS struct {
}

func (w WindowOS) GetFilePath() string {
	panic("implement me")
}

func (w WindowOS) GetMinX() int {
	panic("implement me")
}

func (w WindowOS) GetMinY() int {
	panic("implement me")
}

func (w WindowOS) GetHeight() int {
	panic("implement me")
}

func (w WindowOS) GetWidth() int {
	panic("implement me")
}

func NewWindowOS() *WindowOS {
	w := new(WindowOS)
	return w
}
