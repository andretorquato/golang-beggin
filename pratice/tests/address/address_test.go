package address

import "testing"

func TestAddressType(t *testing.T) {
	inputAddressType := "street queiroz"
	expectedAddressType := "street"
	receivedAddressType := AddressType(inputAddressType)
	if receivedAddressType != expectedAddressType {
		t.Errorf("AddressType(%s) == %s, expected %s", inputAddressType, receivedAddressType, expectedAddressType)
	}
}
