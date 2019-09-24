package types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func populateSet(count int, start int) *Set {
	set := Set{}
	for i := start; i < (start + count); i++ {
		set.Add(fmt.Sprintf("item%d", i))
	}
	return &set
}

func TestAdd(t *testing.T) {
	newSet := Set{}
	newSet.Add("item1")
	assert.Equal(t, newSet.Size(), 1)
	set := populateSet(3, 0)
	assert.Equal(t, set.Size(), 3)
	newSet.Add("item1")
	assert.Equal(t, newSet.Size(), 1)
}

func TestDelete(t *testing.T) {
	set := populateSet(3, 0)
	set.Delete("item2")
	assert.Equal(t, set.Size(), 2)
	set.Delete("item2")
	assert.Equal(t, set.Size(), 2)
}

func TestHas(t *testing.T) {
	set := populateSet(3, 0)
	has := set.Has("item2")
	assert.True(t, has)

	set.Delete("item2")
	has = set.Has("item2")
	assert.False(t, has)
}

func TestSize(t *testing.T) {
	set := populateSet(3, 0)
	size := set.Size()
	assert.Equal(t, size, 3)
}
