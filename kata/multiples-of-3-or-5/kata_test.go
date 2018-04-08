package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiple3And5(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		assert.Equal(t, 23, Multiple3And5(10))
	})

	t.Run("handle zero", func(t *testing.T) {
		assert.Equal(t, 0, Multiple3And5(0))
	})

	t.Run("multiples of both 3 and 5", func(t *testing.T) {
		assert.Equal(t, 1, Multiple3And5(15))
	})
}
