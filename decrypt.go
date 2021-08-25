package sym

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
	"os"
)

// DecryptBytes uses key to return the decrypted content of 'ciphertextBytes'.
func DecryptBytes(key, ciphertextBytes []byte) ([]byte, error) {
	if len(ciphertextBytes) == 0 || ciphertextBytes == nil {
		return nil, errors.New("empty payload")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create new block cipher: %w", err)
	}

	if len(ciphertextBytes) < aes.BlockSize {
		return nil, errors.New("cipher text does not contain a valid initialization vector length")
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

	plaintext, err := DecryptBytes([]byte(key), ciphertext)
	if err != nil {
		return fmt.Errorf("failed to encrypt %q: %w", in, err)
	}
	return os.WriteFile(out, plaintext, 0777)
}

// DecryptString uses key to return the decrypted string contents of ciphertext.
func DecryptString(key, ciphertext string) (string, error) {
	plaintextBytes, err := DecryptBytes([]byte(key), []byte(ciphertext))
	if err != nil {
		return "", err
	}
	return string(plaintextBytes), nil
}
