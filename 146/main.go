package _146

type ListNode struct {
    key   int
    value int
    last  *ListNode
    next  *ListNode
}

type LRUCache struct {
    capacity int
    size     int
    index    map[int]*ListNode
    head     *ListNode
    tail     *ListNode
}

func Constructor(capacity int) LRUCache {
    return LRUCache{capacity: capacity, index: map[int]*ListNode{}}
}

func (this *LRUCache) find(key int) *ListNode {
    if current, ok := this.index[key]; ok {
        if current.key == key {
            if current == this.head {
                return current
            }
            if this.tail == current {
                this.tail = current.last
            }
            current.last.next = current.next
            if current.next != nil {
                current.next.last = current.last
            }
            current.next = this.head
            this.head.last = current
            this.head = current

            current.last = nil

            return current
        }
        current = current.next
    }
    return nil
}

func (this *LRUCache) Get(key int) int {
    result := this.find(key)
    if result != nil {
        return result.value
    }
    return -1
}

func (this *LRUCache) Put(key int, value int) {
    result := this.find(key)
    if result == nil {
        if this.head == nil {
            this.head = &ListNode{key: key, value: value}
            this.tail = this.head
            this.size = 1
        } else {
            newNode := ListNode{key: key, value: value, next: this.head}
            this.head.last = &newNode
            this.head = &newNode
            this.size++
            if this.size > this.capacity {
                delete(this.index, this.tail.key)
                if this.tail.last != nil {
                    this.tail = this.tail.last
                    this.tail.next = nil
                }
                this.size--
            }
        }
        this.index[key] = this.head
    } else {
        result.value = value
    }
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
