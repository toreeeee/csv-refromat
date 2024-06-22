package extra_math_test

import (
	"csv/utils/extra_math"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMin(t *testing.T) {
	assert.Equal(t, 1, extra_math.Min(1, 2))
	assert.Equal(t, 1, extra_math.Min(2, 1))

	assert.Equal(t, -2, extra_math.Min(-1, -2))
	assert.Equal(t, -55, extra_math.Min(55, -55))

	assert.Equal(t, 1, extra_math.Min(1, 1))
}
