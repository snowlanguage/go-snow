package snowerror

type SnowErrType string

const (
	ILLEGAL_CHAR_ERROR  SnowErrType = "Illegal character error"
	TRAILING_DOT_ERROR  SnowErrType = "Trailing dot error"
	MULTIPLE_DOTS_ERROR SnowErrType = "Multiple dots error"
)
