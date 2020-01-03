package _403

//方法一：直接能够想到的O(n^2)的算法
func canCross1(stones []int) bool {
    step := make([]map[int]bool, len(stones))
    for i := 0; i < len(stones); i++ {
        step[i] = map[int]bool{}
    }
    step[0][0] = true
    for i := 1; i < len(stones); i++ {
        for j := i - 1; j >= 0; j-- {
            interval := stones[i] - stones[j]
            //上一步能够k步到达
            if _, ok := step[j][interval]; ok {
                step[i][interval] = true
                continue
            }
            //上一步能够k+1步到达
            if _, ok := step[j][interval-1]; ok {
                step[i][interval] = true
                continue
            }
            //上一步能够k-1步到达
            if _, ok := step[j][interval+1]; ok {
                step[i][interval] = true
                continue
            }
        }
    }
    //fmt.Println(step)
    return len(step[len(step)-1]) > 0
}

//方法一换了个方向，复杂度没变，但是性能提升了很多
func canCross(stones []int) bool {
    step := make([]map[int]bool, len(stones))
    stonesMap := map[int]int{}
    for i := 0; i < len(stones); i++ {
        stonesMap[stones[i]] = i
        step[i] = map[int]bool{}
    }
    step[0][0] = true
    for i := 0; i < len(stones); i++ {
        for k := range step[i] {
            if nextPos, ok := stonesMap[stones[i]+k-1]; stones[i]+k-1 > 0 && ok {
                step[nextPos][k-1] = true
            }
            if nextPos, ok := stonesMap[stones[i]+k]; ok {
                step[nextPos][k] = true
            }
            if nextPos, ok := stonesMap[stones[i]+k+1]; ok {
                step[nextPos][k+1] = true
            }
        }
    }
    //fmt.Println(step)
    return len(step[len(step)-1]) > 0
}
