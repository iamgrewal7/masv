package util

import "fmt"

// ValidateInput validates each input in the list by
// verifying the first character to be between '0' and '9' and rest be between 'a' and 'z'
func ValidateInput(inputList []string) error {
	for i, s := range inputList {
		if s[0] < '0' || s[0] > '9' {
			return fmt.Errorf("expected a digit at the beginning of the argument no. %d (%s)", i+1, s)
		}

		if len(s) == 1 {
			return fmt.Errorf("expected a lowercase letter after digit in the argument no. %d (%s)", i+1, s)
		}

		for j, v := range s[1:] {
			if v < 'a' || v > 'z' {
				return fmt.Errorf("expected a lowercase letter at index %d in the argument no. %d (%s)", j+1, i+1, s)
			}
		}
	}
	return nil
}

// Rotate function increases each letter by 'num' positions
func Rotate(letter rune, num rune) rune {
	letter += num
	if letter > 'z' {
		letter -= 26
	}
	return letter
}
