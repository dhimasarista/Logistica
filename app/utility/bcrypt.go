package utility

import "golang.org/x/crypto/bcrypt"

//  GenerateHash menghasilkan hash dari password dengan menggunakan bcrypt
func GenerateHash(value string) (string, error) {
	// GenerateFromPassword akan menghasilkan hash password dan menambahkan salt secara otomatis
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// ValidatePassword memeriksa apakah password yang dimasukkan sesuai dengan hash yang disimpan
func ValidateTextHashed(hashedPassword, textToCompare string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(textToCompare))
	return err == nil
}
