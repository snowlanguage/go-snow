package runtimevalues

import (
	"fmt"

	"github.com/snowlanguage/go-snow/position"
	snowerror "github.com/snowlanguage/go-snow/snowError"
	"github.com/snowlanguage/go-snow/token"
)

type RTError struct {
	snowerror.SnowError
	environment *Environment
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
	}

	if y != nil {
		msg = fmt.Sprintf("unable to %s '%s' with value of '%s' to %s %s value of '%s'", op, x.GetType(), x.ValueToString(), withBy, y.GetType(), y.ToString())
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