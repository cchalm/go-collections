package collections

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSet_Empty(t *testing.T) {
	s := MakeSet[int]()
	require.Equal(t, 0, s.Len())
	require.False(t, s.Contains(10))
}

func TestSet_Insert(t *testing.T) {
	s := MakeSet[int]()
	inserted := s.Insert(10)
	require.True(t, inserted)
	require.Equal(t, 1, s.Len())
	require.True(t, s.Contains(10))
}

func TestSet_InsertDuplicate(t *testing.T) {
	s := MakeSet[int]()
	s.Insert(10)
	inserted := s.Insert(10)
	require.False(t, inserted)
	require.Equal(t, 1, s.Len())
	require.True(t, s.Contains(10))
}

func TestSet_InsertMultiple(t *testing.T) {
	s := MakeSet[int]()
	s.Insert(10)
	inserted := s.Insert(20)
	require.True(t, inserted)
	require.Equal(t, 2, s.Len())
	require.True(t, s.Contains(10))
	require.True(t, s.Contains(20))
}

func TestSet_InitialElements(t *testing.T) {
	s := MakeSet(10, 20, 30)
	require.Equal(t, 3, s.Len())
	require.True(t, s.Contains(10))
	require.True(t, s.Contains(20))
	require.True(t, s.Contains(30))
}

func TestSet_String(t *testing.T) {
	s := MakeSet(10, 20)
	str := s.String()
	require.True(t, "{10, 20}" == str || "{20, 10}" == str, str)
}

func TestSet_ToSlice(t *testing.T) {
	s := MakeSet(10, 20)
	require.ElementsMatch(t, []int{10, 20}, s.ToSlice())
}
