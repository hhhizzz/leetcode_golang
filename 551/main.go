package _551

func checkRecord(s string) bool {
    present := 0
    late := 0
    for _, c := range s {
        if c == 'A' {
            present++
        }
        if c == 'L' {
            late++
        } else {
            late = 0
        }
        if present > 1 || late > 2 {
            return false
        }
    }
    return true
}
