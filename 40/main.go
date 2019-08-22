package _40

var nums []int

var result [][]int
var purpose int

func dfs(i int, sum int, current *[]int) {
    *current = append(*current, nums[i])
    sum = sum + nums[i]
    defer func() {
        *current = (*current)[:len(*current)-1]
        sum -= nums[i]
    }()

    if sum == purpose {
        new := make([]int, len(*current))
        copy(new, *current)
        result = append(result, new)
    } else if sum > purpose {
        return
    } else {
        for j := i + 1; j < len(nums); j++ {
            if j == i+1 || nums[j] != nums[j-1] {
                dfs(j, sum, current)
            }
        }
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
    nums[0], nums[pivot] = nums[pivot], nums[0]
    qsort(nums[:pivot])
    qsort(nums[pivot+1:])
}

func combinationSum2(candidates []int, target int) [][]int {
    nums = candidates
    result = make([][]int, 0)
    purpose = target

    qsort(nums)

    for i := 0; i < len(nums); i++ {
        if i == 0 || nums[i] != nums[i-1] {
            dfs(i, 0, &[]int{})
        }
    }
    return result
}
