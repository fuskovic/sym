package sym

import (
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func setupTestFiles(t *testing.T, expected []byte) (string, string, func()) {
	t.Helper()

	// create in file
	inFilePath := randomStringOfLen(t, 10) + "test_in_file.txt"
	require.NoError(t, os.WriteFile(inFilePath, expected, 0777))

	// create out file
	outFilePath := randomStringOfLen(t, 11) + "test_out_file.txt"
	outFile, err := os.Create(outFilePath)
	require.NoError(t, err)

	return inFilePath, outFilePath, func() {
		require.NoError(t, outFile.Close())
		require.NoError(t, os.Remove(inFilePath))
		require.NoError(t, os.Remove(outFilePath))
	}
}

func randomStringOfLen(t *testing.T, n int) string {
	t.Helper()

	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.NewSource(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func randomBytesOfLen(t *testing.T, n int) []byte {
	t.Helper()
	return []byte(randomStringOfLen(t, n))
}
