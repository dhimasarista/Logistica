package utility

import (
	"strconv"
	"unicode"
)

// Fungsi untuk memeriksa apakah sebuah string berisi angka
func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// Fungsi untuk memeriksa apakah sebuah string berisi huruf
func IsAlphabetic(s string) bool {
	for _, char := range s {
		if !unicode.IsLetter(char) {
			return false
		}
	}
	return true
}
