package _581

func findUnsortedSubarray(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    maxNums := make([]int, len(nums))
    minNums := make([]int, len(nums))
    maxNums[0] = nums[0]
    minNums[len(nums)-1] = nums[len(nums)-1]
    for i := 1; i < len(nums); i++ {
        if nums[i] > maxNums[i-1] {
            maxNums[i] = nums[i]
        } else {
            maxNums[i] = maxNums[i-1]
        }
    }
    for i := len(nums) - 2; i >= 0; i-- {
        if nums[i] < minNums[i+1] {
            minNums[i] = nums[i]
        } else {
            minNums[i] = minNums[i+1]
        }
    }

    left := 0
    right := len(nums) - 1
    for right >= 0 && nums[right] == maxNums[right] {
        right--
    }
    for left <= right && nums[left] == minNums[left] {
        left++
    }
    // fmt.Println(right)
    // fmt.Println(left)
    return right - left + 1
}
