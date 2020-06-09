package _277

/**
 * The knows API is already defined for you.
 *     knows := func(a int, b int) bool
 */
func solution(knows func(a int, b int) bool) func(n int) int {
	return func(n int) int {
		famous := 0
		//从0开始寻找潜在名人
		for i := 1; i < n; i++ {
			//如果发现这个人认识i 那说明:
			//1. 目前的潜在名人肯定不是名人
			//2. i之前的肯定都不是名人，因为他要么要么没有被潜在名人认识要么曾经存在1的情况
			//因此可以把i作为下一个潜在名人进行寻找
			if famous != i && knows(famous, i) {
				famous = i
			}
		}
		//进行double check检查是不是真正的名人
		for i := 0; i < n; i++ {
			if famous != i {
				if knows(famous, i) || !knows(i, famous) {
					return -1
				}
			}
		}
		return famous
	}
}
