package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveFirstLast(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		assert.Equal(t, "b", RemoveFirstLast("abc"))
		assert.Equal(t, "bc", RemoveFirstLast("abcd"))
		assert.Equal(t, "bcd", RemoveFirstLast("abcde"))
		assert.Equal(t, "bcde", RemoveFirstLast("abcdef"))
	})
}
