package _325

//O(n^2)的算法 加了剪枝
func maxSubArrayLen1(nums []int, k int) int {
    sums := make([]int, len(nums))
    for i := 0; i < len(nums); i++ {
        if i == 0 {
            sums[i] = nums[i]
        } else {
            sums[i] = sums[i-1] + nums[i]
        }
    }

    result := 0
    for i := 0; i < len(nums); i++ {
        if sums[i] == k {
            result = i + 1
        } else {
            for j := i - result; j >= 0; j-- {
                if j != 0 && sums[i]-sums[j-1] == k {
                    result = i - j + 1
                }
            }
        }
    }
    return result
}
