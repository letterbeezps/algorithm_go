package main

import (
	"fmt"
	"sync"
)

func main() {

	arrs := []int{3, 5, 2, 5, 7, 1, 3, 5, 25, 2, 15}
	fmt.Printf("before arr: %v\n", arrs)

	quick_sort(arrs, 0, len(arrs)-1)
	fmt.Printf("after arr: %v\n", arrs)

	arrs2 := []int{3, 5, 2, 5, 7, 1, 3, 5, 25, 2, 15}
	fmt.Printf("before arr: %v\n", arrs2)

	quick_sort_thread(arrs2, 0, len(arrs2)-1)
	fmt.Printf("after arr: %v\n", arrs2)
}

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

func quick_sort_thread(arrs []int, l, r int) {
	if l >= r {
		return
	}

	var wg sync.WaitGroup
	left, right := l-1, r+1
	mid := (l + r) / 2
	x := arrs[mid]

	for left < right {
		left++
		for arrs[left] < x {
			left++
		}

		right--
		for arrs[right] > x {
			right--
		}

		if left < right {
			arrs[left], arrs[right] = arrs[right], arrs[left]
		}
	}

	wg.Add(2)
	go func() {
		quick_sort_thread(arrs, l, right)
		wg.Done()
	}()

	go func() {
		quick_sort_thread(arrs, right+1, r)
		wg.Done()
	}()

	wg.Wait()
}
