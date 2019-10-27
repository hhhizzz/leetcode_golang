package _454

func fourSumCount(A []int, B []int, C []int, D []int) int {
    ABMap := map[int]int{}
    CDMap := map[int]int{}
    result := 0

    for _, numA := range A {
        for _, numB := range B {
            sum := numA + numB
            ABMap[sum] += 1
        }
    }

    for _, numC := range C {
        for _, numD := range D {
            sum := numC + numD
            CDMap[sum] += 1
        }
    }

    for k, v := range ABMap {
        if v2, ok := CDMap[-k]; ok {
            result += v * v2
        }
    }

    return result
}
