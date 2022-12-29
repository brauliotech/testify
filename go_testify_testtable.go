package main

import (
	"fmt"
	"regexp"
	"unicode"
)

// IsAllLowerCase returns true if the given input contains only lowercase characters.
// Otherwise, returns false.
func IsAllLowerCase(input string) (bool, error) {
	re, err := regexp.Compile(`[0-9]+`) // regex for finding a digit in the input
	if err != nil {
		// returns an error if a digit is found
		return false, err
	}

	loc := re.FindStringIndex(input)
	if len(loc) > 0 {
		return false, fmt.Errorf("input must contain only characters")
	}

	for _, ch := range input {
		if unicode.IsUpper(ch) {
			return false, nil
		}
	}

	return true, nil
}
