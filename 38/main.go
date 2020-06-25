package _38

func countAndSay(n int) string {
	var res []byte
	for i := 0; i < n; i++ {
		if i == 0 {
			res = []byte{'1'}
		} else {
			var next []byte
			for j := 0; j < len(res); {
				k := j + 1
				var number byte = 1
				for k < len(res) {
					if res[k] == res[j] {
						k++
						number++
					} else {
						break
					}
				}
				next = append(next, number+'0')
				next = append(next, res[j])
				j = k
			}
			res = next
		}
	}
	return string(res)
}
