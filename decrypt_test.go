package sym

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	validSymmetricKey      = "rand16CharString"
	validSymmetricKeyBytes = []byte(validSymmetricKey)
)

func TestDecrypt(t *testing.T) {
	t.Parallel()
	t.Run("string", func(t *testing.T) {
		t.Parallel()
		t.Run("OK", func(t *testing.T) {
			t.Parallel()
			// encrypt
			expected := "skafiskafnjak"
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
			_, err := DecryptString("", "")
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
			// encrypt
			expected := []byte("skafiskafnjak")
			ciphertext, err := EncryptBytes(validSymmetricKeyBytes, expected)
			require.NoError(t, err)
			require.NotEqual(t, expected, ciphertext)

			// decrypt
			got, err := DecryptBytes(validSymmetricKeyBytes, ciphertext)
			require.NoError(t, err)
			require.Equal(t, expected, got)
		})
		t.Run("should fail if symmetric key length is invalid", func(t *testing.T) {
			t.Parallel()
			invalidSymmetricKeyBytes := []byte("")
			_, err := DecryptBytes(invalidSymmetricKeyBytes, nil)
			require.Error(t, err)
		})
		t.Run("should fail if plaintext bytes are empty", func(t *testing.T) {
			t.Parallel()
			_, err := DecryptBytes(validSymmetricKeyBytes, []byte{})
			require.Error(t, err)
		})
		t.Run("should fail if plaintext bytes are nil", func(t *testing.T) {
			t.Parallel()
			_, err := DecryptBytes(validSymmetricKeyBytes, nil)
			require.Error(t, err)
		})
	})
	t.Run("file", func(t *testing.T) {
		t.Parallel()
		t.Run("OK", func(t *testing.T) {
			t.Parallel()
			// create in file
			expectedPlaintextBytes := []byte("skafiskafnjak")
			inFilePath := "test_in_file.txt"
			require.NoError(t, os.WriteFile(inFilePath, expectedPlaintextBytes, 0777))
			defer os.Remove(inFilePath)

			// create out file
			outFilePath := "test_out_file.txt"
			outFile, err := os.Create(outFilePath)
			require.NoError(t, err)
			defer func() {
				_ = outFile.Close()
				_ = os.Remove(outFilePath)
			}()

			// encrypt
			require.NoError(t, EncryptFile(validSymmetricKey, inFilePath, outFilePath))

			// assert that the ciphertext has been written to the out file.
			ciphertextBytes, err := os.ReadFile(outFilePath)
			require.NoError(t, err)
			require.NotNil(t, ciphertextBytes)
			require.NotEqual(t, expectedPlaintextBytes, ciphertextBytes)

			// decrypt the out file contents into a new file
			decryptedFilePath := "decrypted.txt"
			require.NoError(t, DecryptFile(validSymmetricKey, outFilePath, decryptedFilePath))
			defer os.Remove(decryptedFilePath)

			// assert that the plaintext file contents of the decrypted file are the same as the original in file.
			gotPlaintextBytes, err := os.ReadFile(decryptedFilePath)
			require.NoError(t, err)
			require.Equal(t, expectedPlaintextBytes, gotPlaintextBytes)
		})
	})
}
