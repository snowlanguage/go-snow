package position

import "github.com/snowlanguage/go-snow/file"

type SimplePos struct {
	Col int
	Ln  int
	Idx int
}

func NewSimplePos(col int, ln int, idx int) *SimplePos {
	return &SimplePos{
		Col: col,
		Ln:  ln,
		Idx: idx,
	}
}

func (simplePos SimplePos) AsSEPos(file *file.File) *SEPos {
	return &SEPos{
		Start: *NewSimplePos(simplePos.Col, simplePos.Ln, simplePos.Idx),
		End:   *NewSimplePos(simplePos.Col, simplePos.Ln, simplePos.Idx),
		File:  file,
	}
}
