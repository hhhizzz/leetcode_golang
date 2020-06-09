package _218

import "sort"

type heights []int

func (h *heights) find(height int) int {
	left := 0
	right := len(*h)
	for left < right {
		mid := left + (right-left)>>1
		if (*h)[mid] == height {
			return mid
		} else if (*h)[mid] < height {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func (h *heights) add(height int) {
	find := h.find(height)
	*h = append((*h)[:find], append([]int{height}, (*h)[find:]...)...)
}

func (h *heights) erase(height int) {
	find := h.find(height)
	*h = append((*h)[:find], (*h)[find+1:]...)
}

func getSkyline(buildings [][]int) [][]int {
	var all [][]int
	for _, building := range buildings {
		//正数代表左端点，负数代表右端点
		all = append(all, []int{building[0], building[2]})
		all = append(all, []int{building[1], -building[2]})
	}
	sort.Slice(all, func(i, j int) bool {
		if all[i][0] == all[j][0] {
			return all[i][1] > all[j][1]
		} else {
			return all[i][0] < all[j][0]
		}
	})
	h := heights{}
	lastHeight := []int{0, 0}
	var result [][]int
	for _, building := range all {
		if building[1] > 0 {
			h.add(building[1])
		} else {
			h.erase(-building[1])
		}
		var maxHeight int
		if len(h) > 0 {
			maxHeight = h[len(h)-1]
		} else {
			maxHeight = 0
		}
		if maxHeight != lastHeight[1] {
			lastHeight[0] = building[0]
			lastHeight[1] = maxHeight
			result = append(result, []int{lastHeight[0], lastHeight[1]})
		}

	}
	return result
}
