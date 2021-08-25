package sym

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
	"os"
)

// ErrEmptyPayload is returned when the ciphertext string or bytes provided for an encrypt/decrypt operation are empty or nil.
var ErrEmptyPayload = errors.New("empty payload")

// ErrInvalidIvLen is returned when the provided ciphertext for a decrypt operation is not at least the length of a valid initialization vector.
var ErrInvalidIvLen = errors.New("cipher text does not contain a valid initialization vector length")

// DecryptBytes uses key to return the decrypted content of 'ciphertextBytes'.
func DecryptBytes(key string, ciphertextBytes []byte) ([]byte, error) {
	if len(ciphertextBytes) == 0 || ciphertextBytes == nil {
		return nil, ErrEmptyPayload
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, fmt.Errorf("failed to create new block cipher: %w", err)
	}

	if len(ciphertextBytes) < aes.BlockSize {
		return nil, ErrInvalidIvLen
	}

	// Get the 16 byte IV.
	iv := ciphertextBytes[:aes.BlockSize]

	// Remove it.
	ciphertextBytes = ciphertextBytes[aes.BlockSize:]

	// Return decrypted contents.
	plaintext := make([]byte, len(ciphertextBytes))
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(plaintext, ciphertextBytes)
	return plaintext, nil
}

// DecryptFile opens in and uses key to decrypt it's ciphertext contents; writing the plaintext contents to out.
func DecryptFile(key, in, out string) error {
	ciphertext, err := os.ReadFile(in)
	if err != nil {
		return fmt.Errorf("failed to read in file %q: %w", in, err)
	}

	plaintext, err := DecryptBytes(key, ciphertext)
	if err != nil {
		return fmt.Errorf("failed to encrypt %q: %w", in, err)
	}
	return os.WriteFile(out, plaintext, 0777)
}

// DecryptString uses key to return the decrypted string contents of ciphertext.
func DecryptString(key, ciphertext string) (string, error) {
	plaintextBytes, err := DecryptBytes(key, []byte(ciphertext))
	if err != nil {
		return "", err
	}
	return string(plaintextBytes), nil
}
