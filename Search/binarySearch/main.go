package main

import (
	"fmt"
	"sort"
)

var arrs = []int{1, 2, 4, 6, 8, 8, 8}

func main() {

	fmt.Printf("before arr: %v\n", arrs)

	up_equal := binary_search_up_equal(arrs, 9)

	fmt.Printf("the index of element which first ge %d is %d\n", 9, up_equal)

	index := sort.SearchInts(arrs, 8)
	fmt.Println(index)

	up := binary_search_up(arrs, 8)

	fmt.Printf("the index of element which first g %d is %d\n", 8, up)

	down_equal := binary_search_down_equal(arrs, 8)

	fmt.Printf("the index of element which first le %d is %d\n", 8, down_equal)

	down := binary_search_down(arrs, 8)

	fmt.Printf("the index of element which first l %d is %d\n", 8, down)
}

// 第一个大于等于 8
func binary_search_up_equal(arrs []int, target int) int {
	l, r := 0, len(arrs)
	for l < r {
		mid := (l + r) / 2
		if arrs[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

// 第一个大于 8
func binary_search_up(arrs []int, target int) int {
	l, r := 0, len(arrs)-1
	for l < r {
		mid := (l + r) / 2
		if arrs[mid] <= target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

// 最后一个小于等于 8
func binary_search_down_equal(arrs []int, target int) int {
	l, r := 0, len(arrs)-1
	for l < r {
		mid := (l + r + 1) / 2
		if arrs[mid] > target {
			r = mid - 1
		} else {
			l = mid
		}
	}
	return l
}

// 最后一个小于 8
func binary_search_down(arrs []int, target int) int {
	l, r := 0, len(arrs)-1
	for l < r {
		mid := (l + r + 1) / 2
		if arrs[mid] >= target {
			r = mid - 1
		} else {
			l = mid
		}
	}
	return l
}
