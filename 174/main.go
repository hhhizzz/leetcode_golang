package _174

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func calculateMinimumHP(dungeon [][]int) int {
    dp := make([][]int, len(dungeon))
    for i := 0; i < len(dp); i++ {
        dp[i] = make([]int, len(dungeon[i]))
    }
    for i := len(dungeon) - 1; i >= 0; i-- {
        for j := len(dungeon[i]) - 1; j >= 0; j-- {
            var nextHealth int
            if i == len(dungeon)-1 {
                if j == len(dungeon[i])-1 {
                    if dungeon[i][j] >= 0 {
                        nextHealth = 1 - dungeon[i][j]
                    } else {
                        nextHealth = 1
                    }
                } else {
                    nextHealth = dp[i][j+1]
                }
            } else {
                if j == len(dungeon[i])-1 {
                    nextHealth = dp[i+1][j]
                } else {
                    nextHealth = min(dp[i+1][j], dp[i][j+1])
                }
            }
            dp[i][j] = nextHealth - dungeon[i][j]
            if dp[i][j] < 1 {
                dp[i][j] = 1
            }
        }
    }
    //fmt.Println(dp)
    return dp[0][0]
}
