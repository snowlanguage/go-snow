package snow

type RTValue interface {
	ToString() string
	ValueToString() string
	GetType() RTType
	GetValue() interface{}
	GetEnvironment() *Environment
	Dot(other Token, position SEPos) (RTValue, error)
	SetAttribute(other string, value RTValue, position SEPos) (RTValue, error)
	Add(other RTValue, position SEPos) (RTValue, error)
	Subtract(other RTValue, position SEPos) (RTValue, error)
	Multiply(other RTValue, position SEPos) (RTValue, error)
	Divide(other RTValue, position SEPos) (RTValue, error)
	Equals(other RTValue, position SEPos) (RTValue, error)
	NotEquals(other RTValue, position SEPos) (RTValue, error)
	GreaterThan(other RTValue, position SEPos) (RTValue, error)
	GreaterThanEquals(other RTValue, position SEPos) (RTValue, error)
	LessThan(other RTValue, position SEPos) (RTValue, error)
	LessThanEquals(other RTValue, position SEPos) (RTValue, error)
	Not(position SEPos) (RTValue, error)
	ToBool(SEPos) (RTValue, error)
	Call(arguments []RTValue, position SEPos, interpreter *Interpreter) (RTValue, error)
}
