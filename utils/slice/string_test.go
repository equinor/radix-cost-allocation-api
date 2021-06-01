package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_StringSliceContains(t *testing.T) {
	actual := StringSliceContains([]string{"a", "b"}, "a")
	assert.True(t, actual)
	actual = StringSliceContains([]string{"a", "b"}, "b")
	assert.True(t, actual)
	actual = StringSliceContains([]string{"a", "b"}, "c")
	assert.False(t, actual)
}
