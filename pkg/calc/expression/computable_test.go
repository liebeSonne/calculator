package expression

import "calculator/pkg/calc/operation"

func newMockComputableNil() Computable {
	return &mockComputable{
		value: nil,
	}
}

func newMockComputable(value float64) Computable {
	return &mockComputable{
		value: &value,
	}
}

type mockComputable struct {
	value *float64
}

func (m mockComputable) Value() *float64 {
	return m.value
}

func newMockOperationPlus() operation.Operation {
	return &mockOperation{}
}

type mockOperation struct {
}

func (m mockOperation) Calculate(operand1 operation.Operand, operand2 operation.Operand) *float64 {
	if operand1 == nil || operand2 == nil {
		return nil
	}
	val1 := operand1.Value()
	val2 := operand2.Value()
	if val1 == nil || val2 == nil {
		return nil
	}
	result := *val1 + *val2
	return &result
}

func mockValue(value float64) *float64 {
	return &value
}
