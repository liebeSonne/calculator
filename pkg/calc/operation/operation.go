package operation

import "calculator/pkg/calc/computable"

type Operand interface {
	computable.Computable
}

type Operation interface {
	Calculate(operand1 *Operand, operand2 *Operand) *float64
}
