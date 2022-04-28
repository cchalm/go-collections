package collections

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPriorityQueue_Empty(t *testing.T) {
	pq := MakePriorityQueue(func(a, b int) int { return a - b })
	require.Equal(t, 0, pq.Len())
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