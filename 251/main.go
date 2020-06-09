package _251

type Vector2D struct {
	next     int
	capacity int
	data     [][]int
}

func Constructor(v [][]int) Vector2D {
	capacity := 0
	for i := 0; i < len(v); i++ {
		capacity += len(v[i])
	}
	return Vector2D{0, capacity, v}
}

func (this *Vector2D) Next() int {
	current := this.next
	for i := 0; i < len(this.data); i++ {
		if current < len(this.data[i]) {
			this.next++
			return this.data[i][current]
		}
		current -= len(this.data[i])
	}
	return -1
}

func (this *Vector2D) HasNext() bool {
	return this.next < this.capacity
}

/**
 * Your Vector2D object will be instantiated and called as such:
 * obj := Constructor(v);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
