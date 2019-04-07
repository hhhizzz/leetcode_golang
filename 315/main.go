package _315

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

type segTree struct {
    tree [] int
    n    int
}

func (st *segTree) add(index int, number int) {
    index += st.n
    for index > 0 {
        st.tree[index] += number
        index = index >> 1
    }
}

func (st *segTree) sum(left int, right int) int {
    left += st.n
    right += st.n
    sum := 0
    for left < right {
        if left&1 == 1 {
            sum += st.tree[left]
            left++
        }
        if right&1 == 1 {
            right--
            sum += st.tree[right]
        }
        left = left >> 1
        right = right >> 1
    }
    return sum
}

func countSmaller(nums []int) []int {
    sorted := make([]int, len(nums))
    copy(sorted, nums)
    qsort(sorted)
    ranks := make(map[int]int, len(sorted))
    for i := range sorted {
        ranks[sorted[i]] = i
    }
    for i := range nums {
        nums[i] = ranks[nums[i]]
    }
    st := segTree{tree: make([]int, len(nums)*2), n: len(nums)}
    result := make([]int, len(nums))
    for i := len(nums) - 1; i >= 0; i-- {
        st.add(nums[i], 1)
        result[i] = st.sum(0, nums[i])
    }
    return result
}
