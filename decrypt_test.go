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
			ciphertext, err := EncryptString(validSymmetricKey, expected)
			require.NoError(t, err)
			require.NotEqual(t, expected, ciphertext)
			got, err := DecryptString(validSymmetricKey, ciphertext)
			require.NoError(t, err)
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
			ciphertext, err := EncryptBytes(validSymmetricKey, expected)
			require.NoError(t, err)
			require.NotEqual(t, expected, ciphertext)
			got, err := DecryptBytes(validSymmetricKey, ciphertext)
			require.NoError(t, err)
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
			// create in file
			expectedPlaintextBytes := []byte("skafiskafnjak")
			inFilePath := randomStringOfLen(10) + "test_in_file.txt"
			require.NoError(t, os.WriteFile(inFilePath, expectedPlaintextBytes, 0777))
			defer os.Remove(inFilePath)

			// create out file
			outFilePath := randomStringOfLen(10) + "test_out_file.txt"
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
			decryptedFilePath := randomStringOfLen(10) + "decrypted.txt"
			require.NoError(t, DecryptFile(validSymmetricKey, outFilePath, decryptedFilePath))
			defer os.Remove(decryptedFilePath)

			// assert that the plaintext file contents of the decrypted file are the same as the original in file.
			gotPlaintextBytes, err := os.ReadFile(decryptedFilePath)
			require.NoError(t, err)
			require.Equal(t, expectedPlaintextBytes, gotPlaintextBytes)
		})
		t.Run("should fail if file does not exist", func(t *testing.T) {
			t.Parallel()
			require.Error(t, DecryptFile(validSymmetricKey, "doesntexist", "doesntexist"))
		})
	})
}
