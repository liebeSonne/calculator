package operation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperationDivision_Calculate(t *testing.T) {
	testCases := []struct {
		testName string
		operand1 Operand
		operand2 Operand
		result   *float64
	}{
		{
			"two positive",
			newMockValueOperand(8.0),
			newMockValueOperand(4.0),
			mockValue(2.0),
		},
		{
			"two negative",
			newMockValueOperand(-8.0),
			newMockValueOperand(-4.0),
			mockValue(2.0),
		},
		{
			"positive and negative",
			newMockValueOperand(8.0),
			newMockValueOperand(-4.0),
			mockValue(-2.0),
		},
		{
			"negative and positive",
			newMockValueOperand(-8.0),
			newMockValueOperand(4.0),
			mockValue(-2.0),
		},
		{
			"value and zero",
			newMockValueOperand(-8.0),
			newMockValueOperand(0.0),
			nil,
		},
		{
			"zero and value",
			newMockValueOperand(0.0),
			newMockValueOperand(4.0),
			mockValue(0.0),
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
			operation := NewDivisionOperation()
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
