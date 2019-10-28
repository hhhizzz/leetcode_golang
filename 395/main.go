package _395

import "sort"

func helper(array []byte, k int) int {
    if len(array) < k {
        return 0
    }
    m := map[byte][]int{}
    for index, c := range array {
        m[c] = append(m[c], index)
    }
    var posList []int
    for _, pos := range m {
        if len(pos) < k {
            posList = append(posList, pos...)
        }
    }
    if len(posList) == 0 {
        return len(array)
    }
    sort.Ints(posList)
    result := helper(array[0:posList[0]], k)
    for i := 0; i < len(posList); i++ {
        var current int
        if i != len(posList)-1 {
            current = helper(array[posList[i]+1:posList[i+1]], k)
        } else {
            current = helper(array[posList[i]+1:], k)
        }
        if current > result {
            result = current
        }
    }
    return result
}

func longestSubstring(s string, k int) int {
    array := []byte(s)
    return helper(array, k)
}
