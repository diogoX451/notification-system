package pkg

import "golang.org/x/crypto/bcrypt"

func HashPassword(value string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	return string(hashed)
}
