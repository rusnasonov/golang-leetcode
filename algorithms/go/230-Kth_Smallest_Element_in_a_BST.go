package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Iterative
func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	result := []int{}
	node := root
	for len(stack) != 0 || node != nil || len(result) < k {
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
	return result[k-1]
}

// Channel
func kthSmallestChannel(root *TreeNode, k int) int {
	for item := range inorderChannel(root) {
		k--
		if k == 0 {
			return item
		}
	}
	return 0
}

func inorderChannel(root *TreeNode) chan int {
	stack := []*TreeNode{}
	ch := make(chan int)
	go func(){
		for len(stack) != 0 || root != nil {
			for root != nil {
				stack = append(stack, root)
				root = root.Left
			}
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ch <- root.Val
			root = root.Right
		}
	}()
	return ch
}

// Recursive
func kthSmallestRecursive(root *TreeNode, k int) int {
	return inorder(root, []int{})[k-1]
}

func inorder(tree *TreeNode, res []int) []int {
	if tree == nil {
		return res
	}
	res = inorder(tree.Left, res)
	res = append(res, tree.Val)
	res = inorder(tree.Right, res)
	return res
}