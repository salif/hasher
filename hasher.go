package hasher

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"

	"golang.org/x/crypto/argon2"
)

// Returns hash with length 43 and salt with length 43
func Hash(password string) (string, string, error) {
	salt, err := newSalt(32)
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
	passwordHash := hashPassword(password, saltB)
	hashB, err := toBytes(hash)
	if err != nil {
		return false, err
	}
	return subtle.ConstantTimeCompare(passwordHash, hashB) == 1, nil
}

func newSalt(length int) ([]byte, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	return salt, err
}

func hashPassword(password string, salt []byte) []byte {
	return argon2.IDKey([]byte(password), salt, 1, 65536, 4, 32)
}

func toString(b []byte) string {
	return base64.RawStdEncoding.EncodeToString(b)
}

func toBytes(s string) ([]byte, error) {
	return base64.RawStdEncoding.DecodeString(s)
}
