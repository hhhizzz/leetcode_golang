package _60

import "strconv"

var factorial []int

func helper(array []int, k int) string {
    var result string
    if len(array) == 0 {
        return result
    }
    if k <= 0 {
        for i := range array {
            result += strconv.Itoa(array[i])
        }
        return result
    }
    times := k / factorial[len(array)-1]
    nextK := k % factorial[len(array)-1]

    result += strconv.Itoa(array[times])
    array = append(array[:times], array[times+1:]...)
    return result + helper(array, nextK)
}

func getPermutation(n int, k int) string {
    var array []int

    factorial = make([]int, n+1)
    array = make([]int, n)
    factorial[0] = 1

    for i := 1; i <= n; i++ {
        factorial[i] = factorial[i-1] * i
    }
    for i := range array {
        array[i] = i + 1
    }
    return helper(array, k-1)
}
