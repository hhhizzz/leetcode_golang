package _621

import (
	"container/heap"
	"sort"
)

type task struct {
	name byte
	cool int
	left int
}

type taskHeap struct {
	array []task
}

func (t *taskHeap) Len() int {
	return len(t.array)
}

func (t *taskHeap) Less(i, j int) bool {
	if t.array[i].cool == t.array[j].cool {
		return t.array[i].left > t.array[j].left
	}
	return t.array[i].cool < t.array[j].cool
}

func (t *taskHeap) Swap(i, j int) {
	t.array[i], t.array[j] = t.array[j], t.array[i]
}

func (t *taskHeap) Push(x interface{}) {
	t.array = append(t.array, x.(task))
}

func (t *taskHeap) Pop() interface{} {
	item := t.array[len(t.array)-1]
	t.array = t.array[:len(t.array)-1]
	return item
}

func leastInterval1(tasks []byte, n int) int {
	tHeap := taskHeap{array: []task{}}
	heap.Init(&tHeap)
	m := map[byte]int{}
	for _, task := range tasks {
		m[task] += 1
	}
	for k, v := range m {
		heap.Push(&tHeap, task{k, 0, v})
	}
	time := 0
	for tHeap.Len() != 0 {
		current := heap.Pop(&tHeap).(task)
		if current.cool <= time {
			//fmt.Printf("%c -> ", current.name)
			current.cool += n + 1
			current.left -= 1
			if current.left > 0 {
				heap.Push(&tHeap, current)
			}
		} else {
			//fmt.Printf("%s -> ", "null")
			heap.Push(&tHeap, current)
		}
		time++
	}
	return time
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func leastInterval(tasks []byte, n int) int {
	m := make([]int, 26)
	for _, task := range tasks {
		m[task-'A']++
	}
	sort.Ints(m)
	maxVal := m[25] - 1
	slotLeft := maxVal * n
	for i := 24; i >= 0; i-- {
		if m[i] != 0 {
			slotLeft -= min(maxVal, m[i])
		}
	}
	if slotLeft <= 0 {
		return len(tasks)
	} else {
		return slotLeft + len(tasks)
	}
}
