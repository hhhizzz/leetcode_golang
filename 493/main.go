package _493

type BIT struct {
    tree []int
}

func lowBit(num int) int {
    return num & (-num)
}

func (b *BIT) add(index int, number int) {
    index += 1
    for index > 0 {
        b.tree[index] += number
        index -= lowBit(index)
    }
}

func (b *BIT) getSum(index int) int {
    index += 1
    result := 0
    for index < len(b.tree) {
        result += b.tree[index]
        index += lowBit(index)
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

func lowerBound(nums []int, number int) int {
    left := 0
    right := len(nums)
    for left < right {
        mid := (left + right) >> 1
        if number <= nums[mid] {
            right = mid
        } else {
            left = mid + 1
        }
    }
    return left
}

func reversePairs(nums []int) int {
    numsCopy := make([]int, len(nums))
    copy(numsCopy, nums)
    qsort(numsCopy)
    b := BIT{tree: make([]int, len(nums)+1)}
    result := 0
    for i := range numsCopy {
        dual := lowerBound(numsCopy, nums[i]*2+1)
        result += b.getSum(dual)
        b.add(lowerBound(numsCopy, nums[i]), 1)
    }
    return result
}
