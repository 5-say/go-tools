package change_test

import (
	"testing"

	"github.com/5-say/go-tools/tools/change"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, change.ToString(uint8(66)), "66")
}
