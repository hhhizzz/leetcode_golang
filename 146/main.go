package _146

import (
    "container/list"
    "fmt"
)

type LRUCache struct {
    data     map[int]*list.Element
    l        *list.List
    capacity int
    length   int
}

type keyValue struct {
    key   int
    value int
}

func Constructor(capacity int) LRUCache {
    return LRUCache{data: map[int]*list.Element{}, l: list.New(), capacity: capacity, length: 0}
}

func (this *LRUCache) Get(key int) int {
    if ele, ok := this.data[key]; ok {
        kv := ele.Value.(keyValue)
        if ele != this.l.Front() {
            this.l.MoveToFront(ele)
        }
        return kv.value
    }
    return -1
}

func (this *LRUCache) Put(key int, value int) {
    if ele, ok := this.data[key]; ok {
        ele.Value = keyValue{key: key, value: value}
        if ele != this.l.Front() {
            this.l.MoveToFront(ele)
        }
    } else {
        if this.length >= this.capacity {
            ele := this.l.Back()
            this.l.Remove(ele)
            fmt.Printf("envicts: %d\n", ele.Value.(keyValue).value)
            delete(this.data, ele.Value.(keyValue).key)
            this.length--
        }
        ele := this.l.PushFront(keyValue{key: key, value: value})
        this.data[key] = ele
        this.length++
    }
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
