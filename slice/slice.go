package slice

type AnyType interface{}
type GenericSlice []AnyType

type combiner func(acc AnyType, e AnyType) AnyType
type transformer func(e AnyType) AnyType
type predicate func(e AnyType) bool
type sliceCombiner func(acc GenericSlice, e AnyType) GenericSlice

func (s GenericSlice) Contains(e AnyType) bool {
	for _, i := range s {
		if i == e {
			return true
		}
	}

	return false
}

func (s GenericSlice) MapTo(f transformer) GenericSlice {
	return s.foldToSlice(GenericSlice{}, func(acc GenericSlice, e AnyType) GenericSlice {
		return append(acc, f(e))
	})
}

func (s GenericSlice) Filter(p predicate) GenericSlice {
	return s.foldToSlice(GenericSlice{}, func(acc GenericSlice, e AnyType) GenericSlice {
		if p(e) {
			return append(acc, e)
		} else {
			return acc
		}
	})
}

func (s GenericSlice) Fold(acc AnyType, f combiner) AnyType {
	for _, e := range s {
		acc = f(acc, e)
	}

	return acc
}

func New(e ...AnyType) GenericSlice {
	return append(GenericSlice{}, e...)
}

func (s GenericSlice) foldToSlice(acc GenericSlice, f sliceCombiner) GenericSlice {
	return s.Fold(acc, func(innerAcc AnyType, e AnyType) AnyType {
		return f(innerAcc.(GenericSlice), e)
	}).(GenericSlice)
}

func (s GenericSlice) QuickSort(comparer func(AnyType, AnyType) int) GenericSlice {
	if len(s) <= 1 {
		return s
	}

	pivot := s[0]
	lessEqual, greater := New(), New()
	for _, e := range s {
		d := comparer(e, pivot)
		if d <= 0 {
			lessEqual = append(lessEqual, e)
		} else {
			greater = append(greater, e)
		}
	}

	return append(
		lessEqual.QuickSort(comparer),
		greater.QuickSort(comparer)...,
	)
}
