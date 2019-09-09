package _552

//注意1e9表示10^9，10e9表示10*10^9=10^10
const M int64 = 1e9 + 7

func checkRecord(n int) int {
    //表示最新一天是这些的情况
    var present int64 = 1
    var presentWithAbsent int64 = 0
    var late int64 = 1
    var lateWithAbsent int64 = 0
    var lateLate int64 = 0
    var lateLateWithAbsent int64 = 0
    var absent int64 = 1
    for i := 2; i <= n; i++ {
        newPresent := (present + late + lateLate) % M
        newLate := present % M
        newLateLate := late % M

        newPresentWithAbsent := (presentWithAbsent + lateWithAbsent + lateLateWithAbsent + absent) % M
        newLateWithAbsent := (presentWithAbsent + absent) % M
        newLateLateWithAbsent := lateWithAbsent
        newAbsent := (present + late + lateLate) % M

        present = newPresent
        presentWithAbsent = newPresentWithAbsent
        late = newLate
        lateWithAbsent = newLateWithAbsent
        lateLate = newLateLate
        lateLateWithAbsent = newLateLateWithAbsent
        absent = newAbsent
    }
    return int((present + presentWithAbsent + late + lateWithAbsent + lateLate + lateLateWithAbsent + absent) % M)
}
