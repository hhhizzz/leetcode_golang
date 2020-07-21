package _96

// 方法一：用于理解，复杂度过不了，节点数为n，从1到n每个数都可以做根节点，当i作为根节点的时候，左边有i-1个数作为节点形成子树，右边同理n-i个数
// 因此将问题拆解，左边子树的个数乘以右子树的个数即得到结果，利用递归很容易写出
func numTrees1(n int) int {
	if n <= 1 {
		return 1
	}
	var result int
	for i := 1; i <= n; i++ {
		result += numTrees(i-1) * numTrees(n-i)
	}
	return result
}

// 递归的方式容易理解，但是很容易发现有大量的重复计算，因此可以用动态规划将其重构
func numTrees(n int) int {
	dp := make([]int, 20)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
