package _89

func grayCode(n int) []int {
    m := map[int]bool{0: true}
    result := []int{0}

    var mask []int
    num := 1
    for i := 1; i <= n; i++ {
        mask = append(mask, num)
        num = num << 1
    }

    for i := 1; i < num; i++ {
        for _, ma := range mask {
            current := result[len(result)-1] ^ ma
            if _, ok := m[current]; !ok {
                m[current] = true
                result = append(result, current)
                break
            }
        }
    }
    return result
}
