package _34

func searchRange(nums []int, target int) []int {
    if len(nums) == 0 {
        return []int{-1, -1}
    }
    left := 0
    right := len(nums) - 1
    targetLeft := 0
    targetRight := 0
    for left < right-1 {
        mid := left + ((right - left) >> 1)
        if nums[mid] >= target {
            right = mid
        } else {
            left = mid
        }
    }
    if nums[left] == target {
        targetLeft = left
    } else if nums[right] == target {
        targetLeft = right
    } else {
        return []int{-1, -1}
    }
    left = 0
    right = len(nums) - 1
    for left < right-1 {
        mid := left + ((right - left) >> 1)
        if nums[mid] <= target {
            left = mid
        } else {
            right = mid
        }
    }
    if nums[right] == target {
        targetRight = right
    } else if nums[left] == target {
        targetRight = left
    }
    return []int{targetLeft, targetRight}
}
