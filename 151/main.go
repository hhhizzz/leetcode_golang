package _151

import (
    "strings"
)

func reverseWords(s string) string {
    splits := strings.Split(s, " ")
    var resultArray []string
    for _, split := range splits {
        if split == "" {
            continue
        }
        resultArray = append(resultArray, split)
    }
    var result string
    for i := len(resultArray) - 1; i >= 0; i-- {
        result += resultArray[i]
        if i != 0 {
            result += " "
        }
    }
    return result
}
