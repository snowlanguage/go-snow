package position

import "github.com/snowlanguage/go-snow/file"

type SEPos struct {
	Start SimplePos
	End   SimplePos
	File  *file.File
}

func NewSEP(start SimplePos, end SimplePos, file *file.File) *SEPos {
	return &SEPos{
		Start: start,
		End:   end,
		File:  file,
	}
}
