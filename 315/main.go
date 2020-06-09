package _315

type BIT struct {
	tree []int
}

func lowBit(x int) int {
	return x & (-x)
}

func (b *BIT) getSum(index int) int {
	index += 1
	result := 0
	for index > 0 {
		result += b.tree[index]
		index -= lowBit(index)
	}
	return result
}

func (b *BIT) add(index, number int) {
	index += 1
	for index < len(b.tree) {
		b.tree[index] += number
		index += lowBit(index)
	}
}

func qsort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	pivot := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[0] {
			pivot++
			nums[pivot], nums[i] = nums[i], nums[pivot]
		}
	}
	nums[pivot], nums[0] = nums[0], nums[pivot]
	qsort(nums[:pivot])
	qsort(nums[pivot+1:])
}

func countSmaller(nums []int) []int {
	result := make([]int, len(nums))
	numsSort := make([]int, len(nums))
	copy(numsSort, nums)
	qsort(numsSort)
	ranks := make(map[int]int)
	for i := 0; i < len(numsSort); i++ {
		ranks[numsSort[i]] = i
	}
	b := BIT{tree: make([]int, len(nums)+1)}
	for i := len(nums) - 1; i >= 0; i-- {
		rank := ranks[nums[i]]
		b.add(rank, 1)
		result[i] = b.getSum(rank - 1)
	}
	return result
}
