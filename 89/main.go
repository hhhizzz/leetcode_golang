package _89

func grayCode(n int) []int {
	m := map[int]bool{0: true}
	res := []int{0}
	var mask []int

	num := 1
	for i := 0; i < n; i++ {
		mask = append(mask, 1<<i)
		num = num << 1
	}
	for i := 0; i < num; i++ {
		for _, ma := range mask {
			next := res[len(res)-1] ^ ma
			if _, ok := m[next]; !ok {
				res = append(res, next)
				m[next] = true
				break
			}
		}
	}
	return res
}
