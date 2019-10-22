package _334

func increasingTriplet(nums []int) bool {
    var current []int

    for _, num := range nums {
        if len(current) == 0 || current[len(current)-1] < num {
            current = append(current, num)
        } else {
            for i := len(current) - 1; i >= 0; i-- {
                if i == 0 || num > current[i-1] {
                    current[i] = num
                    break
                }
            }
        }
        if len(current) >= 3 {
            return true
        }
    }
    return false
}
