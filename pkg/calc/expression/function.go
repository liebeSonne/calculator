package expression

import "calculator/pkg/calc/operation"

type Function interface {
	Computable
	ID() string
}

func NewLetFunction(id string, let *Computable) Function {
	return &functionLet{
		id:  id,
		let: let,
	}
}

type functionLet struct {
	id  string
	let *Computable
}

func (f *functionLet) ID() string {
	return f.id
}

func (f *functionLet) Value() *float64 {
	if f.let == nil {
		return nil
	}
	return (*f.let).Value()
}

func NewOperationFunction(id string, operand1 *Computable, operation operation.Operation, operand2 *Computable) Function {
	return &function{
		id:        id,
		operand1:  operand1,
		operation: operation,
		operand2:  operand2,
	}
}

type function struct {
	id        string
	operand1  *Computable
	operation operation.Operation
	operand2  *Computable
}

func (f *function) ID() string {
	return f.id
}

func (f *function) Value() *float64 {
	return f.operation.Calculate((operation.Operand)(*f.operand1), (operation.Operand)(*f.operand2))
}
