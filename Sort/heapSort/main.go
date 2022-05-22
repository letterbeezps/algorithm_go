package main

import "fmt"

func main() {
	// 要将数组看作一个完全二叉树
	arrs := []int{3, 2, 7, 5, 9, 1, 7, 8, 4, 6, 5, 7}
	fmt.Printf("before arr is: %v\n", arrs)
	heapSort(arrs)
	fmt.Printf("after arr is: %v\n", arrs)
}

func heapSort(arrs []int) {
	l := len(arrs)
	initHeap(arrs, l)

	// 当前已经建立一个大根堆
	// 开始排序
	for i := l - 1; i > 0; i-- {
		// 取出堆顶元素
		arrs[0], arrs[i] = arrs[i], arrs[0]

		// 此时堆顶已经不再是最大的元素，需要对整个堆做一次sink
		// 待排序的堆是arrs[:l]
		l--
		sink(arrs, 0, l)
	}
}

// 初始化一个堆
func initHeap(arrs []int, l int) {
	for i := l / 2; i >= 0; i-- {
		sink(arrs, i, l)
	}
}

// sink可以生成一个以i为根节点的大根堆
func sink(arrs []int, i, l int) {
	// 左右子孩子
	left, right := 2*i+1, 2*i+2

	// largest用来记录i和两个子孩子中最大的那个
	largest := i

	if left < l && arrs[left] > arrs[largest] {
		largest = left
	}

	if right < l && arrs[right] > arrs[largest] {
		largest = right
	}

	if largest != i {
		arrs[largest], arrs[i] = arrs[i], arrs[largest]
		// 此时 arrs[i] 已经是 i, left, right三个里最大的了
		// largest的值变化了，要从新sink 以largest为root的堆
		sink(arrs, largest, l)
	}
}
