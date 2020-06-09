package algorithm

type BIT struct {
	tree []int
}

func lowBit(x int) int {
	return x & (-x)
}

func (b *BIT) new(arr []int) {
	b.tree = make([]int, len(arr)+1)
	for i, number := range arr {
		b.update(i, number)
	}
}

func (b *BIT) update(index int, number int) {
	index += 1
	for index < len(b.tree) {
		b.tree[index] += number
		index += lowBit(index)
	}
}

func (b *BIT) sumPre(index int) int {
	index += 1
	result := 0
	for index >= 0 {
		result += b.tree[index]
		index -= lowBit(index)
	}
	return result
}
