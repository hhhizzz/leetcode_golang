package _96

//想出来就很简单，BST就是中序遍历为1~n的树，本质上也就是问n个节点的树有多少种结构
func numTrees(n int) int {
	result := make([]int, n+1)
	result[0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			result[i] += result[j] * result[i-j-1]
		}
	}
	return result[n]
}
