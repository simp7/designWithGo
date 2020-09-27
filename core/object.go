package core

type Object interface {
	Add(Object)
	Clone(*Pos) Object
	GetChild(int) Object
	Highlight()
	OnClick()
	Draw()
	Remove(Object)
}

type object struct {
	Object
	child *object
}

type objects struct {
	groupAlgorithm
}

type groupAlgorithm interface {
	find(Object)
	get() Object
}

func NewObject() *object {
	o := new(object)
	return o
}
