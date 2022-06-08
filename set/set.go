package set

import "github.com/BooleanCat/go-functional/iter"

type Set[T comparable] interface {
	Add(elements ...T) Set[T]
	Remove(elements T) Set[T]
	Has(element T) bool
	Elements() []T
	IterElements() iter.Iterator[T]
	Len() int
	IsEmpty() bool
	Equals(other Set[T]) bool
	Union(other Set[T]) Set[T]
	Intersection(other Set[T]) Set[T]
	Difference(other Set[T]) Set[T]
	SymmetricDifference(other Set[T]) Set[T]
	String() string
}
