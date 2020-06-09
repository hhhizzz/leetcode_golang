package algorithm

import "sync"

func BubbleSort(nums []int) {
	for i := range nums {
		for j := 1; j < len(nums)-i; j++ {
			if nums[j-1] > nums[j] {
				temp := nums[j]
				nums[j] = nums[j-1]
				nums[j-1] = temp
			}
		}
	}
}

func QuickSort(nums []int) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	quickSort(nums, &wg)
}

func quickSort(nums []int, lastWg *sync.WaitGroup) {
	defer lastWg.Done()
	if len(nums) <= 1 {
		return
	}
	left := 0
	right := len(nums) - 1
	pivotIndex := right
	pivot := nums[pivotIndex]
	for left < right {
		for left < right && nums[left] <= pivot {
			left++
		}
		for left < right && nums[right] >= pivot {
			right--
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	nums[left], nums[pivotIndex] = nums[pivotIndex], nums[left]
	wg := sync.WaitGroup{}
	wg.Add(2)
	go quickSort(nums[:left], &wg)
	go quickSort(nums[right+1:], &wg)
	wg.Wait()
}

func quickSort2(num []int, lastWg *sync.WaitGroup) {
	defer lastWg.Done()
	if len(num) <= 1 {
		return
	}
	pivotIndex := 0
	for i := 1; i < len(num); i++ {
		if num[i] < num[0] {
			pivotIndex++
			num[i], num[pivotIndex] = num[pivotIndex], num[i]
		}
	}
	num[0], num[pivotIndex] = num[pivotIndex], num[0]
	wg := sync.WaitGroup{}
	wg.Add(2)
	quickSort2(num[:pivotIndex], &wg)
	quickSort2(num[pivotIndex+1:], &wg)
}
func QuickSort2(num []int) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	quickSort2(num, &wg)
}

func mergerSort(num []int) {
	if len(num) <= 1 {
		return
	}
	mid := len(num) >> 1
	mergerSort(num[:mid])
	mergerSort(num[mid:])
	left := make([]int, mid)
	right := make([]int, len(num)-mid)
	copy(left, num[:mid])
	copy(right, num[mid:])
	li := 0
	ri := 0
	for i := 0; i < len(num); i++ {
		if li < len(left) && (ri == len(right) || left[li] <= right[ri]) {
			num[i] = left[li]
			li++
		} else {
			num[i] = right[ri]
			ri++
		}
	}
}
