package play

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCellView(t *testing.T) {
	assert := assert.New(t)
	cell := cellView('W')
	input := "" +
		"▛▜\n" +
		"▙▟"
	assert.Equal(input, cell)
}

func TestRowView(t *testing.T) {
	assert := assert.New(t)

	input := "#W        W#"
	expected := "" +
		"▛▜                ▛▜\n" +
		"▙▟                ▙▟"
	assert.Equal(rowView([]rune(input)), expected)
}

func TestPlayfieldView(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)
	m, err := New()
	require.NoError(err)

	expected := "" +
		"▛▜                ▛▜\n" +
		"▙▟                ▙▟\n" +
		"▛▜                ▛▜\n" +
		"▙▟                ▙▟\n" +
		"▛▜                ▛▜\n" +
		"▙▟                ▙▟\n" +
		"▛▜                ▛▜\n" +
		"▙▟                ▙▟\n" +
		"▛▜                ▛▜\n" +
		"▙▟                ▙▟\n" +
		"▛▜                ▛▜\n" +
		"▙▟                ▙▟\n" +
		"▛▜                ▛▜\n" +
		"▙▟                ▙▟\n" +
		"▛▜                ▛▜\n" +
		"▙▟                ▙▟\n" +
		"▛▜                ▛▜\n" +
		"▙▟                ▙▟\n" +
		"▛▜                ▛▜\n" +
		"▙▟                ▙▟\n" +
		"▛▜                ▛▜\n" +
		"▙▟                ▙▟\n" +
		"▛▜                ▛▜\n" +
		"▙▟                ▙▟\n" +
		"▛▜                ▛▜\n" +
		"▙▟                ▙▟\n" +
		"▛▜                ▛▜\n" +
		"▙▟                ▙▟\n" +
		"▛▜                ▛▜\n" +
		"▙▟                ▙▟\n" +
		"▛▜                ▛▜\n" +
		"▙▟                ▙▟\n" +
		"▛▜                ▛▜\n" +
		"▙▟                ▙▟\n" +
		"▛▜                ▛▜\n" +
		"▙▟                ▙▟\n" +
		"▛▜                ▛▜\n" +
		"▙▟                ▙▟\n" +
		"▛▜▛▜▛▜▛▜▛▜▛▜▛▜▛▜▛▜▛▜\n" +
		"▙▟▙▟▙▟▙▟▙▟▙▟▙▟▙▟▙▟▙▟"

	assert.Equal(m.playView(), expected)
}
