package handlers

import "golang.org/x/crypto/bcrypt"

var jwtSecret = []byte("JWT_SECRET")

func HashPassword(password string) (string, error) {
  bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
  return string(bytes), err
}

func VerifyPassword(hashedPassword, password string) error {
  return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
