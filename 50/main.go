package _50

func myPow(x float64, n int) float64 {
	neg := false
	if n < 0 {
		neg = true
		n = -n
	} else if n == 0 {
		return 1.0
	}

	res := x
	currentMul := x
	current := 1
	power := 1
	for power < n {
		if power+current <= n {
			res = res * currentMul
			power = power + current
			currentMul = currentMul * currentMul
			current = current * 2
		} else {
			current = 1
			currentMul = x
		}
	}
	if !neg {
		return res
	} else {
		return 1 / res
	}
}
