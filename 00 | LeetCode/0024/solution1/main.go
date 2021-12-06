package solution1

import (
	. "_/definition/single-linkedlist"
)

func swapPairs(head *ListNode) *ListNode {
	if done(head) {
		return head
	} else {
		return swap(head)
	}
}

func done(head *ListNode) bool {
	return head == nil || head.Next == nil
}

// 1 -> 2 -> 3 -> nil
func swap(one *ListNode) *ListNode {
	two := one.Next
	// 1 -> 3 -> nil
	one.Next = swapPairs(two.Next)
	// 2 -> 1 -> 3 -> nil
	two.Next = one
	return two
}