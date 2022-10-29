package address

import (
	"testing"
)

type CasesTest struct {
	inputAddressType    string
	expectedAddressType string
}

func TestAddressType(t *testing.T) {
	t.Parallel()
	cases := []CasesTest{
		{"street one", "street"},
		{"home", "home"},
		{"work", "work"},
		{"unknown", "other"},
		{"unknown", "other"},
	}

	for _, c := range cases {
		output := AddressType(c.inputAddressType)
		if output != c.expectedAddressType {
			t.Errorf("AddressType(%q) == %q, want %q", c.inputAddressType, output, c.expectedAddressType)
		}
	}
}
