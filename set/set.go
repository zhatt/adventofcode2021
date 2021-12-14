// Package set implements a set abstraction.
package set

var exists = struct{}{}

type stringSet struct {
	items map[string]struct{}
}

func NewStringSet() *stringSet {
	set := &stringSet{}
	set.items = make(map[string]struct{})
	return set
}

func (s *stringSet) Add(value string) {
	s.items[value] = exists
}

func (s *stringSet) Remove(value string) {
	delete(s.items, value)
}

func (s *stringSet) Contains(value string) bool {
	_, contains := s.items[value]
	return contains
}

func (s *stringSet) Size() int {
	return len(s.items)
}
