package linkedList

import lList "github.com/DickyMPriangga/LeetCodeSolution/linkedList/utilities"

///LeetCode Problems : Remove Nth Node From End of List
///Link : https://leetcode.com/problems/remove-nth-node-from-end-of-list/

func RemoveNthFromEnd(head *lList.ListNode, n int) *lList.ListNode {
	var arrayNode []*lList.ListNode

	var listPtr *lList.ListNode = head
	for {
		if listPtr == nil {
			break
		}
		arrayNode = append(arrayNode, listPtr)
		listPtr = listPtr.Next
	}

	if len(arrayNode) == 1 {
		return nil
	}

	if n == 1 {
		arrayNode[len(arrayNode)-(n+1)].Next = nil
	} else if n == len(arrayNode) {
		head = arrayNode[1]
	} else {
		arrayNode[len(arrayNode)-(n+1)].Next = arrayNode[len(arrayNode)-(n-1)]
	}

	return head
}
