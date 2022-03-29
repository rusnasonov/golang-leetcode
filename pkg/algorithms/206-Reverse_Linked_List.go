package algorithms

type ListNode struct {
     Val int
     Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	rev := &ListNode{Val: head.Val, Next: nil}
	for node:=head.Next; node != nil; node = node.Next {
		curr := &ListNode{}
		curr.Val = node.Val
		curr.Next = rev
		rev = curr
	}
	return rev
}

func reverseListRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
    out := reverseListRecursive(head.Next);
    head.Next.Next = head
    head.Next = nil;
    return out;
}

func sliceToLinkedList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	list := &ListNode{Val: arr[0]}
	node := list
	for i := 1; i < len(arr); i++ {
		node.Next = &ListNode{Val: arr[i]}
		node = node.Next
	}
	return list
}

func linkedListToSlice(list *ListNode) []int {
	out := []int{}
	if list == nil {
		return out
	}
	for node:=list; node != nil; node = node.Next {
		out = append(out, node.Val)
	}
	return out
}