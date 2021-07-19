package linkedList

import lList "github.com/DickyMPriangga/LeetCodeSolution/linkedList/utilities"

///LeetCode Problems : Swap Nodes in Pairs
///Link : https://leetcode.com/problems/swap-nodes-in-pairs/

func SwapPairs(head *lList.ListNode) *lList.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	ptr1 := head
	ptr2 := head.Next

	ptr1.Next = SwapPairs(ptr2.Next)
	ptr2.Next = ptr1

	return ptr2
}
