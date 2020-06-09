package _242

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m1 := make(map[int32]int)
	m2 := make(map[int32]int)
	for _, c := range s {
		if _, ok := m1[c]; ok {
			m1[c] += 1
		} else {
			m1[c] = 1
		}
	}
	for _, c := range t {
		if _, ok := m2[c]; ok {
			m2[c] += 1
		} else {
			m2[c] = 1
		}
	}
	for k, v1 := range m1 {
		if v2, ok := m2[k]; ok && v1 == v2 {
			continue
		} else {
			return false
		}
	}
	for k, v2 := range m2 {
		if v1, ok := m2[k]; ok && v1 == v2 {
			continue
		} else {
			return false
		}
	}
	return true
}
