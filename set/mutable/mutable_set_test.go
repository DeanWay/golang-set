package mutableset

import (
	"testing"

	"github.com/DeanWay/golang-set/set"
	"github.com/stretchr/testify/require"
)

func TestEmpty(t *testing.T) {
	var s set.Set[int] = Empty[int]()
	require.Equal(t, s.Len(), 0)
	require.Equal(t, s.Elements(), []int{})
	require.True(t, s.IsEmpty())
}

func TestFromSlice(t *testing.T) {
	list := []string{"a", "b", "c", "c"}
	var s set.Set[string] = FromSlice(list)
	require.Equal(t, s.Len(), 3)
	require.ElementsMatch(t, s.Elements(), []string{"a", "b", "c"})
}

func TestAdd_SingleElement(t *testing.T) {
	var s set.Set[int] = Empty[int]()
	s = s.Add(1)
	require.ElementsMatch(t, s.Elements(), []int{1})
	s = s.Add(1)
	require.ElementsMatch(t, s.Elements(), []int{1})

	s = s.Add(2)
	require.ElementsMatch(t, s.Elements(), []int{1, 2})
	s = s.Add(2)
	require.ElementsMatch(t, s.Elements(), []int{1, 2})
}

func TestAdd_ManyElements(t *testing.T) {
	var s set.Set[int] = Empty[int]()
	s = s.Add(1, 1, 2, 2)
	require.ElementsMatch(t, s.Elements(), []int{1, 2})
}

func TestRemove(t *testing.T) {
	var s set.Set[int] = FromSlice([]int{1, 2, 3})
	s = s.Remove(1)
	require.ElementsMatch(t, s.Elements(), []int{2, 3})
	s = s.Remove(2)
	require.ElementsMatch(t, s.Elements(), []int{3})
	s = s.Remove(3)
	require.ElementsMatch(t, s.Elements(), []int{})
}

func TestEquals(t *testing.T) {
	require.True(t, Empty[void]().Equals(Empty[void]()))
	require.True(t, FromSlice([]int{1, 2}).Equals(FromSlice([]int{2, 1})))
	require.False(t, FromSlice([]int{1}).Equals(FromSlice([]int{1, 2})))
	require.False(t, FromSlice([]int{2, 1}).Equals(FromSlice([]int{1})))
	require.False(t, FromSlice([]string{"a", "b"}).Equals(FromSlice([]string{"A", "B"})))
}

func TestUnion(t *testing.T) {
	require.True(t, Empty[void]().Union(Empty[void]()).Equals(Empty[void]()))
	require.True(
		t,
		FromSlice([]int{1}).Union(
			FromSlice([]int{1}),
		).Equals(
			FromSlice([]int{1}),
		),
	)
	require.True(
		t,
		FromSlice([]int{1}).Union(
			FromSlice([]int{2}),
		).Equals(
			FromSlice([]int{1, 2}),
		),
	)
	require.True(
		t,
		FromSlice([]int{1, 2}).Union(
			FromSlice([]int{2, 3}),
		).Equals(
			FromSlice([]int{1, 2, 3}),
		),
	)
}

func TestIntersection(t *testing.T) {
	require.True(t, Empty[void]().Intersection(Empty[void]()).Equals(Empty[void]()))
	require.True(
		t,
		FromSlice([]int{1}).Intersection(
			FromSlice([]int{1}),
		).Equals(
			FromSlice([]int{1}),
		),
	)
	require.True(
		t,
		FromSlice([]int{1}).Intersection(
			Empty[int](),
		).Equals(
			Empty[int](),
		),
	)
	require.True(
		t,
		FromSlice([]int{1, 2}).Intersection(
			FromSlice([]int{2, 3}),
		).Equals(
			FromSlice([]int{2}),
		),
	)
}

func TestDifference(t *testing.T) {
	require.True(
		t,
		Empty[void]().Difference(
			Empty[void](),
		).Equals(
			Empty[void](),
		),
	)
	require.True(
		t,
		FromSlice([]int{1}).Difference(
			FromSlice([]int{1}),
		).Equals(
			Empty[int](),
		),
	)
	require.True(
		t,
		FromSlice([]int{1}).Difference(
			Empty[int](),
		).Equals(
			FromSlice([]int{1}),
		),
	)
	require.True(
		t,
		Empty[int]().Difference(
			FromSlice([]int{1}),
		).Equals(
			Empty[int](),
		),
	)
	require.True(
		t,
		FromSlice([]int{1, 2}).Difference(
			FromSlice([]int{2, 3}),
		).Equals(
			FromSlice([]int{1}),
		),
	)
	require.True(
		t,
		FromSlice([]int{2, 3}).Difference(
			FromSlice([]int{1, 2}),
		).Equals(
			FromSlice([]int{3}),
		),
	)

}

func TestSymmetricDifference(t *testing.T) {
	require.True(
		t,
		Empty[void]().SymmetricDifference(
			Empty[void](),
		).Equals(
			Empty[void](),
		),
	)
	require.True(
		t,
		FromSlice([]int{1}).SymmetricDifference(
			FromSlice([]int{1}),
		).Equals(
			Empty[int](),
		),
	)
	require.True(
		t,
		FromSlice([]int{1}).SymmetricDifference(
			Empty[int](),
		).Equals(
			FromSlice([]int{1}),
		),
	)
	require.True(
		t,
		Empty[int]().SymmetricDifference(
			FromSlice([]int{1}),
		).Equals(
			FromSlice([]int{1}),
		),
	)
	require.True(
		t,
		FromSlice([]int{1, 2}).SymmetricDifference(
			FromSlice([]int{2, 3}),
		).Equals(
			FromSlice([]int{1, 3}),
		),
	)
	require.True(
		t,
		FromSlice([]int{2, 3}).SymmetricDifference(
			FromSlice([]int{1, 2}),
		).Equals(
			FromSlice([]int{1, 3}),
		),
	)
}

func TestIterElements(t *testing.T) {
	list := []string{"a", "b", "c", "c"}
	var s set.Set[string] = FromSlice(list)
	iter := s.IterElements()
	result := []string{}
	for elem := iter.Next(); !elem.IsNone(); elem = iter.Next() {
		result = append(result, elem.Unwrap())
	}
	require.ElementsMatch(t, result, []string{"a", "b", "c"})
}

func TestFromIter(t *testing.T) {
	var s1 set.Set[string] = FromSlice([]string{"a", "b", "c", "c"})
	s2 := FromIterator(s1.IterElements())
	require.True(t, s1.Equals(s2))
}

func TestString(t *testing.T) {
	var s set.Set[string] = FromSlice([]string{"a"})
	require.Equal(t, s.String(), "[a]")
}
