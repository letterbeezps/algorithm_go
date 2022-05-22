package main

import (
	"testing"
)

// func main() {

// 	arrs := []int{3, 5, 2, 5, 7, 1, 3, 5, 25, 2, 15}
// 	fmt.Printf("before arr: %v\n", arrs)

// 	quick_sort(arrs, 0, len(arrs)-1)
// 	fmt.Printf("after arr: %v\n", arrs)
// }

func BenchmarkQuickSortThread(b *testing.B) {
	for n := 0; n < b.N; n++ {
		arrs := []int{3, 5, 2, 5, 7, 1, 3, 5, 25, 2, 15}
		quick_sort_thread(arrs, 0, len(arrs)-1)
		// fmt.Println(arrs)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		arrs := []int{3, 5, 2, 5, 7, 1, 3, 5, 25, 2, 15}
		quick_sort(arrs, 0, len(arrs)-1)
		// fmt.Println(arrs)
	}
}
