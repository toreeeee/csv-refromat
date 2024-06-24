package array_test

import (
	"csv/utils/array"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	output := array.Map(input, func(v *int, i int) string {
		return fmt.Sprintf("%d", *v)
	})

	assert.Equal(t, len(input), len(output), "map output length should be the same as input; else something went missing")
	assert.NotEqual(t, input, output, "input shouldn't match output when transformed")
	assert.Equal(t, []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}, output, "should match int input as string")
}

func TestFilter(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	output := array.Filter(input, func(v *int) bool {
		return *v > 10
	})

	assert.NotEqual(t, len(input), len(output), "output shouldn't match input length after filtering")
	assert.Equal(t, []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, output)
}
