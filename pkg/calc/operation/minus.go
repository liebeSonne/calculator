package operation

func NewMinusOperation() Operation {
	return &operationMinus{}
}

type operationMinus struct {
}

func (p *operationMinus) Calculate(operand1 Operand, operand2 Operand) *float64 {
	if operand1 == nil || operand2 == nil {
		return nil
	}
	val1 := operand1.Value()
	val2 := operand2.Value()
	if val1 == nil || val2 == nil {
		return nil
	}
	result := *val1 - *val2
	return &result
}
