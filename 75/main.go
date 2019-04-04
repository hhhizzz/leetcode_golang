package _75

func quickSort(nums []int) {
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
    quickSort(nums[:pivot])
    quickSort(nums[pivot+1:])
}

func sortColors(nums []int) {
    quickSort(nums)
}


