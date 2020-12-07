package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewNode(val int, left *TreeNode, right *TreeNode) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  left,
		Right: right,
	}
}

const (
	WHITE = 0
	GRAY  = 1
	BLACK = 2
)

func buildTree(values []int) *TreeNode {
	if len(values) == 0 {
		return nil
	}
	root := NewNode(values[0], nil, nil)
	queue := []*TreeNode{root}
	for i := 1; i < len(values); i += 2 {
		node := queue[0]
		if i < len(values) && values[i] != -1 {
			node.Left = NewNode(values[i], nil, nil)
			queue = append(queue, node.Left)
		}
		if i+1 < len(values) && values[i+1] != -1 {
			node.Right = NewNode(values[i+1], nil, nil)
			queue = append(queue, node.Right)
		}
		queue = queue[1:]
	}
	return root
}

func isPalindromic(values []int) bool {
	counter := map[int]int{}
	for _, val := range values {
		counter[val]++
	}
	once := false
	for _, v := range counter {
		if v%2 != 0 {
			if len(values)%2 == 0 {
				return false
			} else if !once {
				once = true
			} else {
				return false
			}
		}
	}
	return true
}

func pseudoPalindromicPaths(root *TreeNode) int {
	pathsCount := 0
	if root == nil {
		return pathsCount
	}
	stack := []*TreeNode{root}
	nodeColors := map[*TreeNode]int{root: WHITE}
	path := []int{}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		color := nodeColors[node]
		if color == WHITE {
			nodeColors[node] = GRAY
			path = append(path, node.Val)
			if node.Left != nil && nodeColors[node.Left] == WHITE {
				stack = append(stack, node.Left)
			}
			if node.Right != nil && nodeColors[node.Right] == WHITE {
				stack = append(stack, node.Right)
			}

		} else if color == GRAY {
			nodeColors[node] = BLACK
			if node.Left == nil && node.Right == nil {
				if isPalindromic(path) {
					pathsCount++
				}
			}
			stack = stack[:len(stack)-1]
			path = path[:len(path)-1]
		}
	}
	return pathsCount
}
