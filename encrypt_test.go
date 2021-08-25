package sym

import (
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestEncrypt(t *testing.T) {
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
			_, err := EncryptBytes(invalidSymmetricKeyBytes, nil)
			require.Error(t, err)
		})
		t.Run("should fail if plaintext bytes are empty", func(t *testing.T) {
			t.Parallel()
			_, err := EncryptBytes(validSymmetricKeyBytes, []byte{})
			require.Error(t, err)
		})
		t.Run("should fail if plaintext bytes are nil", func(t *testing.T) {
			t.Parallel()
			_, err := EncryptBytes(validSymmetricKeyBytes, nil)
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
			defer func() {
				_ = os.Remove(inFilePath)
			}()

			// create out file
			outFilePath := randomStringOfLen(11) + "test_out_file.txt"
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
		})
		t.Run("should fail if file does not exist", func(t *testing.T) {
			t.Parallel()
			require.Error(t, EncryptFile(validSymmetricKey, "doesntexist", "doesntexist"))
		})
	})
}

func randomStringOfLen(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
