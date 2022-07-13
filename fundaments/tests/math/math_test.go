package math

import "testing"

const errorMessage = "expected value %v, but the result is %v"

func TestAverage(t *testing.T) {
	t.Parallel()
	expectedValue := 5.0
	actualValue := Average(1, 2, 3, 4, 5, 6, 7, 8, 9)

	if actualValue != expectedValue {
		t.Errorf(errorMessage, expectedValue, actualValue)
	}
}
