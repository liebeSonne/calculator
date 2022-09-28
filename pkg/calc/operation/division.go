package operation

func NewDivisionOperation() Operation {
	return &operationDivision{}
}

type operationDivision struct {
}

func (p *operationDivision) Calculate(operand1 *Operand, operand2 *Operand) *float64 {
	if operand1 == nil || operand2 == nil {
		return nil
	}
	val1 := (*operand1).Value()
	val2 := (*operand2).Value()
	if val1 == nil || val2 == nil {
		return nil
	}
	if *val2 == 0.0 {
		return nil
	}
	result := *val1 / *val2
	return &result
}
