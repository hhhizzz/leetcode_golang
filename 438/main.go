package _438

func findAnagrams(s string, p string) []int {
	needed := map[byte]int{}
	neededLen := 0
	for _, c := range p {
		needed[byte(c)] += 1
		if needed[byte(c)] == 1 {
			neededLen += 1
		}
	}
	left := 0
	right := 0
	var result []int
	array := []byte(s)
	window := map[byte]int{}
	matchLen := 0
	for right < len(s) {
		current := array[right]
		window[current] += 1
		if needed[current] == 0 {
			right++
			left = right
			window = map[byte]int{}
			matchLen = 0
			continue
		}
		if window[current] == needed[current] {
			matchLen++
		}
		//fmt.Printf("left %d right %d window %v matchLen %d\n",left,right,window,matchLen)
		for matchLen == neededLen {
			if right-left+1 == len(p) {
				result = append(result, left)
			}
			current = array[left]
			window[current]--
			left++
			if window[current] < needed[current] {
				matchLen--
				break
			}
		}
		right++
	}
	return result
}
