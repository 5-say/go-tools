package password_test

import (
	"testing"

	"github.com/5-say/go-tools/tools/password"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	inputPassword := "demo"
	hashedPassword, err := password.Generate(inputPassword)
	if assert.NoError(t, err) {
		err = password.Compare(inputPassword, hashedPassword)
		assert.NoError(t, err)
	}
}
