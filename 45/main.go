package _45

func max(a, b int) int {
    if a>b{
        return a
    }
    return b
}

//题解里说这算是贪心，我认为应该是动态规划，因为其实并不是一直选最远的地方，而是选下一步能够到达的地方，有一个状态转移的过程
func jump(nums []int) int {
    step := 0
    //maxPos表示当前能够到达的最远的地方，也可以认为是当前这里通过两步能到达最远的地方
    maxPos := 0
    //end表示目前到的这一步最远到的地方
    end := 0
    //如果到达len(nums)-1表示到达目的地，不需要再更新step了
    for i := 0; i < len(nums)-1; i++ {
        maxPos = max(nums[i]+i,maxPos)
        //如果到达这步的边界，就更新到新的边界，同时更新step
        if i==end{
            step++
            end = maxPos
        }
    }
    return step
}
