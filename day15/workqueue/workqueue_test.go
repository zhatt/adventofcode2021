package workqueue

import (
	"testing"
	"zhatt/aoc2021/coord"

	"github.com/stretchr/testify/assert"
)

var items = map[coord.Coord]int{
	{Xval: 3, Yval: 0}: 3,
	{Xval: 2, Yval: 0}: 2,
	{Xval: 4, Yval: 0}: 4,
}

func create() WorkQueue {
	wq := NewWorkQueue()
	for location, distance := range items {
		work := NewItem(location, distance)
		wq.Push(work)
	}
	return wq
}

func TestCreatePushPop(t *testing.T) {
	wq := create()
	assert.Equal(t, 3, wq.Len())
	assert.Equal(t, coord.Coord{Xval: 2, Yval: 0}, wq.Pop().Location)
	assert.Equal(t, coord.Coord{Xval: 3, Yval: 0}, wq.Pop().Location)
	assert.Equal(t, coord.Coord{Xval: 4, Yval: 0}, wq.Pop().Location)
	assert.Equal(t, 0, wq.Len())
}

func TestUpdate(t *testing.T) {
	wq := create()

	work := NewItem(coord.Coord{Xval: 1, Yval: 0}, 1)
	assert.Equal(t, -1, work.index)
	wq.Push(work)
	assert.Equal(t, 0, work.index)
	assert.Equal(t, coord.Coord{Xval: 1, Yval: 0}, wq.Pop().Location)
	assert.Equal(t, -1, work.index)
	wq.Push(work)
	assert.Equal(t, 0, work.index)
	work.Distance = 10
	wq.Update(work)
	assert.NotEqual(t, 0, work.index)

	assert.Equal(t, 4, wq.Len())
	assert.Equal(t, coord.Coord{Xval: 2, Yval: 0}, wq.Pop().Location)
	assert.Equal(t, coord.Coord{Xval: 3, Yval: 0}, wq.Pop().Location)
	assert.Equal(t, coord.Coord{Xval: 4, Yval: 0}, wq.Pop().Location)
	assert.Equal(t, coord.Coord{Xval: 1, Yval: 0}, wq.Pop().Location)
	assert.Equal(t, 0, wq.Len())
}
