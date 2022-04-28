package collections

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPriorityQueue_Empty(t *testing.T) {
	pq := MakePriorityQueue(func(a, b int) int { return a - b })
	require.Equal(t, 0, pq.Len())
	require.Panics(t, func() { pq.Pop() })
}

func TestPriorityQueue_PushPopInt(t *testing.T) {
	pq := MakePriorityQueue(func(a, b int) int { return a - b })

	pq.Push(10)
	require.Equal(t, 1, pq.Len())

	pq.Push(30)
	require.Equal(t, 2, pq.Len())

	require.Equal(t, 10, pq.Pop())
	require.Equal(t, 1, pq.Len())

	pq.Push(20)
	require.Equal(t, 2, pq.Len())

	require.Equal(t, 20, pq.Pop())
	require.Equal(t, 1, pq.Len())

	require.Equal(t, 30, pq.Pop())
	require.Equal(t, 0, pq.Len())
}

type intWrapper struct {
	i int
}

func TestPriorityQueue_PushPopStruct(t *testing.T) {
	pq := MakePriorityQueue(func(a, b intWrapper) int { return b.i - a.i })

	pq.Push(intWrapper{i: 20})
	pq.Push(intWrapper{i: 30})
	require.Equal(t, intWrapper{i: 30}, pq.Pop())
	pq.Push(intWrapper{i: 10})
	require.Equal(t, intWrapper{i: 20}, pq.Pop())
	require.Equal(t, intWrapper{i: 10}, pq.Pop())
}

func getSeededPriorityQueue(r *rand.Rand, n int) *PriorityQueue[int] {
	seedData := make([]int, n)
	for i := range seedData {
		seedData[i] = r.Int()
	}
	return MakePriorityQueue(func(a, b int) int { return a - b }, seedData...)
}

func BenchmarkPriorityQueue_PushPop_1e3(b *testing.B) {
	// Use a constant seed to get an arbitrary, deterministic sequence of values
	r := rand.New(rand.NewSource(230427))
	pq := getSeededPriorityQueue(r, 1e3)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		pq.Push(r.Int())
		pq.Pop()
	}
}

func BenchmarkPriorityQueue_PushPop_1e6(b *testing.B) {
	// Use a constant seed to get an arbitrary, deterministic sequence of values
	r := rand.New(rand.NewSource(458326))
	pq := getSeededPriorityQueue(r, 1e6)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		pq.Push(r.Int())
		pq.Pop()
	}
}
