package Q0002

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	p := head
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		a := 0
		b := 0
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}
		carry += a + b
		p.Next = &ListNode{carry % 10, nil}
		p = p.Next
		carry /= 10
	}
	if carry != 0 {
		p.Next = &ListNode{carry, nil}
	}
	return head.Next
}
