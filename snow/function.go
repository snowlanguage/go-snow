package snow

import (
	"fmt"
)

type RTFunction struct {
	Name        string
	Parameters  []Token
	Block       Stmt
	Pos         SEPos
	Environment *Environment
}

func NewRTFunction(name string, parameters []Token, block Stmt, pos SEPos, env *Environment) *RTFunction {
	return &RTFunction{
		Name:        name,
		Parameters:  parameters,
		Block:       block,
		Pos:         pos,
		Environment: env,
	}
}

func (rTFunction *RTFunction) ToString() string {
	return fmt.Sprintf("(FUNCTION: %s)", rTFunction.Name)
}

func (rTFunction *RTFunction) ValueToString() string {
	return fmt.Sprintf("FUNCTION %s", rTFunction.Name)
}

func (rTFunction *RTFunction) GetType() RTType {
	return RTT_FUNCTION
}

func (rTFunction *RTFunction) GetValue() interface{} {
	return fmt.Sprintf("FUNCTION %s", rTFunction.Name)
}

func (rTFunction *RTFunction) GetEnvironment() *Environment {
	return rTFunction.Environment
}

func (rTFunction *RTFunction) Dot(other Token, position SEPos) (RTValue, error) {
	return nil, NewInvalidAttributeRTError(rTFunction, other, position, rTFunction.Environment)
}

func (rTFunction *RTFunction) SetAttribute(other string, value RTValue, position SEPos) (RTValue, error) {
	return nil, NewUnableToAssignAttributeRTError(rTFunction, other, value, position, rTFunction.Environment)
}

func (rTFunction *RTFunction) Add(other RTValue, position SEPos) (RTValue, error) {
	return nil, NewValueRTError(
		PLUS,
		rTFunction,
		other,
		position,
		rTFunction.Environment,
	)
}

func (rTFunction *RTFunction) Subtract(other RTValue, position SEPos) (RTValue, error) {
	return nil, NewValueRTError(
		DASH,
		rTFunction,
		other,
		position,
		rTFunction.Environment,
	)
}

func (rTFunction *RTFunction) Multiply(other RTValue, position SEPos) (RTValue, error) {
	return nil, NewValueRTError(
		STAR,
		rTFunction,
		other,
		position,
		rTFunction.Environment,
	)
}

func (rTFunction *RTFunction) Divide(other RTValue, position SEPos) (RTValue, error) {
	return nil, NewValueRTError(
		SLASH,
		rTFunction,
		other,
		position,
		rTFunction.Environment,
	)
}

func (rTFunction *RTFunction) Equals(other RTValue, position SEPos) (RTValue, error) {
	if other.GetType() == RTT_FUNCTION && other.(*RTFunction).Name == rTFunction.Name && other.(*RTFunction).Environment == rTFunction.Environment {
		return NewRTBool(position, true, rTFunction.Environment), nil
	}

	return NewRTBool(position, false, rTFunction.Environment), nil
}

func (rTFunction *RTFunction) NotEquals(other RTValue, position SEPos) (RTValue, error) {
	if other.GetType() == RTT_FUNCTION && other.(*RTFunction).Name == rTFunction.Name && other.(*RTFunction).Environment == rTFunction.Environment {
		return NewRTBool(position, false, rTFunction.Environment), nil
	}

	return NewRTBool(position, true, rTFunction.Environment), nil
}

func (rTFunction *RTFunction) GreaterThan(other RTValue, position SEPos) (RTValue, error) {
	return nil, NewValueRTError(
		GREATER_THAN,
		rTFunction,
		other,
		position,
		rTFunction.Environment,
	)
}

func (rTFunction *RTFunction) GreaterThanEquals(other RTValue, position SEPos) (RTValue, error) {
	return nil, NewValueRTError(
		GREATER_THAN_EQUALS,
		rTFunction,
		other,
		position,
		rTFunction.Environment,
	)
}

func (rTFunction *RTFunction) LessThan(other RTValue, position SEPos) (RTValue, error) {
	return nil, NewValueRTError(
		LESS_THAN,
		rTFunction,
		other,
		position,
		rTFunction.Environment,
	)
}

func (rTFunction *RTFunction) LessThanEquals(other RTValue, position SEPos) (RTValue, error) {
	return nil, NewValueRTError(
		LESS_THAN_EQUALS,
		rTFunction,
		other,
		position,
		rTFunction.Environment,
	)
}

func (rTFunction *RTFunction) Not(position SEPos) (RTValue, error) {
	return NewRTBool(position, false, rTFunction.Environment), nil
}

func (rTFunction *RTFunction) ToBool(position SEPos) (RTValue, error) {
	return NewRTBool(position, true, rTFunction.Environment), nil
}

func (rTFunction *RTFunction) Call(arguments []RTValue, position SEPos, interpreter *Interpreter) (RTValue, error) {
	fmt.Println(arguments, rTFunction.Parameters)
	if len(arguments) > len(rTFunction.Parameters) {
		return nil, NewTooManyArgumentsRTError(rTFunction, len(rTFunction.Parameters), len(arguments), position, interpreter.environment)
	} else if len(arguments) < len(rTFunction.Parameters) {
		return nil, NewTooFewArgumentsRTError(rTFunction, len(rTFunction.Parameters), len(arguments), position, interpreter.environment)
	}

	runEnv := NewEnvironment(rTFunction.Environment, rTFunction.Name, rTFunction.Pos.Start.Ln, rTFunction.Pos.File.Name, false)

	for index, v := range arguments {
		err := runEnv.Declare(false, rTFunction.Parameters[index].Value, v, position)
		if err != nil {
			return nil, err
		}
	}

	_, err := interpreter.execute(rTFunction.Block, runEnv)
	if err != nil {
		return nil, err
	}

	val := interpreter.returnVal
	interpreter.returnVal = nil

	return val, nil
}
