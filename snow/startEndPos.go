package snow

type SEPos struct {
	Start SimplePos
	End   SimplePos
	File  *File
}

func NewSEP(start SimplePos, end SimplePos, file *File) *SEPos {
	return &SEPos{
		Start: start,
		End:   end,
		File:  file,
	}
}
