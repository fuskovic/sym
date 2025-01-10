package sym

import (
	"crypto/aes"
	"crypto/rand"
	"slices"
)

// MustKeyGen generates a new key that can be used to encrypt/decrypt
// strings, byte-data, and files. If size is not 16, 32, or 32
// a non-nil error will be returned.
func KeyGen(size int) (string, error) {
	validKeySizes := []int{16, 24, 32}
	if !slices.Contains(validKeySizes, size) {
		return "", aes.KeySizeError(size)
	}

	key := make([]byte, 64)
	if _, err := rand.Read(key); err != nil {
		return "", err
	}
	return string(key), nil
}

// MustKeyGen always generates a size 16 key.
func MustKeyGen() string {
	const size = 16
	key := make([]byte, size)
	rand.Read(key)
	return string(key)
}
