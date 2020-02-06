package _974

func subarraysDivByK(A []int, K int) int {
    result := 0
    sumMod := 0

    modCount := make([]int, K)
    modCount[0] = 1
    for i := 0; i < len(A); i++ {
        sumMod += A[i]
        sumMod %= K

        //由于golang下%不会把数变正
        for sumMod < 0 {
            sumMod += K
        }
        result += modCount[sumMod]
        modCount[sumMod]++
    }
    return result
}
