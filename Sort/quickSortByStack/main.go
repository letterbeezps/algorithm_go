package main

import "fmt"

func main() {
	arrs := []int{3, 5, 2, 5, 7, 1, 3, 5, 25, 2, 15}
	// arrs := []int{3,5,2,5,7}
	fmt.Printf("before arr: %v\n", arrs)

	quick_sort(arrs, 0, len(arrs)-1)
	fmt.Printf("after arr: %v\n", arrs)
}

func partition(arr []int, l int, r int) int {
	var i, j int = l - 1, r + 1
	var mid int = (l + r) / 2

	x := arr[mid]
	for i < j {
		i++
		for arr[i] < x {
			i++
		}

		j--
		for arr[j] > x {
			j--
		}

		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	return j
}

func quick_sort(arr []int, l int, r int) {
	if r <= 0 || l >= r {
		return
	}
	stack := []int{}
	i := 0
	j := 0
	stack = append(stack, r)
	stack = append(stack, l)

	for len(stack) > 0 {
		i = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		j = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if i < j {
			k := partition(arr, i, j)
			if k > i {
				stack = append(stack, k)
				stack = append(stack, i)
			}
			if k < j {
				stack = append(stack, j)
				stack = append(stack, k+1)
			}
		}
	}
}
