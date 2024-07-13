package play

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCellView(t *testing.T) {
	assert := assert.New(t)
	cell := cellView('w')
	input := "" +
		"▛▜\n" +
		"▙▟"
	assert.Equal(input, cell)
}
