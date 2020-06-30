package _32

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
	需要对括号串的了解比较深，算法写起来不难但很难想到
	一个符合要求的子串可以判断两个条件，一个很容易想到的是左右括号数量相等，第二个是它的任意前缀都是左括号数大于等于右括号的
	于是可以可以使用栈来辅助查找这样的子串
*/
func longestValidParentheses(s string) int {
	var stack []int
	res := 0
	// split作为上一个不符合要求的子串的末尾，这个末尾一定是个右括号，从0开始计数，因此初始值可以设置为-1
	split := -1
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			// 左括号入栈操作，记录括号位置
			stack = append(stack, i)
		} else {
			//右括号出栈操作
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
				var current int
				if len(stack) > 0 {
					// 出栈后，如果还剩有左括号 类似(()的情况，通过栈顶元素来计算最长符合要求的后缀的长度
					current = i - stack[len(stack)-1]
				} else {
					// 如果不剩左括号，说明当前子串就是符合要求的子串，通过split来计算当前子串长度
					current = i - split
				}
				res = max(res, current)
			} else {
				// 如果栈空，说明右括号数在目前位置大于了左括号数，该子串不合要求，尝试下一次查找，重设split值
				split = i
			}
		}
	}
	return res
}
