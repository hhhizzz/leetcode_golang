package _88

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := m - 1; i >= 0; i-- {
		nums1[i+n] = nums1[i]
	}
	for i, j, pos := n, 0, 0; i < n+m || j < n; pos++ {
		if j == n || (i != n+m && nums1[i] < nums2[j]) {
			nums1[pos] = nums1[i]
			i++
		} else {
			nums1[pos] = nums2[j]
			j++
		}
	}
}
