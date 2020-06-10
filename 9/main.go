package _9

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	tmp := x
	res := 0
	for x > 0 {
		res *= 10
		res += x % 10
		x /= 10
	}
	return res == tmp
}
