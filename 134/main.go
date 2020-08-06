package _134

func canCompleteCircuit1(gas []int, cost []int) int {
	// 比较直观的写法，尝试每个点进行旅程，如果到达某个点汽油用完了，说明这段区间任意一个点出发都是不行的，可以直接从下一个点开始尝试,总复杂度O(N)
	n := len(gas)
	for i := 0; i < n; i++ {
		if gas[i]-cost[i] < 0 {
			continue
		} else {
			hold := gas[i] - cost[i]
			step := 1
			for step = 1; step < n; step++ {
				j := (i + step) % n
				hold += gas[j] - cost[j]
				if hold < 0 {
					break
				}
			}
			if step == n {
				return i
			} else {
				if i+step < n {
					i += step
				} else {
					break
				}
			}
		}
	}
	return -1
}

func canCompleteCircuit(gas []int, cost []int) int {
	// 优化版的写法，首先如果所有gas加起来比cost加起来小说明结果不存在
	// 如果这个点开始能走到结尾那么它一定能走一圈，可以反证法，首先这个点左边的点已经遍历过不能走一圈，然后这个点右边的点如果能走一起圈那么这个点也一定能(因为这个点能走到它右边的所有点)，这与题目描述的点唯一矛盾，因此这个点就是要求的点
	var hold int
	var total int
	var start int
	for i := 0; i < len(gas); i++ {
		hold += gas[i] - cost[i]
		total += gas[i] - cost[i]
		if hold < 0 {
			hold = 0
			start = i + 1
		}
	}
	if total < 0 {
		return -1
	}
	return start
}
