package _170

type TwoSum struct {
    data map[int]int
}

/** Initialize your data structure here. */
func Constructor() TwoSum {
    return TwoSum{data: map[int]int{}}
}

/** Add the number to an internal data structure.. */
func (this *TwoSum) Add(number int) {
    this.data[number] += 1
}

/** Find if there exists any pair of numbers which sum is equal to the value. */
func (this *TwoSum) Find(value int) bool {
    for k := range this.data {
        if number, ok := this.data[value-k]; ok {
            if value-k != k || number > 1 {
                return true
            }
        }
    }
    return false
}
