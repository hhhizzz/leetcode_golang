package _70

func climbStairs(n int) int {
	tuple := []int{1, 2, 3}

	for i := 4; i <= n; i++ {
		current := (i - 1) % 3
		next := (current + 1) % 3
		nNext := (next + 1) % 3
		tuple[current] = tuple[next] + tuple[nNext]
	}
	return tuple[(n-1)%3]
}
