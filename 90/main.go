package _90

//巧用异或算法符合交换律和结合律的特性
func singleNumber(nums []int) int {
    a := 0
    for _, nums := range nums {
        a ^= nums
    }
    return a
}
