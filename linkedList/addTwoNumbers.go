package linkedList

type ListNode struct {
	Val  int
	Next *ListNode
}

func ListFromArray(arr []int) ListNode {
	list := new(ListNode)
	listPtr := list
	for i := 0; i < len(arr); i++ {
		listPtr.Val = arr[i]
		if i != len(arr)-1 {
			listPtr.Next = &ListNode{}
			listPtr = listPtr.Next
		}
	}
	return *list
}

func (l *ListNode) ListToArray() []int {
	var result []int
	var listPtr *ListNode = l
	for {
		if listPtr == nil {
			break
		}
		result = append(result, l.Val)
		listPtr = listPtr.Next
	}

	return result
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result ListNode
	var p1, p2, pr *ListNode
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
				pr.Next = &ListNode{Val: 1, Next: nil}
			}
			break
		} else {
			pr.Next = &ListNode{}
			pr = pr.Next
		}
	}

	return &result
}
