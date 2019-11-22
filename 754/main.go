package _754

func reachNumber(target int) int {
    if target < 0 {
        target = -target
    }
    step := 0
    for target > 0 {
        step += 1
        target -= step
    }
    if target%2 == 0 {
        return step
    }
    if (step+1)%2 == 1 {
        //说明下一个step和target同为奇数
        return step + 1
    } else {
        return step + 2
    }
}
