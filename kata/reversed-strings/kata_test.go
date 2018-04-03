package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		assert.Equal(t, Solution("world"), "dlrow")
		assert.Equal(t, Solution("hello"), "olleh")
	})
}
