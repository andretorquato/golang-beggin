package address

import "strings"

func AddressType(address string) string {
	validTypes := []string{"street", "home", "work", "other"}

	wordKey := strings.Split(address, " ")[0]
	wordKey = strings.ToLower(wordKey)

	isValid := false
	for _, validType := range validTypes {
		if wordKey == validType {
			isValid = true
		}
	}

	if isValid {
		return wordKey
	}

	return "other"
}
