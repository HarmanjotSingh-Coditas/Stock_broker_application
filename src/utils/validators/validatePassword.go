package validators

import "unicode"

func ValidatePassword(s string) bool {
	if len(s) < 8 || len(s) > 20 { 
		return false
	}

	var hasUpper, hasLower, hasNumber, hasSpecial bool

	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
		if hasUpper && hasLower && hasNumber && hasSpecial {
			return true
		}
	}
	return hasUpper && hasLower && hasNumber && hasSpecial
}
