package main

import "unicode"

func isPalindrome(s string) bool {
	letterRunes := []rune{}

	for _, str := range s {
		if unicode.IsLetter(str) {
			lowerLetter := unicode.ToLower(str)
			letterRunes = append(letterRunes, lowerLetter)
		} else if unicode.IsNumber(str) {
			letterRunes = append(letterRunes, str)
		}
	}

	finalIndex := len(letterRunes) - 1

	if finalIndex >= 0 {
        halfway := finalIndex / 2
		for i := 0; i <= halfway; i++ {
			if letterRunes[i] != letterRunes[finalIndex-i] {
				return false
			}
		}
	}

	return true
}
