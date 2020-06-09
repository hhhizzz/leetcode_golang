package _1182

func shortestDistanceColor(colors []int, queries [][]int) []int {
	dp := make([][]int, 3)
	for i := 0; i < 3; i++ {
		dp[i] = make([]int, len(colors))
	}
	color1 := -1
	color2 := -1
	color3 := -1
	for i := 0; i < len(colors); i++ {
		if colors[i] == 1 {
			color1 = i
		} else if colors[i] == 2 {
			color2 = i
		} else if colors[i] == 3 {
			color3 = i
		}
		dp[0][i] = color1
		dp[1][i] = color2
		dp[2][i] = color3
	}
	color1 = -1
	color2 = -1
	color3 = -1
	for i := len(colors) - 1; i >= 0; i-- {
		if colors[i] == 1 {
			color1 = i
		} else if colors[i] == 2 {
			color2 = i
		} else if colors[i] == 3 {
			color3 = i
		}
		if color1 != -1 && (dp[0][i] == -1 || color1-i < i-dp[0][i]) {
			dp[0][i] = color1
		}
		if color2 != -1 && (dp[1][i] == -1 || color2-i < i-dp[1][i]) {
			dp[1][i] = color2
		}
		if color3 != -1 && (dp[2][i] == -1 || color3-i < i-dp[2][i]) {
			dp[2][i] = color3
		}
	}
	//fmt.Println(dp)
	result := make([]int, len(queries))
	for i, query := range queries {
		pos := dp[query[1]-1][query[0]]
		if pos == -1 {
			result[i] = -1
		} else {
			result[i] = abs(pos - query[0])
		}
	}
	return result
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
