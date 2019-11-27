package _914

func hasGroupsSizeX(deck []int) bool {
    m := map[int]int{}
    for _, d := range deck {
        m[d] += 1
    }
    list := []int{}
    max := 0
    for _, v := range m {
        list = append(list, v)
        if v > max {
            max = v
        }
    }
    for i := 2; i <= max; i++ {
        exists := true
        for j := 0; j < len(list); j++ {
            if list[j]%i != 0 {
                exists = false
                break
            }
        }
        if exists {
            return true
        }
    }
    return false
}
