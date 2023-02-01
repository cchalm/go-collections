package collections

import (
	"math/rand"
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

func TestSet_Remove(t *testing.T) {
	s := MakeSet[int]()
	s.Insert(10)
	require.Equal(t, 1, s.Len())
	found := s.Remove(20)
	require.False(t, found)
	require.Equal(t, 1, s.Len())
	found = s.Remove(10)
	require.True(t, found)
	require.Equal(t, 0, s.Len())
	require.False(t, s.Contains(10))
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

func TestSet_CompareStructs(t *testing.T) {
	type a struct {
		x int
		y string
	}

	in1 := a{1, "foo"}
	in2 := a{2, "bar"}
	out1 := a{3, "baz"}

	s := MakeSet(in1, in2)
	require.True(t, s.Contains(in1))
	require.True(t, s.Contains(in2))
	require.False(t, s.Contains(out1))
}

func TestSet_CompareInterfaces(t *testing.T) {
	type a struct {
		x int
		y string
	}

	var in1 any = a{1, "foo"}
	var in2 any = "bar"
	var out1 any = a{3, "baz"}
	var out2 any = "foo"
	var out3 any = 10

	s := MakeSet(in1, in2)
	require.True(t, s.Contains(in1))
	require.True(t, s.Contains(in2))
	require.False(t, s.Contains(out1))
	require.False(t, s.Contains(out2))
	require.False(t, s.Contains(out3))
}

func getSeededSet(r *rand.Rand, n int) *Set[int] {
	seedData := make([]int, n)
	for i := range seedData {
		seedData[i] = r.Int()
	}
	return MakeSet(seedData...)
}

func BenchmarkSet_InsertRemove_1e3(b *testing.B) {
	// Use a constant seed to get an arbitrary, deterministic sequence of values
	r := rand.New(rand.NewSource(230427))
	pq := getSeededSet(r, 1e3)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		n := r.Int()
		pq.Insert(n)
		pq.Remove(n)
	}
}

func BenchmarkSet_InsertRemove_1e6(b *testing.B) {
	// Use a constant seed to get an arbitrary, deterministic sequence of values
	r := rand.New(rand.NewSource(458326))
	pq := getSeededSet(r, 1e6)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		n := r.Int()
		pq.Insert(n)
		pq.Remove(n)
	}
}
