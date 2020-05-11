package hasher

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/argon2"
)

// Returns hash with length 43 and salt with length 22
func Hash(password string) (string, string, error) {
	salt, err := generateSalt(16)
	if err != nil {
		return "", "", err
	}
	hash := hashPassword(password, salt)
	return toString(hash), toString(salt), nil
}

func Verify(password string, hash string, salt string) (bool, error) {
	saltB, err := toBytes(salt)
	if err != nil {
		return false, err
	}
	hashB := hashPassword(password, saltB)
	return hash == toString(hashB), nil
}

func generateSalt(length int) ([]byte, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func hashPassword(password string, salt []byte) []byte {
	return argon2.IDKey([]byte(password), salt, 1, 65536, 4, 32)
}

func toString(b []byte) string {
	return base64.RawStdEncoding.EncodeToString(b)
}

func toBytes(s string) ([]byte, error) {
	result, err := base64.RawStdEncoding.DecodeString(s)
	return result, err
}
