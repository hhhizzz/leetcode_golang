package _347

import "container/heap"

type count struct {
    value int
    times int
}

type countHeap struct {
    array []count
}

func (h *countHeap) Less(i, j int) bool {
    return h.array[i].times > h.array[j].times
}

func (h *countHeap) Swap(i, j int) {
    h.array[i], h.array[j] = h.array[j], h.array[i]
}

func (h *countHeap) Len() int {
    return len(h.array)
}

func (h *countHeap) Push(x interface{}) {
    h.array = append(h.array, x.(count))
}

func (h *countHeap) Pop() interface{} {
    last := h.array[len(h.array)-1]
    h.array = h.array[:len(h.array)-1]
    return last
}

func topKFrequent(nums []int, k int) []int {
    m := map[int]int{}
    for _, num := range nums {
        if value, ok := m[num]; ok {
            m[num] = value + 1
        } else {
            m[num] = 1
        }
    }
    h := countHeap{array: []count{}}
    heap.Init(&h)
    for k, v := range m {
        c := count{value: k, times: v}
        heap.Push(&h, c)
    }
    var result []int
    for i := 0; i < k; i++ {
        next := heap.Pop(&h).(count).value
        result = append(result, next)
    }
    return result
}
