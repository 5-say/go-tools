package t_test

import (
	"testing"

	tool "github.com/5-say/go-tools/tools/t"
	"github.com/stretchr/testify/assert"
)

func TestRPCError(t *testing.T) {
	err := tool.RPCError("privateMessage", "publicMessage")
	e := tool.RPCErrorParse(err)

	assert.Equal(t, "privateMessage", e.PrivateMessage)
	assert.Equal(t, "publicMessage", e.PublicMessage)
}
