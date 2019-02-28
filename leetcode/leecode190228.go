package leetcode

import "fmt"

func LeetTest190228() {
	l1head := &ListNode{2, &ListNode{2, &ListNode{6, nil}}}
	l2head := &ListNode{1, &ListNode{3, &ListNode{3, nil}}}

	res := mergeTwoLists(l1head, l2head)

	for {
		if res == nil {
			break
		}

		fmt.Printf("%d  ", res.val)
		res = res.next
	}
}

// 21
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	var one, two *ListNode
	if l1.val <= l2.val {
		one = l1
		two = l2
	} else {
		one = l2
		two = l1
	}

	l1now := one
	l1pre := one
	l2now := two
	for {
		if l2now == nil {
			break
		}

		if l1now.val <= l2now.val {
			if l1now.next == nil {
				l1now.next = l2now
				break
			} else {
				l1pre = l1now
				l1now = l1now.next
			}

		} else {
			l1pre.next = l2now
			l2next := l2now.next

			l2now.next = l1now
			l1pre = l2now

			l2now = l2next
		}
	}
	return one
}
