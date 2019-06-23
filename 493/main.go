package _493

func merge(nums []int, left, right int) int {
    if left >= right-1 {
        return 0
    }
    result := 0
    mid := (left + right) >> 1
    result += merge(nums, left, mid)
    result += merge(nums, mid, right)
    length := right - left
    temp := make([]int, length)
    lIndex := left
    rIndex := mid
    rIndex2 := mid
    index := 0
    for lIndex < mid {
        for rIndex2 < right && nums[lIndex] > 2*nums[rIndex2] {
            rIndex2++
        }
        result += rIndex2 - mid
        for rIndex < right && nums[lIndex] > nums[rIndex] {
            temp[index] = nums[rIndex]
            index++
            rIndex++
        }
        temp[index] = nums[lIndex]
        index++
        lIndex++
    }
    for rIndex < right {
        temp[index] = nums[rIndex]
        index++
        rIndex++
    }
    for i := 0; i < length; i++ {
        nums[left+i] = temp[i]
    }
    return result
}

func reversePairs(nums []int) int {
    return merge(nums, 0, len(nums))
}
