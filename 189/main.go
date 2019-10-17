package _189

func rotate2(nums []int, k int) {
    k %= len(nums)
    start := 0
    next := start
    temp := nums[start]
    for i := 0; i < len(nums); i++ {
        next = (next + k) % len(nums)
        temp2 := nums[next]
        nums[next] = temp
        temp = temp2
        if i != 0 && next == start {
            start += 1
            temp = nums[start]
            next = start
        }
    }
    nums[next] = temp
}

func reverse(nums []int) {
    for i := 0; i < len(nums)>>1; i++ {
        nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
    }
}

func rotate(nums []int, k int) {
    k %= len(nums)
    reverse(nums)
    reverse(nums[:k])
    reverse(nums[k:])
}
