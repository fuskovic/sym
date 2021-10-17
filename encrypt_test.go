package sym

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncrypt(t *testing.T) {
	t.Parallel()
	t.Run("string", func(t *testing.T) {
		t.Parallel()
		t.Run("OK", func(t *testing.T) {
			t.Parallel()
			key := randomStringOfLen(t, 16)
			expected := randomStringOfLen(t, 10)
			ciphertext, err := EncryptString(key, expected)
			require.NoError(t, err)
			require.NotEqual(t, expected, ciphertext)
		})
		t.Run("should fail if symmetric key length is invalid", func(t *testing.T) {
			t.Parallel()
			_, err := EncryptString("", "plaintext")
			require.Error(t, err)
		})
		t.Run("should fail if plaintext string is empty", func(t *testing.T) {
			t.Parallel()
			key := randomStringOfLen(t, 16)
			_, err := EncryptString(key, "")
			require.Equal(t, err, ErrEmptyPayload)
		})
	})
	t.Run("bytes", func(t *testing.T) {
		t.Parallel()
		t.Run("OK", func(t *testing.T) {
			t.Parallel()

			// encrypt
			key := randomStringOfLen(t, 16)
			expected := randomBytesOfLen(t, 11)
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
			_, err := EncryptBytes(invalidSymmetricKey, nil)
			require.Error(t, err)
		})
		t.Run("should fail if plaintext bytes are empty", func(t *testing.T) {
			t.Parallel()
			key := randomStringOfLen(t, 16)
			_, err := EncryptBytes(key, []byte{})
			require.Equal(t, err, ErrEmptyPayload)
		})
		t.Run("should fail if plaintext bytes are nil", func(t *testing.T) {
			t.Parallel()
			key := randomStringOfLen(t, 16)
			_, err := EncryptBytes(key, nil)
			require.Equal(t, err, ErrEmptyPayload)
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

			// assert the encrypted contents have been written to the outfile
			ciphertextBytes, err := os.ReadFile(outFilePath)
			require.NoError(t, err)
			require.NotNil(t, ciphertextBytes)
			require.NotEqual(t, expected, ciphertextBytes)
		})
		t.Run("should fail if symmetric key length is invalid", func(t *testing.T) {
			t.Parallel()
			expected := randomBytesOfLen(t, 10)
			inFilePath, outFilePath, cleanUp := setupTestFiles(t, expected)
			defer cleanUp()
			require.Error(t, EncryptFile("", inFilePath, outFilePath))
		})
		t.Run("should fail if file does not exist", func(t *testing.T) {
			t.Parallel()
			key := randomStringOfLen(t, 16)
			require.Error(t, EncryptFile(key, "doesntexist", "doesntexist"))
		})
	})
}
