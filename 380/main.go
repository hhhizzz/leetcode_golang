package _380

import "math/rand"

type RandomizedSet struct {
    m    map[int]int
    list []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
    return RandomizedSet{m: map[int]int{}, list: []int{}}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
    if _, ok := this.m[val]; ok {
        return false
    } else {
        this.list = append(this.list, val)
        this.m[val] = len(this.list) - 1
        return true
    }
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
    if pos, ok := this.m[val]; ok {
        delete(this.m, val)
        if len(this.list) == 1 {
            this.list = []int{}
        } else {
            lastItem := this.list[len(this.list)-1]
            if pos == len(this.list)-1 {
                this.list = this.list[:len(this.list)-1]
            } else {
                this.list[pos] = lastItem
                this.m[lastItem] = pos
                this.list = this.list[:len(this.list)-1]
            }
        }
        return true
    } else {
        return false
    }
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
    next := rand.Int() % len(this.list)
    return this.list[next]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
