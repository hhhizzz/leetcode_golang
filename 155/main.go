package _155

type MinStack struct {
    stack []int
    min   []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{stack: []int{}, min: []int{}}
}

func (this *MinStack) Push(x int) {
    this.stack = append(this.stack, x)
    if len(this.min) == 0 || x < this.min[len(this.min)-1] {
        this.min = append(this.min, x)
    } else {
        this.min = append(this.min, this.min[len(this.min)-1])
    }
}

func (this *MinStack) Pop() {
    this.stack = this.stack[:len(this.stack)-1]
    this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
    return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
    return this.min[len(this.min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
