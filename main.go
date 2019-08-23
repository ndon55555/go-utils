package main

import (
	"fmt"

	"github.com/ndon55555/go-utils/slice"
)

func main() {
	s1 := slice.New(4, 2, 3, 1, 5, 9, 7, 8, 6, 0)

	fmt.Println(s1.Contains(3))
	fmt.Println(s1.Contains(0))

	fmt.Println(s1.MapTo(func(e slice.AnyType) slice.AnyType {
		return e.(int) + 1
	}))

	fmt.Println(s1.Filter(func(e slice.AnyType) bool {
		return e.(int)%2 == 0
	}))

	fmt.Println(s1.MapTo(func(e slice.AnyType) slice.AnyType {
		return e.(int) + 1
	}).Filter(func(e slice.AnyType) bool {
		return e.(int)%2 == 0
	}))

	fmt.Println(s1.Fold(s1, func(acc slice.AnyType, e slice.AnyType) slice.AnyType {
		return append(acc.(slice.GenericSlice), e)
	}))

	fmt.Println(s1.Fold(slice.GenericSlice{}, func(acc slice.AnyType, e slice.AnyType) slice.AnyType {
		for i := 1; i <= 2; i++ {
			acc = append(acc.(slice.GenericSlice), e)
		}

		return acc
	}))

	fmt.Println(s1.Fold(slice.GenericSlice{}, func(acc slice.AnyType, e slice.AnyType) slice.AnyType {
		return append(slice.New(e), acc.(slice.GenericSlice)...)
	}))

	fmt.Println(s1.Fold(slice.GenericSlice{}, func(acc slice.AnyType, e slice.AnyType) slice.AnyType {
		innerSlice := slice.GenericSlice{}
		for i := 1; i <= 2; i++ {
			innerSlice = append(innerSlice, e)
		}

		return append(acc.(slice.GenericSlice), innerSlice)
	}))

	fmt.Println(s1.QuickSort(func(x slice.AnyType, y slice.AnyType) int {
		return x.(int) - y.(int)
	}))
}
