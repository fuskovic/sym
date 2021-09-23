package sym

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const validSymmetricKey = "rand16CharString"

func TestDecrypt(t *testing.T) {
	t.Parallel()
	t.Run("string", func(t *testing.T) {
		t.Parallel()
		t.Run("OK", func(t *testing.T) {
			t.Parallel()
			expected := "skafiskafnjak"

			// encrypt
			ciphertext, err := EncryptString(validSymmetricKey, expected)
			require.NoError(t, err)
			require.NotEqual(t, expected, ciphertext)

			// decrypt
			got, err := DecryptString(validSymmetricKey, ciphertext)
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
			_, err := DecryptString(validSymmetricKey, "")
			require.Error(t, err)
		})
	})
	t.Run("bytes", func(t *testing.T) {
		t.Parallel()
		t.Run("OK", func(t *testing.T) {
			t.Parallel()

			expected := []byte("skafiskafnjak")

			// encrypt
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
			_, err := DecryptBytes(invalidSymmetricKey, []byte("dummyciphertext"))
			require.Error(t, err)
		})
		t.Run("should fail if ciphertext bytes are empty", func(t *testing.T) {
			t.Parallel()
			_, err := DecryptBytes(validSymmetricKey, []byte{})
			require.Error(t, err)
		})
		t.Run("should fail if ciphertext bytes are nil", func(t *testing.T) {
			t.Parallel()
			_, err := DecryptBytes(validSymmetricKey, nil)
			require.Error(t, err)
		})
		t.Run("should fail if ciphertext bytes is not at least the valid length of an initialization vector", func(t *testing.T) {
			t.Parallel()
			_, err := DecryptBytes(validSymmetricKey, []byte("tooshort"))
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

			// assert ciphertext has been written to out file
			ciphertextBytes, err := os.ReadFile(outFilePath)
			require.NoError(t, err)
			require.NotNil(t, ciphertextBytes)
			require.NotEqual(t, expected, ciphertextBytes)

			// decrypt
			decryptedFilePath := randomStringOfLen(t, 10) + "decrypted.txt"
			require.NoError(t, DecryptFile(validSymmetricKey, outFilePath, decryptedFilePath))
			defer os.Remove(decryptedFilePath)

			// assert equality
			got, err := os.ReadFile(decryptedFilePath)
			require.NoError(t, err)
			require.Equal(t, expected, got)
		})
		t.Run("should fail if file does not exist", func(t *testing.T) {
			t.Parallel()
			require.Error(t, DecryptFile(validSymmetricKey, "doesntexist", "doesntexist"))
		})
	})
}
