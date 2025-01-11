package sym

import (
	"crypto/aes"
	"crypto/rand"
	"encoding/hex"
	"os"
	"slices"
)

var validKeySizes = []int{16, 24, 32}

// KeyGen generates a new key that can be used to encrypt/decrypt
// strings, byte-data, and files. If size is not 16, 32, or 32
// a non-nil error will be returned.
func KeyGen(size int) (string, error) {
	if !slices.Contains(validKeySizes, size) {
		return "", aes.KeySizeError(size)
	}
	key := make([]byte, size)
	if _, err := rand.Read(key); err != nil {
		return "", err
	}
	return hex.EncodeToString(key), nil
}

// MustKeyGen always generates a size 16 key.
func MustKeyGen() string {
	const size = 16
	key := make([]byte, size)
	rand.Read(key)
	return string(key)
}

// KeyFromFile reads the file at path, validates the key,
// and then returns it. If the file does not exist a non-nil error is returned.
// If the file exists and the key is invalid a non-nil error is returned.
func KeyFromFilePath(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	if size := len(b); !slices.Contains(validKeySizes, size) {
		return "", aes.KeySizeError(size)
	}
	return string(b), nil
}
