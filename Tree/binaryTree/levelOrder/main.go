package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := [][]int{}

	level := []*TreeNode{}
	level = append(level, root)
	res = append(res, getVal(level))

	for true {
		newLevel := []*TreeNode{}
		for i := 0; i < len(level); i++ {
			x := level[i]
			if x.Left != nil {
				newLevel = append(newLevel, x.Left)
			}
			if x.Right != nil {
				newLevel = append(newLevel, x.Right)
			}
		}
		if len(newLevel) > 0 {
			res = append(res, getVal(newLevel))
			level = newLevel
		} else {
			break
		}
	}

	return res
}

func getVal(level []*TreeNode) []int {
	res := []int{}
	for i := 0; i < len(level); i++ {
		res = append(res, level[i].Val)
	}
	return res
}
