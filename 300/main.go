package _300

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func find(nums *[]int, key int) int {
    if len(*nums) == 0 {
        return 0
    }
    left, right := 0, len(*nums)
    for left < right {
        mid := (left + right) / 2
        if (*nums)[mid] >= key {
            right = mid
        } else {
            left = mid+1
        }
    }
    if right < len(*nums) && key <= (*nums)[right] {
        return left
    } else {
        return right
    }
}

func lengthOfLIS(nums []int) int {
    lis := make([]int, 0)
    result := 0
    for i := 0; i < len(nums); i++ {
        pos := find(&lis, nums[i])
        if pos >= len(lis) {
            lis = append(lis, nums[i])
        } else {
            lis := lis[:pos+1]
            lis[pos] = nums[i]
        }
        result = max(result, len(lis))
    }
    return result
}
