package strings

import (
	"strings"
	"testing"
)

var errorMessage = "%s expected to be %d, but got %d"

func TestIndexStrings(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input         string
		value         string
		indexExpected int
	}{
		{"Hello World", "World", 6},
		{"Own", "Li", -1},
		{"Ruffy", "Ruffy", 0},
	}

	for _, test := range tests {
		t.Logf("Testing %s", test.input)
		current := strings.Index(test.input, test.value)
		if current != test.indexExpected {
			t.Errorf(errorMessage, test.input, current, test.indexExpected)
		}

	}
}
