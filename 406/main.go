package _406

import "sort"

func reconstructQueue(people [][]int) [][]int {
    result := make([][]int, len(people))

    sort.Slice(people, func(i, j int) bool {
        if people[i][0] == people[j][0] {
            return people[i][1] < people[j][1]
        } else {
            return people[i][0] > people[j][0]
        }
    })

    for i := 0; i < len(people); i++ {
        k := people[i][1]
        if result[k] == nil {
            result[k] = people[i]
        } else {
            last, current := people[i], result[k]
            for j := k; j < len(people) && current != nil; j++ {
                current = result[j]
                result[j] = last
                last = current
            }
        }
        //fmt.Println(result)
    }
    return result

}
