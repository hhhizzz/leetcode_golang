package _76

import "math"

func minWindow(s string, t string) string {
	needs := map[byte]int{}
	needsLen := 0
	for _, c := range t {
		needs[byte(c)] += 1
		if needs[byte(c)] == 1 {
			needsLen++
		}
	}
	minLen := math.MaxInt32
	var result []byte
	left := 0
	right := 0
	window := map[byte]int{}
	array := []byte(s)
	match := 0
	for right < len(array) {
		current := array[right]
		window[current]++
		if times, ok := needs[current]; ok {
			if window[current] == times {
				match++
			}
		}
		if match == needsLen {
			currentLen := right - left + 1
			if currentLen < minLen {
				minLen = currentLen
				result = array[left : right+1]
			}
			for left < right {
				current = array[left]
				window[current]--
				left++
				if window[current] < needs[current] {
					match--
					break
				}
				currentLen = right - left + 1
				if currentLen < minLen {
					minLen = currentLen
					result = array[left : right+1]
				}
			}
		}
		right++
	}
	return string(result)
}
