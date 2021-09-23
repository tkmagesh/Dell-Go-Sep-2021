package calculator

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	//arrange
	x := 10
	y := 20
	expectedResult := 30

	//act
	result := Add(x, y)

	//assert
	/* if result != expectedResult {
		t.Fail()
	} */
	assert.Equal(t, expectedResult, result)

}

type CalculatorTestData struct {
	X              int
	Y              int
	ExpectedResult int
}

func TestCalculator(t *testing.T) {
	testData := []CalculatorTestData{
		{X: 10, Y: 20, ExpectedResult: 30},
		{X: 100, Y: 200, ExpectedResult: 300},
		{X: 1000, Y: 2000, ExpectedResult: 3000},
		{X: 10000, Y: 20000, ExpectedResult: 30000},
	}
	for _, td := range testData {
		t.Run(fmt.Sprintf("Testing Add for %d and %d", td.X, td.Y), func(t *testing.T) {
			result := Add(td.X, td.Y)
			/* if result != td.ExpectedResult {
				t.Fail()
			} */
			assert.Equal(t, td.ExpectedResult, result)
		})
	}
}
