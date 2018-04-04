package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasUniqueChar(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		assert.Equal(t, HasUniqueChar("  nAa"), false)
		assert.Equal(t, HasUniqueChar("abcde"), true)
		assert.Equal(t, HasUniqueChar("++-"), false)
		assert.Equal(t, HasUniqueChar("AaBbC"), true)
	})
}
