package main

import (
	"strconv"
)

func isDigit(r rune) bool {
	if r >= 48 && r <= 57 {
		return true
	}
	return false
}

func digit(r rune) int {
	i, err := strconv.Atoi(string(r))
	if err != nil {
		return 0
	}
	return i
}

func isLetter(r rune) bool {
	// check capital letter
	if r >= 65 && r <= 90 {
		return true
	}

	// check lowercase letter
	if r >= 97 && r <= 122 {
		return true
	}

	return false
}
