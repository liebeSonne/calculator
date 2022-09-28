package operation

func NewMultiplyOperation() Operation {
	return &operationMultiply{}
}

type operationMultiply struct {
}

func (p *operationMultiply) Calculate(operand1 *Operand, operand2 *Operand) *float64 {
	if operand1 == nil || operand2 == nil {
		return nil
	}
	val1 := (*operand1).Value()
	val2 := (*operand2).Value()
	if val1 == nil || val2 == nil {
		return nil
	}
	result := *val1 * *val2
	return &result
}
