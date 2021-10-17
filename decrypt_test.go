package sym

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecrypt(t *testing.T) {
	t.Parallel()
	t.Run("string", func(t *testing.T) {
		t.Parallel()
		t.Run("OK", func(t *testing.T) {
			t.Parallel()

			// encrypt
			key := randomStringOfLen(t, 16)
			expected := randomStringOfLen(t, 10)
			ciphertext, err := EncryptString(key, expected)
			require.NoError(t, err)
			require.NotEqual(t, expected, ciphertext)

			// decrypt
			got, err := DecryptString(key, ciphertext)
			require.NoError(t, err)

			// assert equality
			require.Equal(t, expected, got)
		})
		t.Run("should fail if symmetric key length is invalid", func(t *testing.T) {
			t.Parallel()
			_, err := DecryptString("", "dummyciphertext")
			require.Error(t, err)
		})
		t.Run("should fail if plaintext string is empty", func(t *testing.T) {
			t.Parallel()
			key := randomStringOfLen(t, 16)
			_, err := DecryptString(key, "")
			require.Equal(t, err, ErrEmptyPayload)
		})
	})
	t.Run("bytes", func(t *testing.T) {
		t.Parallel()
		t.Run("OK", func(t *testing.T) {
			t.Parallel()

			// encrypt
			key := randomStringOfLen(t, 16)
			expected := randomBytesOfLen(t, 10)
			ciphertext, err := EncryptBytes(key, expected)
			require.NoError(t, err)
			require.NotEqual(t, expected, ciphertext)

			// decrypt
			got, err := DecryptBytes(key, ciphertext)
			require.NoError(t, err)

			// assert equality
			require.Equal(t, expected, got)
		})
		t.Run("should fail if symmetric key length is invalid", func(t *testing.T) {
			t.Parallel()
			invalidSymmetricKey := ""
			_, err := DecryptBytes(invalidSymmetricKey, []byte("dummyciphertext"))
			require.Error(t, err)
		})
		t.Run("should fail if ciphertext bytes are empty", func(t *testing.T) {
			t.Parallel()
			key := randomStringOfLen(t, 16)
			_, err := DecryptBytes(key, []byte{})
			require.Equal(t, err, ErrEmptyPayload)
		})
		t.Run("should fail if ciphertext bytes are nil", func(t *testing.T) {
			t.Parallel()
			key := randomStringOfLen(t, 16)
			_, err := DecryptBytes(key, nil)
			require.Equal(t, err, ErrEmptyPayload)
		})
		t.Run("should fail if ciphertext bytes is not at least the valid length of an initialization vector", func(t *testing.T) {
			t.Parallel()
			key := randomStringOfLen(t, 16)
			_, err := DecryptBytes(key, []byte("tooshort"))
			require.Equal(t, ErrInvalidIvLen, err)
		})
	})
	t.Run("file", func(t *testing.T) {
		t.Parallel()
		t.Run("OK", func(t *testing.T) {
			t.Parallel()

			// setup
			expected := randomBytesOfLen(t, 10)
			inFilePath, outFilePath, cleanUp := setupTestFiles(t, expected)
			defer cleanUp()

			// encrypt
			key := randomStringOfLen(t, 16)
			require.NoError(t, EncryptFile(key, inFilePath, outFilePath))

			// assert ciphertext has been written to out file
			ciphertextBytes, err := os.ReadFile(outFilePath)
			require.NoError(t, err)
			require.NotNil(t, ciphertextBytes)
			require.NotEqual(t, expected, ciphertextBytes)

			// decrypt
			decryptedFilePath := randomStringOfLen(t, 10) + "decrypted.txt"
			require.NoError(t, DecryptFile(key, outFilePath, decryptedFilePath))
			defer os.Remove(decryptedFilePath)

			// assert equality
			got, err := os.ReadFile(decryptedFilePath)
			require.NoError(t, err)
			require.Equal(t, expected, got)
		})
		t.Run("should fail if symmetric key length is invalid", func(t *testing.T) {
			t.Parallel()
			expected := randomBytesOfLen(t, 10)
			inFilePath, outFilePath, cleanUp := setupTestFiles(t, expected)
			defer cleanUp()
			require.Error(t, DecryptFile("", inFilePath, outFilePath))
		})
		t.Run("should fail if file does not exist", func(t *testing.T) {
			t.Parallel()
			key := randomStringOfLen(t, 16)
			require.Error(t, DecryptFile(key, "doesntexist", "doesntexist"))
		})
	})
}
