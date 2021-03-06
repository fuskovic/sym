package sym

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

// EncryptBytes uses key to encrypt plaintextBytes; returning the ciphertext bytes.
func EncryptBytes(key string, plaintextBytes []byte) ([]byte, error) {
	if len(plaintextBytes) == 0 || plaintextBytes == nil {
		return nil, ErrEmptyPayload
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, fmt.Errorf("failed to create new aes cipher block: %w", err)
	}

	// Create a buffer big enough to hold the ciphertext + the IV.
	ciphertextBytes := make([]byte, aes.BlockSize+len(plaintextBytes))

	// First 16 bytes of the buffer will hold the IV.
	iv := ciphertextBytes[:aes.BlockSize]

	// Fill the IV with 16 random bytes
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, fmt.Errorf("failed to fill initialization vector: %w", err)
	}

	// Fill the rest of the buffer with the encrypted contents of plaintextBytes.
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertextBytes[aes.BlockSize:], plaintextBytes)
	return ciphertextBytes, nil
}

// EncryptFile opens in and uses key to encrypt it's plaintext contents; writing the ciphertext contents to out.
func EncryptFile(key, in, out string) error {
	plaintext, err := os.ReadFile(in)
	if err != nil {
		return fmt.Errorf("failed to read in file %q: %w", in, err)
	}

	ciphertext, err := EncryptBytes(key, plaintext)
	if err != nil {
		return fmt.Errorf("failed to encrypt %q: %w", in, err)
	}
	return os.WriteFile(out, ciphertext, 0777)
}

// EncryptString uses key to encrypt plaintext; returning the ciphertext string.
func EncryptString(key, plaintext string) (string, error) {
	plaintextBytes, err := EncryptBytes(key, []byte(plaintext))
	if err != nil {
		return "", err
	}
	return string(plaintextBytes), nil
}
