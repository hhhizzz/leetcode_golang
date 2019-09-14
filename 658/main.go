package _658

func abs(num int) int {
    if num < 0 {
        return -num
    }
    return num
}

// 方法是通过二分找区间的起始位置
func findClosestElements(arr []int, k int, x int) []int {
    left := 0
    right := len(arr) - k
    for left < right {
        mid := left + ((right - left) >> 1)
        if abs(arr[mid]-x) > abs(arr[mid+k]-x) {
            //如果mid的位置下距离x的距离比mid+k远，说明可以继续往右搜索，因为只要如果往右再挪到一格距离那肯定也是减少了的
            left = mid + 1
        } else {
            right = mid
        }
    }
    return arr[left : left+k]
}
