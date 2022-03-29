package algorithms

type node struct {
	treeNode *TreeNode
	depth int
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 1
	stack := []*node{{root,1}}
	for len(stack) > 0 {
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if n.treeNode.Right != nil {
			stack = append(stack, &node{n.treeNode.Right, n.depth+1})
		}
		if n.treeNode.Left != nil {
			stack = append(stack, &node{n.treeNode.Left, n.depth+1})
		}

		if n.depth > depth {
			depth = n.depth
		}
	}
	return depth
}

func maxDepthRecursive(tree *TreeNode) int {
	if tree == nil {
		return 0
	}
	lchild := maxDepth(tree.Left);
	rchild := maxDepth(tree.Right);
	if lchild <= rchild {
		return rchild+1
	} else {
		return lchild+1
	}
}
