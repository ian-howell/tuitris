package ring

import (
	"fmt"
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

func TestNext(t *testing.T) {
	tests := []struct {
		values   []string
		cursor   int
		expected string
	}{
		{[]string{"foo", "bar", "baz"}, 0, "bar"},
		{[]string{"foo", "bar", "baz"}, 1, "baz"},
		{[]string{"foo", "bar", "baz"}, 2, "foo"},
	}

	for i, tt := range tests {
		testName := fmt.Sprintf("test-%d", i)
		t.Run(testName, func(t *testing.T) {
			assert := assert.New(t)

			r := &ring[string]{tt.values, tt.cursor}

			assert.Equal(tt.values[tt.cursor], r.Get())

			r.Next()
			assert.Equal(tt.expected, r.Get())
		})
	}
}

func TestPrev(t *testing.T) {
	tests := []struct {
		values   []string
		cursor   int
		expected string
	}{
		{[]string{"foo", "bar", "baz"}, 0, "baz"},
		{[]string{"foo", "bar", "baz"}, 1, "foo"},
		{[]string{"foo", "bar", "baz"}, 2, "bar"},
	}

	for i, tt := range tests {
		testName := fmt.Sprintf("test-%d", i)
		t.Run(testName, func(t *testing.T) {
			assert := assert.New(t)

			r := &ring[string]{tt.values, tt.cursor}

			assert.Equal(tt.values[tt.cursor], r.Get())

			r.Prev()
			assert.Equal(tt.expected, r.Get())
		})
	}
}
