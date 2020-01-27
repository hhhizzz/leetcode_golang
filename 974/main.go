package _974

func subarraysDivByK(A []int, K int) int {
    var P int
    m := map[int]int{0: 1}
    result := 0
    for i := 0; i < len(A); i++ {
        if i == 0 {
            P = A[i] % K
        } else {
            P = (P + A[i]) % K
        }
        //注意golang里%不会取正数
        for P < 0 {
            P += K
        }
        if num, ok := m[P]; ok {
            result += num
            m[P] = num + 1
        } else {
            m[P] = 1
        }
    }
    //fmt.Println(m)
    return result
}
