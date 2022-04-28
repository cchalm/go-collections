package collections

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStack_Empty(t *testing.T) {
	s := MakeStack[int]()
	require.Equal(t, 0, s.Len())
	require.Panics(t, func() { s.Pop() })
}

func TestStack_PushPop(t *testing.T) {
	s := MakeStack[int]()

	s.Push(10)
	require.Equal(t, 1, s.Len())

	s.Push(20)
	require.Equal(t, 2, s.Len())

	require.Equal(t, 20, s.Pop())
	require.Equal(t, 1, s.Len())

	s.Push(30)
	require.Equal(t, 2, s.Len())

	require.Equal(t, 30, s.Pop())
	require.Equal(t, 1, s.Len())

	require.Equal(t, 10, s.Pop())
	require.Equal(t, 0, s.Len())
}

func BenchmarkStack_PushPop(b *testing.B) {
	s := MakeStack[int]()

	for i := 0; i < b.N; i++ {
		s.Push(10)
		s.Pop()
	}
}
