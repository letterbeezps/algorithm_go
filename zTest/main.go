package main

import (
	"fmt"
	"testing"
)

// func main() {

// 	arrs := []int{3, 5, 2, 5, 7, 1, 3, 5, 25, 2, 15}
// 	fmt.Printf("before arr: %v\n", arrs)

// 	quick_sort(arrs, 0, len(arrs)-1)
// 	fmt.Printf("after arr: %v\n", arrs)
// }

func quick_sort(arr []int, l int, r int) {
	if l >= r {
		return
	}

	var left, right int = l - 1, r + 1
	mid := (l + r) / 2

	x := arr[mid]
	for left < right {
		left++
		for arr[left] < x {
			left++
		}

		right--
		for arr[right] > x {
			right--
		}

		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}
	// fmt.Printf("temp arr: %v\n", arr)
	// fmt.Printf("left is %d \n", left)
	// fmt.Printf("right is %d \n", right)
	quick_sort(arr, l, right)
	quick_sort(arr, right+1, r)
}

func BenchmarkQuickSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		arrs := []int{3, 5, 2, 5, 7, 1, 3, 5, 25, 2, 15}
		quick_sort(arrs, 0, len(arrs)-1)
		fmt.Println(arrs)
	}
}
