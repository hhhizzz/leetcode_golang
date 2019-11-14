package _432

import "container/list"

type AllOne struct {
    list.List
    data map[string]*list.Element
}

/** Initialize your data structure here. */
func Constructor() AllOne {
    return AllOne{list.List{}, map[string]*list.Element{}}
}

type keyValue struct {
    key   string
    value int
}

/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string) {
    if ele, ok := this.data[key]; ok {
        value := ele.Value.(keyValue)
        value.value += 1
        if ele != this.Back() {
            next := ele.Next()
            this.Remove(ele)
            for next != nil {
                if next.Value.(keyValue).value > value.value {
                    break
                }
                next = next.Next()
            }
            if next == nil {
                this.data[key] = this.PushBack(value)
            } else {
                this.data[key] = this.InsertBefore(value, next)
            }
        } else {
            ele.Value = value
        }
    } else {
        ele = this.PushFront(keyValue{key: key, value: 1})
        this.data[key] = ele
    }
}

/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string) {
    if ele, ok := this.data[key]; ok {
        if ele.Value.(keyValue).value == 1 {
            this.Remove(ele)
            delete(this.data, key)
        } else {
            value := ele.Value.(keyValue)
            value.value -= 1
            if ele != this.Front() {
                prev := ele.Prev()
                this.Remove(ele)
                for prev != nil {
                    if prev.Value.(keyValue).value < value.value {
                        break
                    }
                    prev = prev.Prev()
                }
                if prev == nil {
                    this.data[key] = this.PushFront(value)
                } else {
                    this.data[key] = this.InsertAfter(value, prev)
                }
            } else {
                ele.Value = value
            }
        }
    }
}

/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {
    if len(this.data) == 0 {
        return ""
    } else {
        return this.Back().Value.(keyValue).key
    }
}

/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {
    if len(this.data) == 0 {
        return ""
    } else {
        return this.Front().Value.(keyValue).key
    }
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
