package linkedList

import lList "github.com/DickyMPriangga/LeetCodeSolution/linkedList/utilities"

///LeetCode Problems : Add Two Numbers
///Link : https://leetcode.com/problems/add-two-numbers/

func AddTwoNumbers(l1 *lList.ListNode, l2 *lList.ListNode) *lList.ListNode {
	var result lList.ListNode
	var p1, p2, pr *lList.ListNode
	var remainder = 0

	p1 = l1
	p2 = l2
	pr = &result

	for {
		if p1 == nil {
			pr.Val = p2.Val + remainder
			p2 = p2.Next
		} else if p2 == nil {
			pr.Val = p1.Val + remainder
			p1 = p1.Next
		} else {
			pr.Val = p1.Val + p2.Val + remainder
			p1 = p1.Next
			p2 = p2.Next
		}

		if pr.Val >= 10 {
			remainder = 1
			pr.Val = pr.Val % 10
		} else {
			remainder = 0
		}

		if p1 == nil && p2 == nil {
			if remainder == 1 {
				pr.Next = &lList.ListNode{Val: 1, Next: nil}
			}
			break
		} else {
			pr.Next = &lList.ListNode{}
			pr = pr.Next
		}
	}

	return &result
}
