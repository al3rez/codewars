package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindUniq(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		assert.Equal(t, FindUniq([]float32{1.0, 1.0, 1.0, 2.0, 1.0, 1.0}), float32(2))
		assert.Equal(t, FindUniq([]float32{0, 0, 0.55, 0, 0}), float32(0.55))
	})
}
