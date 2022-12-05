package bubblesort

import (
	"fmt"
	"testing"
)

func TestBubble_sort(t *testing.T) {

	arr := []int{3, 1, 9, 5, 10, 6, 4, 8, 7}
	bubble_sort(arr)
	fmt.Println(arr)
}
