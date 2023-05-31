package random_test

import (
	"testing"

	"github.com/5-say/go-tools/tools/random"
	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {
	r := random.Simple(random.SimpleRandomConfig{
		MapShifting:       3211,
		ConfusionShifting: 9475,
		ValueShifting:     9267,
	})

	var (
		number = int64(849384)
		en     = r.Encode(number, 8)
		de     = r.Decode(en)
	)

	assert.Equal(t, number, de)
}
