package ring

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	tests := map[string]struct {
		values        []string
		expectedError string
	}{
		"empty":    {expectedError: "cannot create empty ring buffer"},
		"nonEmpty": {values: []string{"foo", "bar", "baz"}},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)
			require := require.New(t)

			r, err := New(tt.values...)
			if tt.expectedError != "" {
				require.EqualError(err, tt.expectedError)
			} else {
				require.NoError(err)
				assert.Equal(tt.values, r.Values())
			}

		})
	}
}
