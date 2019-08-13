package _146

import "testing"

func TestLRU(t *testing.T) {
    obj := Constructor(2)
    obj.Put(1, 1)
    obj.Put(2, 2)
    println(obj.Get(1))
    obj.Put(3, 3)
    println(obj.Get(2))
    obj.Put(4, 4)
    obj.Put(2, 2)
    println(obj.Get(1))
    println(obj.Get(3))
    println(obj.Get(4))
}
