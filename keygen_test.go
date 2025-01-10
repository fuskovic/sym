package sym

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKeyGen(t *testing.T) {
	t.Parallel()

	t.Run("should succeed", func(t *testing.T) {
		t.Parallel()
		for _, test := range []struct {
			name  string
			input int
		}{
			{
				name:  "when size is 16",
				input: 16,
			},
			{
				name:  "when size is 24",
				input: 24,
			},
			{
				name:  "when size is 32",
				input: 32,
			},
		} {
			t.Run(test.name, func(t *testing.T) {
				k, err := KeyGen(test.input)
				require.NoError(t, err)
				require.NotNil(t, k)
			})
		}
	})
	t.Run("should fail when size is not valid", func(t *testing.T) {
		t.Parallel()
		k, err := KeyGen(10)
		require.Error(t, err)
		require.Nil(t, k)
	})
}
