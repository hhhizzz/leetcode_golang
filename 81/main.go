package _81

func search(nums []int, target int) bool {
    if len(nums) <= 0 {
        return false
    }
    if len(nums) == 1 {
        return target == nums[0]
    }
    mid := len(nums) >> 1
    if nums[mid] == target {
        return true
    }
    if nums[0] < nums[mid] {
        //左边有序
        if target < nums[mid] && target >= nums[0] {
            return search(nums[:mid], target)
        } else {
            return search(nums[mid+1:], target)
        }
    } else if nums[mid] < nums[len(nums)-1] {
        //右边有序
        if target > nums[mid] && target <= nums[len(nums)-1] {
            return search(nums[mid+1:], target)
        } else {
            return search(nums[:mid], target)
        }
    } else {
        //两边都相等的情况，两边都判断
        return search(nums[:mid], target) || search(nums[mid+1:], target)
    }
}
