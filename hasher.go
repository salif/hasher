package hasher

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"encoding/hex"
	"errors"

	"golang.org/x/crypto/argon2"
)

const VERSION = 2

func Hash(password string) (string, string, int) {
	var salt = newSalt(16)
	return toString(hashPasswordV2(password, salt)), toString(salt), VERSION
}

func Verify(password string, hash string, salt string, version int) bool {
	var toBytes = toBytesV2
	var hashPassword = hashPasswordV2
	if version > VERSION || version < 1 {
		panic(errors.New("version not supported"))
	}
	if version == 1 {
		toBytes = toBytesV1
		hashPassword = hashPasswordV1
	}
	return subtle.ConstantTimeCompare(
		hashPassword(password, toBytes(salt)),
		toBytes(hash)) == 1
}

func NeedsRehash(version int) bool {
	return version < VERSION
}

func hashPasswordV2(password string, salt []byte) []byte {
	return argon2.IDKey([]byte(password), salt, 4, 65536, 2, 32)
}

func hashPasswordV1(password string, salt []byte) []byte {
	return argon2.IDKey([]byte(password), salt, 1, 65536, 4, 32)
}

func newSalt(length int) []byte {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		panic(err)
	}
	return salt
}

func toString(b []byte) string {
	return hex.EncodeToString(b)
}

func toBytesV2(s string) []byte {
	var b, err = hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return b
}

func toBytesV1(s string) []byte {
	var b, err = base64.RawStdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return b
}
