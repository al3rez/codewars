package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultipleOfIndex(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		assert.Equal(t, []int{-6, 32, 25}, multipleOfIndex([]int{22, -6, 32, 82, 9, 25}))
		assert.Equal(t, []int{-1, 10}, multipleOfIndex([]int{68, -1, 1, -7, 10, 10}))
		assert.Equal(t, []int{-11}, multipleOfIndex([]int{11, -11}))
		assert.Equal(t, []int{-85, 72, 0, 68}, multipleOfIndex([]int{-56, -85, 72, -26, -14, 76, -27, 72, 35, -21, -67, 87, 0, 21, 59, 27, -92, 68}))
		assert.Equal(t, []int{38, -44, -99}, multipleOfIndex([]int{28, 38, -44, -99, -13, -54, 77, -51}))
		assert.Equal(t, []int{-49, 8, -60, 35}, multipleOfIndex([]int{-1, -49, -1, 67, 8, -60, 39, 35}))
	})
}
