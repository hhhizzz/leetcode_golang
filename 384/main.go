package _384

import "math/rand"

type Solution struct {
    origin []int
}

func Constructor(nums []int) Solution {
    return Solution{origin: nums}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
    return this.origin
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
    result := make([]int, len(this.origin))
    copy(result, this.origin)
    for i := 0; i < len(result); i++ {
        randomIndex := rand.Int() % len(result)
        result[i], result[randomIndex] = result[randomIndex], result[i]
    }
    return result
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
