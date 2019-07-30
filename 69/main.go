package _69

func mySqrt(x int) int {
    l := 1
    r := x
    for l < r-1 {
        mid := (l + r) >> 1
        if mid*mid < x {
            l = mid
        } else {
            r = mid
        }
    }
    if r*r == x {
        return r
    } else {
        return l
    }
}

