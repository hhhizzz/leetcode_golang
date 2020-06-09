package _817

type ListNode struct {
	Val  int
	Next *ListNode
}

func numComponents(head *ListNode, G []int) int {
	m := map[int]bool{}
	for _, g := range G {
		m[g] = true
	}
	current := head
	currentSum := 0

	result := 0
	for current != nil {
		if _, ok := m[current.Val]; ok {
			currentSum += 1
		} else {
			if currentSum > 0 {
				result++
			}
			currentSum = 0
		}
		current = current.Next
	}
	if currentSum > 0 {
		result++
	}
	return result
}
