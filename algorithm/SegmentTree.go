package algorithm

type SegmentTree struct {
	tree   []int
	length int
}

func (st *SegmentTree) new(arr []int) {
	st.tree = make([]int, len(arr)*2)
	st.length = len(arr)
	for i, num := range arr {
		st.add(i, num)
	}
}

func (st *SegmentTree) add(index, number int) {
	index += st.length
	for index > 0 {
		st.tree[index] += number
		index >>= 1
	}
}

/*
   get the sum of [left,right)
*/
func (st *SegmentTree) sum(left, right int) int {
	left += st.length
	right += st.length
	result := 0
	for left < right {
		if left&1 == 1 {
			result += st.tree[left]
			left++
		}
		if right&1 == 1 {
			right--
			result += st.tree[right]
		}
		left >>= 1
		right >>= 1
	}
	return result
}
