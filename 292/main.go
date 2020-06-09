package _292

func canWinNim(n int) bool {
	if n == 0 {
		return true
	}
	if n%4 == 0 {
		return false
	} else {
		return true
	}
}
