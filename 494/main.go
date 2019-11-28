package _494

func findTargetSumWays(nums []int, S int) int {
    m := map[int]int{0: 1}
    for i := 0; i < len(nums); i++ {
        num := nums[i]
        newM := map[int]int{}
        for k := range m {
            newM[k+num] += m[k]
            newM[k-num] += m[k]
        }
        m = newM
    }
    return m[S]
}
