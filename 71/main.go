package _71

import "strings"

func simplifyPath(path string) string {
    directories := strings.Split(path, "/")
    var result []string

    for _, dic := range directories {
        if dic == "" || dic == "." {
            continue
        } else if dic == ".." {
            if len(result) > 0 {
                result = result[:len(result)-1]
            }
        } else {
            result = append(result, dic)
        }
    }
    resultString := ""
    if len(result) == 0 {
        resultString = "/"
    } else {
        for _, r := range result {
            resultString += "/" + r
        }
    }
    return resultString
}
