package _349

func intersection(nums1 []int, nums2 []int) []int {
	m1 := map[int]bool{}
	m2 := map[int]bool{}
	for _, i := range nums1 {
		m1[i] = true
	}
	for _, i := range nums2 {
		m2[i] = true
	}
	mResult := map[int]bool{}
	for key := range m1 {
		if _, ok := m2[key]; ok {
			mResult[key] = true
		}
	}
	for key := range m2 {
		if _, ok := m1[key]; ok {
			mResult[key] = true
		}
	}
	var result []int
	for key := range mResult {
		result = append(result, key)
	}
	return result
}
