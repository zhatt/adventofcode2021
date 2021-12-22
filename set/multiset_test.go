package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultisetCreate(t *testing.T) {
	set := NewStringMultiSet()
	assert.Equal(t, 0, set.Size())
}

func TestMultiSetAdd(t *testing.T) {
	set := NewStringMultiSet()
	set.Add("a")
	set.Add("a")
	set.Add("b")

	assert.Equal(t, 2, set.Size())
	assert.True(t, set.Contains("a"))
	assert.True(t, set.Contains("b"))
	assert.False(t, set.Contains("c"))
	assert.Equal(t, 2, set.Count("a"))
	assert.Equal(t, 1, set.Count("b"))
}

func TestMultiSetSize(t *testing.T) {
	set := NewStringMultiSet()
	assert.Equal(t, 0, set.Size())
	set.Add("a")
	assert.Equal(t, 1, set.Size())
	set.Add("b")
	assert.Equal(t, 2, set.Size())
}

func TestMultiSetDelete(t *testing.T) {
	set := NewStringMultiSet()
	set.Add("a")
	set.Add("b")
	set.Add("b")
	set.Remove("a")

	assert.Equal(t, 1, set.Size())
	assert.False(t, set.Contains("a"))
	assert.True(t, set.Contains("b"))

	set.Remove("b")
	assert.Equal(t, 1, set.Size())
	assert.False(t, set.Contains("a"))
	assert.True(t, set.Contains("b"))

	set.Remove("b")
	assert.Equal(t, 0, set.Size())
	assert.False(t, set.Contains("a"))
	assert.False(t, set.Contains("b"))
}

func TestMultiSetValues(t *testing.T) {
	set := NewStringMultiSet()
	set.Add("a")
	set.Add("b")

	values := set.Values()

	assert.Equal(t, set.Size(), len(values))
	assert.Contains(t, values, "a")
	assert.Contains(t, values, "b")
}
