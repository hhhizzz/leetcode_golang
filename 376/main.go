package _376

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//首先想到的方法，类似股票问题进行dp，复杂度O(n^2)
func wiggleMaxLength1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	//dp[i][0]以i为摆动序列最后一个元素且是向下的最长长度 dp[i][1]同理表示向上
	dp := make([][2]int, len(nums))
	dp[0][0] = 1
	dp[0][1] = 1
	result := 1
	for i := 1; i < len(nums); i++ {
		dp[i][0] = 1
		dp[i][1] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[j] > nums[i] {
				dp[i][0] = max(dp[i][0], dp[j][1]+1)
				result = max(result, dp[i][0])
			} else if nums[j] < nums[i] {
				dp[i][1] = max(dp[i][1], dp[j][0]+1)
				result = max(result, dp[i][1])
			}
		}
	}
	return result
}

//直接贪心，注意处理等于的情况
func wiggleMaxLength(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	count := 1
	prevDiff := nums[1] - nums[0]
	if prevDiff != 0 {
		count++
	}
	for i := 2; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		//注意这里可以取等号，但是只有发生了diff和preDiff的情况才会升级preDiff，保证之后的preDiff不会为0
		if prevDiff <= 0 && diff > 0 {
			count++
			prevDiff = diff
		} else if prevDiff >= 0 && diff < 0 {
			count++
			prevDiff = diff
		}
	}
	return count
}
