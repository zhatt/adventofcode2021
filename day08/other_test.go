package main

import (
	"strings"
	"testing"

	"github.com/gitchander/permutation"

	"github.com/stretchr/testify/assert"

	"zhatt/aoc2021/set"
)

func TestPermutation(t *testing.T) {
	a := []string{"a", "b", "c"}
	p := permutation.New(permutation.StringSlice(a))

	permutation_set := set.NewStringSet()
	for p.Next() {
		permutation_set.Add(strings.Join(a, ""))
	}

	assert.Equal(t, 6, permutation_set.Size())
	assert.True(t, permutation_set.Contains("abc"))
	assert.True(t, permutation_set.Contains("acb"))
	assert.True(t, permutation_set.Contains("bac"))
	assert.True(t, permutation_set.Contains("bca"))
	assert.True(t, permutation_set.Contains("cab"))
	assert.True(t, permutation_set.Contains("cba"))
}
