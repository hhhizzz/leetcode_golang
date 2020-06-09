package _88

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := m - 1; i >= 0; i-- {
		nums1[i+n] = nums1[i]
	}
	l := n
	r := 0
	i := 0
	for ; l < m+n && r < n; i++ {
		if nums1[l] < nums2[r] {
			nums1[i] = nums1[l]
			l++
		} else {
			nums1[i] = nums2[r]
			r++
		}
	}
	if l == m+n {
		for i < m+n {
			nums1[i] = nums2[r]
			r++
			i++
		}
	} else {
		for i < m+n {
			nums1[i] = nums1[l]
			l++
			i++
		}
	}
}
