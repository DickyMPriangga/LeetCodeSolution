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
		result = append(result, listPtr.Val)
		listPtr = listPtr.Next
	}

	return result
}
