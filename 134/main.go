package _134

func canCompleteCircuit(gas []int, cost []int) int {
	hold := make([]int, len(gas))
	for i := 0; i < len(hold); i++ {
		hold[i] = gas[i] - cost[i]
	}
	//canTry表示这个值需不需要尝试，如果一个位置尝试过不行，那么直到下一次hold为负之前都不用尝试了
	canTry := true
	for i := 0; i < len(hold); i++ {
		if hold[i] < 0 {
			canTry = true
			continue
		} else if canTry {
			current := 0
			for step := 0; step < len(hold); step++ {
				current += hold[(i+step)%len(hold)]
				if current < 0 {
					canTry = false
					break
				}
			}
			if canTry {
				return i
			}
		}
	}
	return -1
}
