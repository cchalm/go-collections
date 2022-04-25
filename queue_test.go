package collections

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueue_Empty(t *testing.T) {
	q := MakeQueue[int]()
	require.Equal(t, 0, q.Len())
}

func TestQueue_PushPop(t *testing.T) {
	q := MakeQueue[int]()

	q.Push(10)
	require.Equal(t, 1, q.Len())

	q.Push(20)
	require.Equal(t, 2, q.Len())

	require.Equal(t, 10, q.Pop())
	require.Equal(t, 1, q.Len())

	q.Push(30)
	require.Equal(t, 2, q.Len())

	require.Equal(t, 20, q.Pop())
	require.Equal(t, 1, q.Len())

	require.Equal(t, 30, q.Pop())
	require.Equal(t, 0, q.Len())
}

func BenchmarkQueue_PushPop(b *testing.B) {
	q := MakeQueue[int]()

	for i := 0; i < b.N; i++ {
		q.Push(10)
		q.Pop()
	}
}
