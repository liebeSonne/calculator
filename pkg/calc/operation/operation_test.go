package operation

func newMockValueOperand(value float64) Operand {
	return &mockOperand{
		value: &value,
	}
}

func newMockNilOperand() Operand {
	return &mockOperand{
		value: nil,
	}
}

type mockOperand struct {
	value *float64
}

func (m mockOperand) Value() *float64 {
	return m.value
}

func mockValue(value float64) *float64 {
	return &value
}
