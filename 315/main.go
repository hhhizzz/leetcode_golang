package _315

type bit struct {
    tree []int
}

func lowBit(index int) int {
    return index & (-index)
}

func (b *bit) add(index int, number int) {
    index += 1
    for index < len(b.tree) {
        b.tree[index] += number
        index += lowBit(index)
    }
}

func (b *bit) preSum(index int) int {
    index += 1
    sum := 0
    for index > 0 {
        sum += b.tree[index]
        index -= lowBit(index)
    }
    return sum
}

func countSmaller(nums []int) []int {
    sorted := make([]int, len(nums))
    copy(sorted, nums)
    qsort(sorted)
    rank := 0
    ranks := make(map[int]int)
    for _, s := range sorted {
        ranks[s] = rank
        rank++
    }
    for i, n := range nums {
        nums[i] = ranks[n]
    }
    result := make([]int, len(nums))
    var bt = bit{tree: make([]int, len(nums)+1)}
    for i := len(nums) - 1; i >= 0; i-- {
        bt.add(nums[i], 1)
        result[i] += bt.preSum(nums[i] - 1)

    }
    return result
}

func qsort(nums []int) {
    if len(nums) <= 1 {
        return
    }
    pivot := 0
    for i := 1; i < len(nums); i++ {
        if nums[i] < nums[0] {
            pivot++
            nums[i], nums[pivot] = nums[pivot], nums[i]
        }
    }
    nums[0], nums[pivot] = nums[pivot], nums[0]
    qsort(nums[:pivot])
    qsort(nums[pivot+1:])
}
