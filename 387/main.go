package _387

func firstUniqChar(s string) int {
	m := map[int32]int{}
	for _, c := range s {
		if _, ok := m[c]; ok {
			m[c] += 1
		} else {
			m[c] = 1
		}
	}
	for i, c := range s {
		if m[c] == 1 {
			return i
		}
	}
	return -1
}
