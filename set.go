package collections

import (
	"fmt"
	"strings"
)

// Set represents a unique set of elements. It has constant-time insert and contains operations
type Set[E comparable] struct {
	m map[E]struct{}
}

// MakeSet returns a Set containing the given initial elements
func MakeSet[E comparable](initialElements ...E) *Set[E] {
	s := &Set[E]{
		m: map[E]struct{}{},
	}
	for _, element := range initialElements {
		s.Insert(element)
	}
	return s
}

// Insert inserts the given element into the set. It returns true if the element was not previously in the set, false
// otherwise
func (s *Set[E]) Insert(element E) bool {
	_, ok := s.m[element]
	s.m[element] = struct{}{}
	return !ok
}

// Contains returns true if the given element is in the set, false otherwise
func (s Set[E]) Contains(element E) bool {
	_, ok := s.m[element]
	return ok
}

func (s Set[E]) Len() int {
	return len(s.m)
}

func (s Set[E]) String() string {
	var b strings.Builder
	b.WriteString("{")
	for e := range s.m {
		if b.Len() != 1 {
			b.WriteString(", ")
		}
		b.WriteString(fmt.Sprintf("%v", e))
	}
	b.WriteString("}")
	return b.String()
}

// ToSlice returns a slice containing shallow copies of each element in the set. Element order is not defined
func (s Set[E]) ToSlice() []E {
	arr := make([]E, 0, len(s.m))
	for e := range s.m {
		arr = append(arr, e)
	}
	return arr
}
