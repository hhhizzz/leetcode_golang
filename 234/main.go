package _234

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	var data []int
	for head != nil {
		data = append(data, head.Val)
		head = head.Next
	}
	for i := 0; i < len(data)>>1; i++ {
		if data[i] != data[len(data)-i] {
			return false
		}
	}
	return true
}
