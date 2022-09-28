package operation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperationPlus_Calculate(t *testing.T) {
	testCases := []struct {
		testName string
		operand1 Operand
		operand2 Operand
		result   *float64
	}{
		{
			"two positive",
			newMockValueOperand(3.0),
			newMockValueOperand(4.0),
			mockValue(7.0),
		},
		{
			"two negative",
			newMockValueOperand(-3.0),
			newMockValueOperand(-4.0),
			mockValue(-7.0),
		},
		{
			"positive and negative, negative result",
			newMockValueOperand(3.0),
			newMockValueOperand(-4.0),
			mockValue(-1.0),
		},
		{
			"positive and negative, positive result",
			newMockValueOperand(-3.0),
			newMockValueOperand(4.0),
			mockValue(1.0),
		},
		{
			"nil operand and value",
			nil,
			newMockValueOperand(4.0),
			nil,
		},
		{
			"value and nil operand",
			newMockValueOperand(3.0),
			nil,
			nil,
		},
		{
			"nil value and value",
			newMockNilOperand(),
			newMockValueOperand(4.0),
			nil,
		},
		{
			"value and nil value",
			newMockValueOperand(3.0),
			newMockNilOperand(),
			nil,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		f := func(t *testing.T) {
			operation := NewPlusOperation()
			result := operation.Calculate(testCase.operand1, testCase.operand2)

			if testCase.result == nil {
				assert.Nil(t, result)
			} else {
				assert.NotNil(t, result)
				assert.Equal(t, testCase.result, result)
			}
		}
		t.Run(testCase.testName, f)
	}
}

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
