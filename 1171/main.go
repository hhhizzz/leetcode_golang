package _1171

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeZeroSumSublists(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	current := head
	currentSum := 0

	result := head
	for current != nil {
		currentSum += current.Val
		if currentSum == 0 {
			//这里有一个小点需要注意的是这里不需要break，要继续找出和为0的点
			result = current.Next
		}
		current = current.Next
	}
	if result != nil {
		result.Next = removeZeroSumSublists(result.Next)
	}

	return result
}
