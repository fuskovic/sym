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
			expected := randomStringOfLen(t, 10)
			ciphertext, err := EncryptString(validSymmetricKey, expected)
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
			_, err := EncryptString(validSymmetricKey, "")
			require.Error(t, err)
		})
	})
	t.Run("bytes", func(t *testing.T) {
		t.Parallel()
		t.Run("OK", func(t *testing.T) {
			t.Parallel()

			// encrypt
			expected := randomBytesOfLen(t, 10)
			ciphertext, err := EncryptBytes(validSymmetricKey, expected)
			require.NoError(t, err)
			require.NotEqual(t, expected, ciphertext)

			// decrypt
			got, err := DecryptBytes(validSymmetricKey, ciphertext)
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
			_, err := EncryptBytes(validSymmetricKey, []byte{})
			require.Error(t, err)
		})
		t.Run("should fail if plaintext bytes are nil", func(t *testing.T) {
			t.Parallel()
			_, err := EncryptBytes(validSymmetricKey, nil)
			require.Error(t, err)
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
			require.NoError(t, EncryptFile(validSymmetricKey, inFilePath, outFilePath))

			// assert the encrypted contents have been written to the outfile
			ciphertextBytes, err := os.ReadFile(outFilePath)
			require.NoError(t, err)
			require.NotNil(t, ciphertextBytes)
			require.NotEqual(t, expected, ciphertextBytes)
		})
		t.Run("should fail if file does not exist", func(t *testing.T) {
			t.Parallel()
			require.Error(t, EncryptFile(validSymmetricKey, "doesntexist", "doesntexist"))
		})
	})
}
