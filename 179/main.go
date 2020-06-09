package _179

import (
	"fmt"
	"strconv"
)

//if a>b
func comp(a, b string) bool {

	for i := 0; i < len(a)*len(b); i++ {
		if a[i%len(a)] > b[i%len(b)] {
			return true
		} else if a[i%len(a)] < b[i%len(b)] {
			return false
		}
	}

	return false
}

func qsort(ss []string) {
	if len(ss) <= 1 {
		return
	}
	pivot := 0
	for i := 1; i < len(ss); i++ {
		if comp(ss[i], ss[0]) {
			pivot++
			ss[i], ss[pivot] = ss[pivot], ss[i]
		}
	}
	ss[0], ss[pivot] = ss[pivot], ss[0]
	qsort(ss[:pivot])
	qsort(ss[pivot+1:])
}

func largestNumber(nums []int) string {
	var numString []string
	for _, num := range nums {
		numString = append(numString, strconv.Itoa(num))
	}
	qsort(numString)
	fmt.Println(numString)
	var result string
	for _, s := range numString {
		result += s
	}
	notZeroIndex := 0
	for notZeroIndex = 0; notZeroIndex < len(result)-1 && result[notZeroIndex] == '0'; notZeroIndex++ {
	}
	return result[notZeroIndex:]
}
