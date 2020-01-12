package _365

func gcd(a, b int) int {
    if a < b {
        a, b = b, a
    }
    if b == 0 {
        return 0
    }
    if a%b == 0 {
        return b
    }
    return gcd(b, a-b)
}

func canMeasureWater(x int, y int, z int) bool {
    if z > x+y {
        return false
    }
    xy := gcd(x, y)
    //fmt.Println(xy)
    if z == 0 {
        return true
    } else if xy == 0 {
        return false
    }
    return z%xy == 0
}
