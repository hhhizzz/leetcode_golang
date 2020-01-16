package _325

import "fmt"

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
    sumPos := map[int]int{}
    sums := make([]int, len(nums))

    result := 0
    for i := 0; i < len(nums); i++ {
        if i == 0 {
            sums[i] = nums[i]
        } else {
            sums[i] = nums[i] + sums[i-1]
        }
        if _, ok := sumPos[sums[i]]; !ok {
            sumPos[sums[i]] = i
        }
    }
    fmt.Println(sumPos)
    for i := 0; i < len(nums); i++ {
        if sums[i] == k {
            result = i + 1
        } else {
            if pos, ok := sumPos[sums[i]-k]; ok && i-pos > result {
                //fmt.Printf("%d %d\n",i,pos)
                result = i - pos
            }
        }
    }

    return result
}
