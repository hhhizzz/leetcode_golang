package _486

func canWin(nums []int, a, b int, aTerm bool) bool {
    if len(nums) == 0 {
        return a >= b
    } else {
        //如果是a的回合，只要有一种方式能保证能赢即可，因此用或
        if aTerm {
            return canWin(nums[1:], a+nums[0], b, !aTerm) || canWin(nums[:len(nums)-1], a+nums[len(nums)-1], b, !aTerm)
        } else {
            //如果是b的回合，需要保证不过他怎么选都不能赢，因此用与
            return canWin(nums[1:], a, b+nums[0], !aTerm) && canWin(nums[:len(nums)-1], a, b+nums[len(nums)-1], !aTerm)
        }
    }
}

func PredictTheWinner(nums []int) bool {
    return canWin(nums, 0, 0, true)
}
