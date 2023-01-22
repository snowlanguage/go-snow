package snowerror

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/snowlanguage/go-snow/position"
)

type SnowError struct {
	ErrType SnowErrType
	Msg     string
	Tip     string
	Pos     position.SEPos
}

func NewSnowError(errType SnowErrType, msg string, tip string, pos position.SEPos) *SnowError {
	return &SnowError{
		ErrType: errType,
		Msg:     msg,
		Tip:     tip,
		Pos:     pos,
	}
}

func (err SnowError) Error() string {
	tip := err.Tip
	if tip != "" {
		tip = tip + "\n"
	}

	codeAtLine := strings.Split(err.Pos.File.Code, "\n")[err.Pos.Start.Ln-1]
	add := len(strconv.Itoa(err.Pos.Start.Ln+1)) + 3
	arrows := strings.Repeat(" ", err.Pos.Start.Col+add) + strings.Repeat("^", err.Pos.End.Col-err.Pos.Start.Col+1)
	return fmt.Sprintf("\033[31m%s\033[0m: %s\n%s%d | %s\n%s", err.ErrType, err.Msg, tip, err.Pos.Start.Ln, codeAtLine, arrows)
}
