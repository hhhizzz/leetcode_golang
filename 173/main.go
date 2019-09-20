package _173

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

type BSTIterator struct {
    list  []int
    index int
}

func Constructor(root *TreeNode) BSTIterator {
    iterator := BSTIterator{}
    if root == nil {
        return iterator
    }
    var stack []*TreeNode
    current := root
    for true {
        for current != nil {
            stack = append(stack, current)
            current = current.Left
        }
        if len(stack) == 0 {
            break
        }
        current = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        iterator.list = append(iterator.list, current.Val)
        current = current.Right
    }
    return iterator
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
    defer func() { this.index += 1 }()
    return this.list[this.index]
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
    if this.index < len(this.list) {
        return true
    }
    return false
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
