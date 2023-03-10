package snow

type SnowErrType string

const (
	ILLEGAL_CHAR_ERROR                 SnowErrType = "Illegal character error"
	TRAILING_DOT_ERROR                 SnowErrType = "Trailing dot error"
	MULTIPLE_DOTS_ERROR                SnowErrType = "Multiple dots error"
	UNTERMINATED_INLINE_COMMENT_ERROR  SnowErrType = "Unterminated inline comment error"
	UNTERMINATED_STRING_ERROR          SnowErrType = "Unterminated string error"
	EXPECTED_TOKEN_ERROR               SnowErrType = "Expected token error"
	INVALID_TOKEN_TYPE_ERROR           SnowErrType = "Invalid token type error"
	INVALID_OP_TOKEN_ERROR             SnowErrType = "Invalid op token error"
	TOO_BIG_VALUE_ERROR                SnowErrType = "Too big value error"
	VALUE_ERROR                        SnowErrType = "Value error"
	DIVISION_BY_ZERO_ERROR             SnowErrType = "Division by zero error"
	UNEXPECTED_TOKEN_ERROR             SnowErrType = "Unexpected token error"
	VARIABLE_ALREADY_DECLARED_ERROR    SnowErrType = "Unexpected token error"
	UNDEFINED_VARIABLE_ERROR           SnowErrType = "Undefined variable error"
	CONSTANT_VARIABLE_ASSIGNMENT_ERROR SnowErrType = "Constant variable assignment error"
	INVALID_ATTRIBUTE_ERROR            SnowErrType = "Invalid attribute error"
	UNABLE_TO_ASSIGN_ATTRIBUTE_ERROR   SnowErrType = "Unable to assign attribute error"
	BREAK_OUTSIDE_OF_LOOP_ERROR        SnowErrType = "Break outside of loop error"
	RETURN_OUTSIDE_OF_LOOP_ERROR       SnowErrType = "Return outside of loop error"
	CONTINUE_OUTSIDE_OF_LOOP_ERROR     SnowErrType = "Continue outside of loop error"
	INVALID_CALL_ERROR                 SnowErrType = "Invalid call error"
	ARGUMENT_ERROR                     SnowErrType = "Argument error"
)
