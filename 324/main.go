package _324

import "math/rand"

func findNth(nums []int, n int) int {
    pivot := 0
    for i := 1; i < len(nums); i++ {
        if nums[i] < nums[0] {
            pivot++
            nums[pivot], nums[i] = nums[i], nums[pivot]
        }
    }
    nums[0], nums[pivot] = nums[pivot], nums[0]

    if pivot == n {
        return nums[pivot]
    } else if n < pivot {
        return findNth(nums[:pivot], n)
    } else {
        return findNth(nums[pivot+1:], n-pivot-1)
    }
}

//理论上这个算法查找nth数的平均复杂度是O(n)，最差是原本有序的情况下是O(n^2)，于是可以通过把这个数组打乱的方法，来使复杂度接近O(n)
func findMedian(nums []int) int {
    for i := 0; i < len(nums); i++ {
        r := rand.Int() % len(nums)
        nums[i], nums[r] = nums[r], nums[i]
    }
    return findNth(nums, len(nums)>>1)
}

//使用虚拟位置法，通过把奇数位置映射到前半段 ，偶数位置映射到后半段，然后把大于median的数放到前半段，等于的放中间,小于的放到后半段，那么结果就是奇数位置是更小的数，偶数位置是更大的数
func wiggleSort(nums []int) {
    median := findMedian(nums)
    n := len(nums)
    left := 0
    right := n - 1

    //这里有个注意点，left一定在i左边或相等，在i左边的元素一定是小于等于median的，当这个元素从left转过来的时候不需要i--
    //反之right转过来的时候不知道这个元素和median的关系，因此还需要判断一次，就需要i--
    //还有需要注意的地方，right当前指向的位置不代表这个数一定大于median，而是下一个存储大于median数的位置，因此最后一次判断是i==right
    for i := 0; i <= right; i++ {
        if nums[newIndex(i, n)] < median {
            nums[newIndex(right, n)], nums[newIndex(i, n)] = nums[newIndex(i, n)], nums[newIndex(right, n)]
            right--
            i--
        } else if nums[newIndex(i, n)] > median {
            nums[newIndex(left, n)], nums[newIndex(i, n)] = nums[newIndex(i, n)], nums[newIndex(left, n)]
            left++
        }
    }

}

//newIndex用于找到这个数的虚拟位置，前一半的数会放到奇数位置，后一半的数会放到偶数位置，例如n为7 那么0 -> 1 1 -> 3 2 -> 5 3 -> 0 4 -> 2 5 -> 4 6 -> 6
func newIndex(index int, n int) int {
    return (2*index + 1) % (n | 1)
}
