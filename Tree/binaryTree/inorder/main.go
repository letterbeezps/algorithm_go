package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := []int{}
	stack := []*TreeNode{}
	cur := root
	for len(stack) > 0 || cur != nil {

		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		if len(stack) > 0 {
			cur = stack[len(stack)-1]
			res = append(res, cur.Val)
			stack = stack[:len(stack)-1]
			cur = cur.Right
		}
	}
	return res
}

func inorderTraversalDFS(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := []int{}
	dfs(root, &res)

	return res
}

func dfs(root *TreeNode, res *[]int) {
	if root != nil {
		dfs(root.Left, res)
		*res = append(*res, root.Val)
		dfs(root.Right, res)
	}
}
