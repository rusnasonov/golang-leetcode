package algorithms


func inorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{}
	result := []int{}
	node := root
	for len(stack) != 0 || node != nil {
		if node != nil {
			stack = append(stack, node)
			node = node.Left
		} else {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = append(result, node.Val)
			node = node.Right
		}
	}
	return result
}