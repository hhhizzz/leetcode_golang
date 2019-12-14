package _87

import (
    "sort"
)

func helper(s1 []byte, s2 []byte) bool {
    if len(s1) == 1 && s1[0] == s2[0] {
        return true
    }

    s1Copy := make([]byte, len(s1))
    copy(s1Copy, s1)
    s2Copy := make([]byte, len(s2))
    copy(s2Copy, s2)

    sort.Slice(s1Copy, func(i, j int) bool {
        return s1Copy[i] < s1Copy[j]
    })
    sort.Slice(s2Copy, func(i, j int) bool {
        return s2Copy[i] < s2Copy[j]
    })
    for i := 0; i < len(s1); i++ {
        if s1Copy[i] != s2Copy[i] {
            //fmt.Printf("not equal %s,%s \n",string(s1),string(s2))
            return false
        }
    }
    for i := 1; i < len(s1); i++ {
        if helper(s1[:i], s2[:i]) && helper(s1[i:], s2[i:]) {
            //fmt.Printf("equal %s,%s \n",string(s1),string(s2))
            return true
        }
        if helper(s1[:i], s2[len(s2)-i:]) && helper(s1[i:], s2[:len(s2)-i]) {
            //fmt.Printf("equal %s,%s \n",string(s1),string(s2))
            return true
        }
    }
    //fmt.Printf("not equal %s,%s \n",string(s1),string(s2))
    return false
}

func isScramble(s1 string, s2 string) bool {
    if len(s1) != len(s2) {
        //fmt.Println("length is not equal")
        return false
    }
    array1 := []byte(s1)
    array2 := []byte(s2)

    return helper(array1, array2)
}
