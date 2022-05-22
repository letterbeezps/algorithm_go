package main

import "fmt"

func main() {

	arr := []int{3, 0, 2, 1, 4}

	for i := 0; i < len(arr); i++ {
		arr[i], arr[arr[i]] = arr[arr[i]], arr[i]
	}
	fmt.Println(arr)
}
