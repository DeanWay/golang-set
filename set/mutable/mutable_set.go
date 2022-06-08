package mutableset

import (
	"fmt"
	"reflect"

	"github.com/BooleanCat/go-functional/iter"
	"github.com/BooleanCat/go-functional/option"
	"github.com/DeanWay/golang-set/set"
)

type void struct{}

type MutableSet[T comparable] struct {
	members map[T]void
}

var _ set.Set[void] = (*MutableSet[void])(nil)

func Empty[T comparable]() *MutableSet[T] {
	return &MutableSet[T]{members: make(map[T]void)}
}

func FromSlice[T comparable](items []T) *MutableSet[T] {
	s := Empty[T]()
	s.Add(items...)
	return s
}

func FromIterator[T comparable](iterator iter.Iterator[T]) *MutableSet[T] {
	s := Empty[T]()
	for item := iterator.Next(); item.IsSome(); item = iterator.Next() {
		s.Add(item.Unwrap())
	}
	return s
}

func (s *MutableSet[T]) Add(elements ...T) set.Set[T] {
	var voidValue void
	m := s.members
	for _, element := range elements {
		m[element] = voidValue
	}
	return s
}

func (s *MutableSet[T]) Remove(element T) set.Set[T] {
	delete(s.members, element)
	return s
}

func (s *MutableSet[T]) Has(item T) bool {
	_, exists := s.members[item]
	return exists
}

type SetIterator[T comparable] struct {
	mapIter *reflect.MapIter
}

func (setIter *SetIterator[T]) Next() option.Option[T] {
	if setIter.mapIter.Next() {
		return option.Some(setIter.mapIter.Key().Interface().(T))
	}
	return option.None[T]()
}

func (s *MutableSet[T]) IterElements() iter.Iterator[T] {
	setReflection := reflect.ValueOf(s.members)
	return &SetIterator[T]{mapIter: setReflection.MapRange()}
}

func (s *MutableSet[T]) Elements() []T {
	result := make([]T, s.Len())
	i := 0
	for item := range s.members {
		result[i] = item
		i++
	}
	return result
}

func (s *MutableSet[T]) Len() int {
	return len(s.members)
}

func (s *MutableSet[T]) IsEmpty() bool {
	return s.Len() == 0
}

func (s *MutableSet[T]) Equals(other set.Set[T]) bool {
	if s.Len() != other.Len() {
		return false
	}
	for elem := range s.members {
		if !other.Has(elem) {
			return false
		}
	}
	return true
}

func (s1 *MutableSet[T]) Union(s2 set.Set[T]) set.Set[T] {
	result := Empty[T]()
	for _, element := range s1.Elements() {
		result.Add(element)
	}
	for _, element := range s2.Elements() {
		result.Add(element)
	}
	return result
}

func (s1 *MutableSet[T]) Intersection(s2 set.Set[T]) set.Set[T] {
	result := Empty[T]()
	for item := range s1.members {
		if s2.Has(item) {
			result.Add(item)
		}
	}
	return result
}

func (s1 *MutableSet[T]) Difference(s2 set.Set[T]) set.Set[T] {
	result := Empty[T]()
	for element := range s1.members {
		if !s2.Has(element) {
			result.Add(element)
		}
	}
	return result
}

func (s1 *MutableSet[T]) SymmetricDifference(s2 set.Set[T]) set.Set[T] {
	result := Empty[T]()
	for item := range s1.members {
		if !s2.Has(item) {
			result.Add(item)
		}
	}
	for _, element := range s2.Elements() {
		if !s1.Has(element) {
			result.Add(element)
		}
	}
	return result
}

func (s *MutableSet[T]) String() string {
	return fmt.Sprint(s.Elements())
}
