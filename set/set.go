// Package set implements a set abstraction.
package set

var exists = struct{}{}

type StringSet struct {
	items map[string]struct{}
}

func NewStringSet() StringSet {
	set := StringSet{}
	set.items = make(map[string]struct{})
	return set
}

func (s *StringSet) Add(value string) {
	s.items[value] = exists
}

func (s *StringSet) Remove(value string) {
	delete(s.items, value)
}

func (s *StringSet) Values() []string {
	values := make([]string, 0, len(s.items))
	for key := range s.items {
		values = append(values, key)
	}

	return values
}

func (s *StringSet) Contains(value string) bool {
	_, contains := s.items[value]
	return contains
}

func (s *StringSet) Size() int {
	return len(s.items)
}
