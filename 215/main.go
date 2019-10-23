package _215

func findKthLargest(nums []int, k int) int {
    pivot := 0
    for i := 1; i < len(nums); i++ {
        if nums[i] < nums[0] {
            pivot++
            nums[pivot], nums[i] = nums[i], nums[pivot]
        }
    }
    nums[pivot], nums[0] = nums[0], nums[pivot]
    if pivot == len(nums)-k {
        return nums[pivot]
    } else if pivot < len(nums)-k {
        return findKthLargest(nums[pivot+1:], k)
    } else {
        return findKthLargest(nums[:pivot], k-(len(nums)-pivot))
    }
}
