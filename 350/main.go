package _350

func intersect(nums1 []int, nums2 []int) []int {
    m1 := map[int]int{}
    m2 := map[int]int{}
    for _, num := range nums1 {
        if value, ok := m1[num]; ok {
            m1[num] = value + 1
        } else {
            m1[num] = 1
        }
    }
    for _, num := range nums2 {
        if value, ok := m2[num]; ok {
            m2[num] = value + 1
        } else {
            m2[num] = 1
        }
    }
    var result []int
    for key, value := range m1 {
        if value2, ok := m2[key]; ok {
            length := value
            if value2 < length {
                length = value2
            }
            for i := 0; i < length; i++ {
                result = append(result, key)
            }
        }
    }
    return result
}
