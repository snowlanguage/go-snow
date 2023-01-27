package runtimevalues

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/snowlanguage/go-snow/position"
	snowerror "github.com/snowlanguage/go-snow/snowError"
	"github.com/snowlanguage/go-snow/token"
)

type RTError struct {
	snowerror.SnowError
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

func NewRuntimeError(errType snowerror.SnowErrType, msg string, tip string, pos position.SEPos, env *Environment) *RTError {
	return &RTError{
		SnowError: *snowerror.NewSnowError(
			errType,
			msg,
			tip,
			pos,
		),
		environment: env,
	}
}

func NewValueRTError(opTType token.TokenType, x RTValue, y RTValue, pos position.SEPos, env *Environment) *RTError {
	var msg string
	var op string
	var withBy string

	switch opTType {
	case token.PLUS:
		op = "add"
		withBy = "to"
	case token.DASH:
		op = "subtract"
		withBy = "by"
	case token.STAR:
		op = "multiply"
		withBy = "by"
	case token.SLASH:
		op = "divide"
		withBy = "by"
	case token.EQUALS:
		op = "check equality between"
		withBy = "and"
	case token.NOT_EQUALS:
		op = "check inequality between"
		withBy = "and"
	case token.GREATER_THAN:
		op = "compare sizes between"
		withBy = "and"
	case token.GREATER_THAN_EQUALS:
		op = "compare sizes between"
		withBy = "and"
	case token.LESS_THAN:
		op = "compare sizes between"
		withBy = "and"
	case token.LESS_THAN_EQUALS:
		op = "compare sizes between"
		withBy = "and"
	}

	if y != nil {
		msg = fmt.Sprintf("unable to %s '%s' with value of '%s' %s '%s' with value of '%s'", op, x.GetType(), x.ValueToString(), withBy, y.GetType(), y.ValueToString())
	} else {
		msg = fmt.Sprintf("unable to %s '%s' with value of '%s'", op, x.GetType(), x.ValueToString())
	}

	return &RTError{
		SnowError: *snowerror.NewSnowError(
			snowerror.VALUE_ERROR,
			msg,
			"",
			pos,
		),
		environment: env,
	}
}

func NewDivisionByZeroRTError(x RTValue, y RTValue, pos position.SEPos, env *Environment) *RTError {
	return &RTError{
		SnowError: *snowerror.NewSnowError(
			snowerror.VALUE_ERROR,
			fmt.Sprintf("unable to divide '%s' with value of '%s' by '%s' with value of '%s'", x.GetType(), x.ValueToString(), y.GetType(), y.ValueToString()),
			"",
			pos,
		),
		environment: env,
	}
}

func NewInvalidAttributeRTError(x RTValue, y token.Token, pos position.SEPos, env *Environment) *RTError {
	return &RTError{
		SnowError: *snowerror.NewSnowError(
			snowerror.INVALID_ATTRIBUTE_ERROR,
			fmt.Sprintf("object of type '%s' has no attribute called '%s'", x.GetType(), y.Value),
			"",
			pos,
		),
		environment: env,
	}
}

func NewUnableToAssignAttributeError(x RTValue, y string, val RTValue, pos position.SEPos, env *Environment) *RTError {
	return &RTError{
		SnowError: *snowerror.NewSnowError(
			snowerror.UNABLE_TO_ASSIGN_ATTRIBUTE_ERROR,
			fmt.Sprintf("unable to assign '%s' of '%s' to a '%s' with value of '%s'", y, x.GetType(), val.GetType(), val.ValueToString()),
			"",
			pos,
		),
		environment: env,
	}
}
