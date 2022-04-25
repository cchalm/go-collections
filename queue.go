package collections

type Queue[E any] struct {
	s []E
}

func MakeQueue[E any]() *Queue[E] {
	return &Queue[E]{}
}

func (q *Queue[E]) Push(e E) {
	q.s = append(q.s, e)
}

func (q *Queue[E]) Pop() E {
	popped := q.s[0]
	q.s = q.s[1:]
	return popped
}

func (q Queue[E]) Len() int {
	return len(q.s)
}
