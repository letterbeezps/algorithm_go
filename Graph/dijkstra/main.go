package main

import "container/heap"

// items [[2,1,1],[2,3,1],[3,4,1]] [u, v, w] 表示节点u于v相邻，且权重为w
func dijkstra(times [][]int, n int, k int) int {

	// 构建图，邻接表结构
	graph := make(map[int][][]int)
	for i := 0; i < len(times); i++ {
		graph[times[i][0]] = append(graph[times[i][0]], times[i][1:])
	}

	// 存储点k到每个节点的最小距离
	distince := make(map[int]int)

	// fmt.Println(graph)

	pq := make(PriorityQueue, 0)
	// 起始节点k, k到k的距离为0
	heap.Push(&pq, &Item{
		value:    []int{0, k},
		priority: 0,
	})

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		dis, node := item.value[0], item.value[1]
		if _, ok := distince[node]; ok {
			continue
		}

		distince[node] = dis

		for i := 0; i < len(graph[node]); i++ {
			nextNode, nextDis := graph[node][i][0], graph[node][i][1]
			if _, ok := distince[nextNode]; !ok {
				newItem := &Item{
					value:    []int{nextDis + dis, nextNode},
					priority: nextDis + dis,
				}
				heap.Push(&pq, newItem)
			}
		}
	}
	// distance[i]表示节点k到节点i的最短路径
	if len(distince) != n {
		return -1
	} else {
		res := -1
		for _, v := range distince {
			res = max(res, v)
		}
		return res
	}

}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

type Item struct {
	value    []int
	priority int
	index    int
}
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = j
	pq[j].index = i
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index -= 1
	*pq = old[0 : n-1]
	return item
}
