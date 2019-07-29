package _169

func majorityElement(nums []int) int {
    m := make(map[int]int)
    n2 := len(nums) >> 1
    for _,num := range nums {
        if _, ok := m[num]; ok {
            m[num] += 1;
        } else {
            m[num] = 1;
        }
        if m[num] > n2 {
            return num
        }
    }
    return nums[len(nums)-1]
}
