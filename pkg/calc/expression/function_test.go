package expression

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"calculator/pkg/calc/operation"
)

func TestFunction_Value(t *testing.T) {
	testCases := []struct {
		testName  string
		id        string
		operand1  Computable
		operation operation.Operation
		operand2  Computable
		result    *float64
	}{
		{
			"plus for two operand",
			"f",
			newMockComputable(3.0),
			newMockOperationPlus(),
			newMockComputable(4.0),
			mockValue(7.0),
		},
		{
			"plus for nil and nil value",
			"f",
			newMockComputableNil(),
			newMockOperationPlus(),
			newMockComputableNil(),
			nil,
		},
		{
			"plus for nil and nil operand",
			"f",
			nil,
			newMockOperationPlus(),
			nil,
			nil,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		f := func(t *testing.T) {
			fn := NewOperationFunction(testCase.id, &testCase.operand1, testCase.operation, &testCase.operand2)

			assert.Equal(t, testCase.id, fn.ID())

			value := fn.Value()

			if testCase.result == nil {
				assert.Nil(t, value)
			} else {
				assert.NotNil(t, value)
				assert.Equal(t, testCase.result, value)
			}
		}
		t.Run(testCase.testName, f)
	}
}
