package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		assert.Equal(t, []string{"ab", "cd"}, Solution("abcd"))
		assert.Equal(t, []string{"ab", "c_"}, Solution("abc"))
	})
}
