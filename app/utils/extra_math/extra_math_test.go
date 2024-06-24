package extra_math_test

import (
	"csv/utils/extra_math"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMin(t *testing.T) {
	assert.Equal(t, 1, extra_math.Min(1, 2), "min value of 1 and 2 should be 1")
	assert.Equal(t, 1, extra_math.Min(2, 1), "min value of 2 and 1 should be 1")

	assert.Equal(t, -2, extra_math.Min(-1, -2), "min value of -1 and -2 should be -2")
	assert.Equal(t, -55, extra_math.Min(55, -55), "min value of 55 and -55 should be -55")

	assert.Equal(t, 1, extra_math.Min(1, 1), "when both values match output should match both")
}
