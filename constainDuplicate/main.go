// https://leetcode.com/problems/contains-duplicate/
package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// 可以在泛型函数中使用 constraint.Ordered，而无需创建新的接口约束。不过，为了便于学习，我们还是创建了自己的约束 AllowData
type AllowedData interface {
	constraints.Ordered
}

type Filter[T AllowedData] map[T]bool

func (r Filter[T]) add(datum T) {
	r[datum] = true
}

func (r Filter[T]) has(datum T) bool {
	_, ok := r[datum]
	return ok
}

func FindDuplicate[T AllowedData](data []T) bool {
	inArray := Filter[T]{}
	for _, datum := range data {
		if inArray.has(datum) {
			return true
		}
		inArray.add(datum)
	}
	return false
}

func main() {
	data := []int{1, 3, 4, 4, 5, 8, 7, 3, 2}     // sample array
	data32 := []int32{1, 3, 4, 4, 5, 8, 7, 3, 2} // sample array
	data64 := []int64{1, 3, 4, 4, 5, 8, 7, 3, 2} // sample array
	fmt.Printf("Duplicate found %t\n", FindDuplicate(data))
	fmt.Printf("Duplicate found %t\n", FindDuplicate(data32))
	fmt.Printf("Duplicate found %t\n", FindDuplicate(data64))
}

// type (
// 	FilterInt   map[int]bool
// 	FilterInt32 map[int32]bool
// 	FilterInt64 map[int64]bool
// )
//
// func main() {
// 	data := []int{1, 3, 4, 4, 5, 8, 7, 3, 2}     // sample array
// 	data32 := []int32{1, 3, 4, 4, 5, 8, 7, 3, 2} // sample array
// 	data64 := []int64{1, 3, 4, 4, 5, 8, 7, 3, 2} // sample array
// 	fmt.Printf("Duplicate found %t\n", FindDuplicateInt(data))
// 	fmt.Printf("Duplicate found %t\n", FindDuplicateInt32(data32))
// 	fmt.Printf("Duplicate found %t\n", FindDuplicateInt64(data64))
// }

// func FindDuplicateInt(data []int) bool {
// 	inArray := FilterInt{}
// 	for _, datum := range data {
// 		if inArray.has(datum) {
// 			return true
// 		}
// 		inArray.add(datum)
// 	}
// 	return false
// }
//
// func FindDuplicateInt32(data []int32) bool {
// 	inArray := FilterInt32{}
// 	for _, datum := range data {
// 		if inArray.has(datum) {
// 			return true
// 		}
// 		inArray.add(datum)
// 	}
// 	return false
// }
//
// func FindDuplicateInt64(data []int64) bool {
// 	inArray := FilterInt64{}
// 	for _, datum := range data {
// 		if inArray.has(datum) {
// 			return true
// 		}
// 		inArray.add(datum)
// 	}
// 	return false
// }
//
// func (r FilterInt) add(datum int) {
// 	r[datum] = true
// }
//
// func (r FilterInt32) add(datum int32) {
// 	r[datum] = true
// }
//
// func (r FilterInt64) add(datum int64) {
// 	r[datum] = true
// }
//
// func (r FilterInt) has(datum int) bool {
// 	_, ok := r[datum]
// 	return ok
// }
//
// func (r FilterInt32) has(datum int32) bool {
// 	_, ok := r[datum]
// 	return ok
// }
//
// func (r FilterInt64) has(datum int64) bool {
// 	_, ok := r[datum]
// 	return ok
// }
