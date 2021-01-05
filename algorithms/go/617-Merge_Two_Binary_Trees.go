package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type nodePair struct {
	first *TreeNode
	second *TreeNode
}

func bfs(tree *TreeNode, f func(node *TreeNode))  {
	if tree == nil {
		return
	}
	queue := []*TreeNode{tree}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		f(node)
		if node != nil && (node.Left != nil || node.Right != nil) {
			queue = append(queue, node.Left, node.Right)
		}
	}
}

func treeValues(tree *TreeNode) []int {
	values := []int{}
	bfs(tree, func(node *TreeNode){
		if node == nil {
			values = append(values, -1)
		} else {
			values = append(values, node.Val)
		}
	})
	return values
}

// Inplace Iterative
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.Val += t2.Val
	queue := []nodePair{{t1, t2}}
	for len(queue) > 0 {
		pair := queue[0]
		node1 := pair.first
		node2 := pair.second
		queue = queue[1:]
		if node1.Left != nil && node2.Left != nil {
			node1.Left.Val += node2.Left.Val
			queue = append(queue, nodePair{node1.Left, node2.Left})
		} else if node2.Left != nil {
			node1.Left = node2.Left
		}
		if node1.Right != nil && node2.Right != nil {
			node1.Right.Val += node2.Right.Val
			queue = append(queue, nodePair{node1.Right, node2.Right})
		} else if node2.Right != nil {
			node1.Right = node2.Right
		}
	}
	return t1
}