package kata

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDuplicateCount(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		assert.Equal(t, duplicateCount("abcde"), 0)
		assert.Equal(t, duplicateCount("abcdea"), 1)
		assert.Equal(t, duplicateCount("abcdeaB11"), 3)
		assert.Equal(t, duplicateCount("indivisibility"), 1)
	})
}
