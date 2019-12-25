package _162

func findPeakElement(nums []int) int {
    left := 0
    right := len(nums)
    for left < right {
        mid := (left + right) >> 1
        if mid != len(nums)-1 && nums[mid] < nums[mid+1] {
            left = mid + 1
        } else if mid == 0 || nums[mid-1] < nums[mid] {
            return mid
        } else {
            right = mid
        }
    }
    return -1
}
