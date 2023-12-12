package utils

import (
	"crypto/sha1"
	"encoding/hex"
)

const salt = "hjqrhjqw124617ajfhajs"

func GeneratePasswordHash(password string) string {
	hasher := sha1.New()
	hasher.Write([]byte(password))
	hashedPassword := hex.EncodeToString(hasher.Sum([]byte(salt)))
	return hashedPassword
}
