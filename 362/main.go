package _362

type HitCounter struct {
    count [300]int
    term  [300]int
}

/** Initialize your data structure here. */
func Constructor() HitCounter {
    return HitCounter{[300]int{}, [300]int{}}
}

/** Record a hit.
  @param timestamp - The current timestamp (in seconds granularity). */
func (this *HitCounter) Hit(timestamp int) {
    term := timestamp / 300
    pos := timestamp % 300
    if term > this.term[pos] {
        this.term[pos] = term
        this.count[pos] = 1
    } else {
        this.count[pos]++
    }
}

/** Return the number of hits in the past 5 minutes.
  @param timestamp - The current timestamp (in seconds granularity). */
func (this *HitCounter) GetHits(timestamp int) int {
    term := timestamp / 300
    pos := timestamp % 300
    result := 0
    for i := 0; i <= pos; i++ {
        if this.term[i] == term {
            result += this.count[i]
        }
    }
    for i := pos + 1; i < 300; i++ {
        if this.term[i] == term-1 {
            result += this.count[i]
        }
    }
    return result
}
