// Package set implements a set abstraction.
package set

type StringMultiSet struct {
	items map[string]int
}

func NewStringMultiSet() StringMultiSet {
	set := StringMultiSet{}
	set.items = make(map[string]int)
	return set
}

func (s *StringMultiSet) Add(value string) {
	s.items[value]++
}

func (s *StringMultiSet) Remove(value string) {
	if count, okay := s.items[value]; okay {
		if count > 1 {
			s.items[value]--
		} else {
			delete(s.items, value)
		}
	}
}

func (s *StringMultiSet) Values() []string {
	values := make([]string, 0, len(s.items))
	for key := range s.items {
		values = append(values, key)
	}

	return values
}

func (s *StringMultiSet) Contains(value string) bool {
	_, contains := s.items[value]
	return contains
}

func (s *StringMultiSet) Count(value string) int {
	return s.items[value]
}

func (s *StringMultiSet) Size() int {
	return len(s.items)
}
