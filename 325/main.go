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

//用map存储可以降到O(n)
func maxSubArrayLen(nums []int, k int) int {
    sum := make([]int, len(nums))
    sumPos := map[int]int{0: -1}

    for i := 0; i < len(nums); i++ {
        if i == 0 {
            sum[i] = nums[i]
        } else {
            sum[i] = sum[i-1] + nums[i]
        }
        if _, ok := sumPos[sum[i]]; !ok {
            sumPos[sum[i]] = i
        }
    }
    result := 0
    for i := 0; i < len(sum); i++ {
        diff := sum[i] - k
        if pos, ok := sumPos[diff]; ok && pos < i {
            if i-pos > result {
                result = i - pos
            }
        }
    }
    return result
}
