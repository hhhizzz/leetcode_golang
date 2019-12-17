package _7

import "math"

func reverse(x int) int {
    result := 0
    for x != 0 {
        current := x % 10
        result += current

        x /= 10
        if x != 0 {
            result *= 10
        }
    }
    if result < math.MinInt32 || result > math.MaxInt32 {
        return 0
    }
    return result
}
