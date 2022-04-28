package collections

import "container/heap"

// PriorityQueue is an ordered collection of elements. Elements can be pushed and popped from a PriorityQueue; when
// popping, elements come out in an order determined by the provided comparator. Elements that compare as "less" than
// others come out first
type PriorityQueue[E any] struct {
	heap *comparatorHeap[E]
}

// Comparator is a function that returns a value indicating whether the first argument is less than, greater than, or
// equal to the second argument.
// If the returned value is less than zero, the first argument is less than the second
// If the returned value is greater than zero, the first argument is greater than the second
// If the returned value is zero, the first argument is equal to the second
type Comparator[E any] func(a, b E) int

func MakePriorityQueue[E any](c Comparator[E], initialElements ...E) *PriorityQueue[E] {
	h := &comparatorHeap[E]{
		arr:     initialElements,
		compare: c,
	}
	heap.Init(h)
	return &PriorityQueue[E]{
		heap: h,
	}
}

func (pq *PriorityQueue[E]) Push(e E) {
	heap.Push(pq.heap, e)
}

func (pq *PriorityQueue[E]) Pop() E {
	e := heap.Pop(pq.heap)
	return e.(E)
}

func (pq PriorityQueue[E]) Len() int {
	return pq.heap.Len()
}

// comparatorHeap is a type that implements the interface required to behave as a heap when used with the built-it heap
// package. It orders elements of any type using a provided comparator
type comparatorHeap[E any] struct {
	arr     []E
	compare Comparator[E]
}

func (ch comparatorHeap[E]) Len() int           { return len(ch.arr) }
func (ch comparatorHeap[E]) Less(i, j int) bool { return ch.compare(ch.arr[i], ch.arr[j]) < 0 }
func (ch comparatorHeap[E]) Swap(i, j int)      { ch.arr[i], ch.arr[j] = ch.arr[j], ch.arr[i] }

func (ch *comparatorHeap[E]) Push(x any) {
	ch.arr = append(ch.arr, x.(E))
}

func (ch *comparatorHeap[E]) Pop() any {
	old := ch.arr
	n := len(old)
	x := old[n-1]
	ch.arr = old[0 : n-1]
	return x
}
