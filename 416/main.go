package _416

func canPartition(nums []int) bool {
    if len(nums) < 2 {
        return false
    }
    sum := 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
    }
    if sum%2 != 0 {
        return false
    }
    target := make([][]bool, (sum>>1)+1)
    for i := 0; i < len(target); i++ {
        target[i] = make([]bool, len(nums))
        for j := 0; j < len(nums); j++ {
            if j > 0 && target[i][j-1] == true {
                target[i][j] = true
            } else if i-nums[j] == 0 {
                target[i][j] = true
            } else if j > 0 && i-nums[j] > 0 {
                target[i][j] = target[i-nums[j]][j-1]
            }
        }
    }

    return target[sum>>1][len(nums)-1]
}
