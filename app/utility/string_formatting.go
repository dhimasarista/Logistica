package utility

import (
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Capitalize(i string) string {
	return toPascalCaseWithSpace(i)
}

func toPascalCaseWithSpace(input string) string {
	// Pisahkan string berdasarkan spasi
	words := strings.Fields(input)

	// Buat instance dari TitleCase dengan bahasa Inggris (English)
	titleCase := cases.Title(language.English)

	// Ubah setiap kata menjadi PascalCase
	for i, word := range words {
		// Konversi huruf pertama menjadi huruf besar
		words[i] = titleCase.String(word)
	}

	// Gabungkan kembali kata-kata menjadi string dengan spasi
	result := strings.Join(words, " ")

	return result
}

func ToLowerCase(str string) string {
	var result strings.Builder

	for _, char := range str {
		// Mengabaikan spasi dan hanya mengonversi huruf
		if unicode.IsLetter(char) {
			charLower := unicode.ToLower(char)
			result.WriteRune(charLower)
		}
	}

	return result.String()
}
