package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	set := NewStringSet()
	assert.Equal(t, 0, set.Size())
}

func TestAdd(t *testing.T) {
	set := NewStringSet()
	set.Add("a")
	set.Add("b")

	assert.Equal(t, 2, set.Size())
	assert.True(t, set.Contains("a"))
	assert.True(t, set.Contains("b"))
	assert.False(t, set.Contains("c"))
}

func TestSize(t *testing.T) {
	set := NewStringSet()
	assert.Equal(t, 0, set.Size())
	set.Add("a")
	assert.Equal(t, 1, set.Size())
	set.Add("b")
	assert.Equal(t, 2, set.Size())
}

func TestDelete(t *testing.T) {
	set := NewStringSet()
	set.Add("a")
	set.Add("b")
	set.Remove("a")

	assert.Equal(t, 1, set.Size())
	assert.False(t, set.Contains("a"))
	assert.True(t, set.Contains("b"))
}
