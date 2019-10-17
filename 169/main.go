package _169

func majorityElement(nums []int) int {
    m := map[int]int{}
    half := len(nums) >> 1

    for _, num := range nums {
        if _, ok := m[num]; ok {
            m[num] += 1
        } else {
            m[num] = 1
        }
        if m[num] > half {
            return num
        }
    }
    return -1
}
