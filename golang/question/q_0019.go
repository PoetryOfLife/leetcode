package question

func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	var nodes []*ListNode
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	for node := dummy; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	prev := nodes[len(nodes)-1-n]
	prev.Next = prev.Next.Next
	return dummy.Next
}
