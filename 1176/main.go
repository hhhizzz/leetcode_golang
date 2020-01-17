package _1176

func dietPlanPerformance(calories []int, k int, lower int, upper int) int {
    window := 0
    for i := 0; i < k-1; i++ {
        window += calories[i]
    }
    point := 0
    for i := k - 1; i < len(calories); i++ {
        if i != k-1 {
            window -= calories[i-k]
        }
        window += calories[i]
        //fmt.Println(window)
        if window < lower {
            point -= 1
        } else if window > upper {
            point += 1
        }
    }
    return point
}
