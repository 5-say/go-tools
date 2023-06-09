package t_test

import (
	"testing"

	tool "github.com/5-say/go-tools/tools/t"
	"github.com/stretchr/testify/assert"
)

func TestPointer(t *testing.T) {
	var input *string
	result := tool.Pointer(input, "default")

	assert.Equal(t, result, "default")
}
