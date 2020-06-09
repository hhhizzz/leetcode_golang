package _225

type MyQueue struct {
	data []int
}

func (queue *MyQueue) Push(x int) {
	queue.data = append(queue.data, x)
}

func (queue *MyQueue) Pop() int {
	result := queue.data[0]
	queue.data = queue.data[1:]
	return result
}

func (queue *MyQueue) Len() int {
	return len(queue.data)
}

type MyStack struct {
	qs      []MyQueue
	current int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{[]MyQueue{MyQueue{[]int{}}, MyQueue{[]int{}}}, 0}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.qs[this.current].Push(x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	next := (this.current + 1) % 2
	for this.qs[this.current].Len() != 1 {
		this.qs[next].Push(this.qs[this.current].Pop())
	}
	result := this.qs[this.current].Pop()
	this.current = next
	return result
}

/** Get the top element. */
func (this *MyStack) Top() int {
	result := this.Pop()
	this.qs[this.current].Push(result)
	return result
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.qs[this.current].Len() == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
