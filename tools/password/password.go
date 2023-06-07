package password

import (
	"golang.org/x/crypto/bcrypt"
)

// Compare .. 比较
func Compare(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// Generate .. 生成
func Generate(password string) (hash string, err error) {
	cost := bcrypt.DefaultCost
	hashByte, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	hash = string(hashByte[:])
	return hash, err
}
