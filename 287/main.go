package _287

//快慢指针法
func findDuplicate(nums []int) int {
	slow := 0
	fast := 0
	for slow == 0 || slow != fast {
		fast = nums[fast]
		fast = nums[fast]
		slow = nums[slow]
	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
