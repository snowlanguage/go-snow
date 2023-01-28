package snow

import (
	"fmt"
	"strconv"
	"strings"
)

type RTError struct {
	SnowError
	environment *Environment
}

func (rTError RTError) Error() string {
	tip := rTError.Tip
	if tip != "" {
		tip = tip + "\n"
	}

	stack := make([]string, 0)
	env := rTError.environment
	for env != nil {
		name := "anonymous block"
		if env.Name != "" {
			name = fmt.Sprintf("'%s'", env.Name)
		}

		if env.IsFile {
			stack = append(stack, fmt.Sprintf("In file '%s'", env.FileName))
		} else {
			stack = append(stack, fmt.Sprintf("In %s starting at line %d in file '%s'", name, env.StartLine, env.FileName))
		}
		env = env.Parent
	}

	for i, j := 0, len(stack)-1; i < j; i, j = i+1, j-1 {
		stack[i], stack[j] = stack[j], stack[i]
	}

	codeAtLine := strings.ReplaceAll(strings.Split(rTError.Pos.File.Code, "\n")[rTError.Pos.Start.Ln-1], "\t", "   ")
	add := len(strconv.Itoa(rTError.Pos.Start.Ln+1)) + 3
	arrows := strings.Repeat(" ", rTError.Pos.Start.Col+add) + strings.Repeat("^", rTError.Pos.End.Col-rTError.Pos.Start.Col+1)
	return fmt.Sprintf("Stack with most recent last:\n%s\n\033[31m%s\033[0m: %s\n%s%d | %s\n%s", strings.Join(stack, "\n"), rTError.ErrType, rTError.Msg, tip, rTError.Pos.Start.Ln, codeAtLine, arrows)
}

func NewRuntimeError(errType SnowErrType, msg string, tip string, pos SEPos, env *Environment) *RTError {
	return &RTError{
		SnowError: *NewSnowError(
			errType,
			msg,
			tip,
			pos,
		),
		environment: env,
	}
}

func NewValueRTError(opTType TokenType, x RTValue, y RTValue, pos SEPos, env *Environment) *RTError {
	var msg string
	var op string
	var withBy string

	switch opTType {
	case PLUS:
		op = "add"
		withBy = "to"
	case DASH:
		op = "subtract"
		withBy = "by"
	case STAR:
		op = "multiply"
		withBy = "by"
	case SLASH:
		op = "divide"
		withBy = "by"
	case EQUALS:
		op = "check equality between"
		withBy = "and"
	case NOT_EQUALS:
		op = "check inequality between"
		withBy = "and"
	case GREATER_THAN:
		op = "compare sizes between"
		withBy = "and"
	case GREATER_THAN_EQUALS:
		op = "compare sizes between"
		withBy = "and"
	case LESS_THAN:
		op = "compare sizes between"
		withBy = "and"
	case LESS_THAN_EQUALS:
		op = "compare sizes between"
		withBy = "and"
	}

	if y != nil {
		msg = fmt.Sprintf("unable to %s '%s' with value of '%s' %s '%s' with value of '%s'", op, x.GetType(), x.ValueToString(), withBy, y.GetType(), y.ValueToString())
	} else {
		msg = fmt.Sprintf("unable to %s '%s' with value of '%s'", op, x.GetType(), x.ValueToString())
	}

	return &RTError{
		SnowError: *NewSnowError(
			VALUE_ERROR,
			msg,
			"",
			pos,
		),
		environment: env,
	}
}

func NewDivisionByZeroRTError(x RTValue, y RTValue, pos SEPos, env *Environment) *RTError {
	return &RTError{
		SnowError: *NewSnowError(
			VALUE_ERROR,
			fmt.Sprintf("unable to divide '%s' with value of '%s' by '%s' with value of '%s'", x.GetType(), x.ValueToString(), y.GetType(), y.ValueToString()),
			"",
			pos,
		),
		environment: env,
	}
}

func NewInvalidAttributeRTError(x RTValue, y Token, pos SEPos, env *Environment) *RTError {
	return &RTError{
		SnowError: *NewSnowError(
			INVALID_ATTRIBUTE_ERROR,
			fmt.Sprintf("object of type '%s' has no attribute called '%s'", x.GetType(), y.Value),
			"",
			pos,
		),
		environment: env,
	}
}

func NewUnableToAssignAttributeRTError(x RTValue, y string, val RTValue, pos SEPos, env *Environment) *RTError {
	return &RTError{
		SnowError: *NewSnowError(
			UNABLE_TO_ASSIGN_ATTRIBUTE_ERROR,
			fmt.Sprintf("unable to assign '%s' of '%s' to a '%s' with value of '%s'", y, x.GetType(), val.GetType(), val.ValueToString()),
			"",
			pos,
		),
		environment: env,
	}
}

func NewInvalidCallRTError(x RTValue, pos SEPos, env *Environment) *RTError {
	return &RTError{
		SnowError: *NewSnowError(
			INVALID_CALL_ERROR,
			fmt.Sprintf("unable to call object of type '%s'", x.GetType()),
			"",
			pos,
		),
		environment: env,
	}
}

func NewTooManyArgumentsRTError(x RTValue, expected int, got int, pos SEPos, env *Environment) *RTError {
	return &RTError{
		SnowError: *NewSnowError(
			ARGUMENT_ERROR,
			fmt.Sprintf("too may arguments, object of type '%s' expected %d arguments but got %d arguments", x.GetType(), expected, got),
			"",
			pos,
		),
		environment: env,
	}
}

func NewTooFewArgumentsRTError(x RTValue, expected int, got int, pos SEPos, env *Environment) *RTError {
	return &RTError{
		SnowError: *NewSnowError(
			ARGUMENT_ERROR,
			fmt.Sprintf("too few arguments, object of type '%s' expected %d arguments but got %d arguments", x.GetType(), expected, got),
			"",
			pos,
		),
		environment: env,
	}
}
