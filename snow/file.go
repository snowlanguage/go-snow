package snow

type File struct {
	Name string
	Code string
}

func NewFile(name string, code string) *File {
	return &File{
		Name: name,
		Code: code,
	}
}
