package _15

func qsort(nums []int) {
    if len(nums) <= 1 {
        return
    }
    pivotIndex := 0
    for i := 1; i < len(nums); i++ {
        if nums[i] < nums[0] {
            pivotIndex++
            nums[pivotIndex], nums[i] = nums[i], nums[pivotIndex]
        }
    }
    nums[0], nums[pivotIndex] = nums[pivotIndex], nums[0]
    qsort(nums[:pivotIndex])
    qsort(nums[pivotIndex+1:])
}

func threeSum(nums []int) [][]int {
    if len(nums) < 3 {
        return make([][]int, 0)
    }
    qsort(nums)
    res := make(map[[3]int]bool)
    for i, v := range nums[:len(nums)-2] {
        if i >= 1 && nums[i] == nums[i-1] {
            continue
        }
        l := i + 1
        r := len(nums) - 1
        for l < r {
            sum := v + nums[l] + nums[r]
            if sum == 0 {
                res[[3]int{v, nums[l], nums[r]}] = true
                for l < r && nums[l] == nums[l+1] {
                    l++
                }
                for l < r && nums[r] == nums[r-1] {
                    r--
                }
                l++
                r--
            } else if sum > 0 {
                r--
            } else if sum < 0 {
                l++
            }
        }
    }
    var result [][]int
    for k := range res {
        slice := make([]int, 3)
        copy(slice, k[:])
        result = append(result, slice)
    }
    return result
}
