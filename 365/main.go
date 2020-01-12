package _365

func gcd(a, b int) int {
    if a < b {
        a, b = b, a
    }
    if b == 0 {
        return 0
    }
    for a%b != 0 {
        a, b = a, a%b
        if a < b {
            a, b = b, a
        }
    }
    return b
}

func canMeasureWater(x int, y int, z int) bool {
    if z > x+y {
        return false
    }
    xy := gcd(x, y)
    if z == 0 {
        return true
    } else if xy == 0 {
        return false
    }
    return z%xy == 0
}
