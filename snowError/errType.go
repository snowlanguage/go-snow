package snowerror

type SnowErrType string

const (
	ILLEGAL_CHAR_ERROR                SnowErrType = "Illegal character error"
	TRAILING_DOT_ERROR                SnowErrType = "Trailing dot error"
	MULTIPLE_DOTS_ERROR               SnowErrType = "Multiple dots error"
	UNTERMINATED_INLINE_COMMENT_ERROR SnowErrType = "Unterminated inline comment error"
	UNTERMINATED_STRING_ERROR         SnowErrType = "Unterminated string error"
	EXPECTED_TOKEN_ERROR              SnowErrType = "Expected token error"
	INVALID_TOKEN_TYPE                SnowErrType = "Invalid token type"
)
