package _202

func isHappy(n int) bool {
    m := map[int]bool{}
    for n != 1 {
        if _, ok := m[n]; ok {
            return false
        }
        m[n] = true
        var list []int
        for n != 0 {
            list = append(list, n%10)
            n /= 10
        }
        for _, num := range list {
            n += num * num
        }
    }
    return true
}
