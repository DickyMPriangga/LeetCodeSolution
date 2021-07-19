package linkedList

import lList "github.com/DickyMPriangga/LeetCodeSolution/linkedList/utilities"

///LeetCode Problems : Merge Two Sorted Lists
///Link : https://leetcode.com/problems/merge-two-sorted-lists/

func MergeTwoLists(l1 *lList.ListNode, l2 *lList.ListNode) *lList.ListNode {
	if l1 == nil && l2 == nil {
		return nil
	} else if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else {
		if l1.Val < l2.Val {
			l1.Next = MergeTwoLists(l1.Next, l2)
			return l1
		} else {
			l2.Next = MergeTwoLists(l1, l2.Next)
			return l2
		}
	}
}
