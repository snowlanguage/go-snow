package snow

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

func (simplePos SimplePos) AsSEPos(file *File) *SEPos {
	return &SEPos{
		Start: *NewSimplePos(simplePos.Col, simplePos.Ln, simplePos.Idx),
		End:   *NewSimplePos(simplePos.Col, simplePos.Ln, simplePos.Idx),
		File:  file,
	}
}

func (simplePos SimplePos) CreateSEPos(end SimplePos, file *File) *SEPos {
	return &SEPos{
		Start: *NewSimplePos(simplePos.Col, simplePos.Ln, simplePos.Idx),
		End:   *NewSimplePos(end.Col, end.Ln, end.Idx),
		File:  file,
	}
}
