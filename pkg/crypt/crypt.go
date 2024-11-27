package crypt

import "golang.org/x/crypto/bcrypt"

// Шифрует токен
func CryptRefreshToken(token string) ([]byte, error) {
	tokenHash, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	if err != nil {
		return []byte{}, err
	}
	return tokenHash, nil
}
