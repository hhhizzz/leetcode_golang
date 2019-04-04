package _57

type Interval struct {
    Start int
    End   int
}

func min(a int, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}

func insert(intervals []Interval, newInterval Interval) []Interval {
    var left, right []Interval
    for _, current := range intervals {
        if current.End < newInterval.Start {
            left = append(left, current)
        } else if current.Start > newInterval.End {
            right = append(right, current)
        } else {
            newInterval.Start = min(newInterval.Start, current.Start)
            newInterval.End = max(newInterval.End, current.End)
        }
    }
    return append(append(left, newInterval), right...)
}
