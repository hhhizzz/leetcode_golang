package _66

func plusOne(digits []int) []int {
	i := len(digits) - 1
	carry := 1

	for carry != 0 && i >= 0 {
		digits[i] += carry
		carry = digits[i] / 10
		digits[i] %= 10
		i--
	}
	if carry != 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
