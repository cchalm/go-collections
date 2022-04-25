package collections

type Stack[E any] struct {
	s []E
}

func MakeStack[E any]() *Stack[E] {
	return &Stack[E]{}
}

func (s *Stack[E]) Push(e E) {
	s.s = append(s.s, e)
}

func (s *Stack[E]) Pop() E {
	i := len(s.s) - 1
	popped := s.s[i]
	s.s = s.s[:i]
	return popped
}

func (s Stack[E]) Len() int {
	return len(s.s)
}
