package _255

import (
    "math"
    "sort"
)

func check(pre, in []int) bool {
    if len(pre) != len(in) {
        return false
    }
    if len(pre) == 0 {
        return true
    } else if len(pre) == 1 {
        return pre[0] == in[0]
    } else {
        root := pre[0]
        find := -1
        for i := 0; i < len(in); i++ {
            if in[i] == root {
                find = i
                break
            }
        }
        if find == -1 {
            return false
        }
        return check(pre[1:1+find], in[:find]) && check(pre[1+find:], in[find+1:])
    }
}

//使用中序递归检查，直接能想到的方法
func verifyPreorder1(preorder []int) bool {
    inorder := make([]int, len(preorder))
    copy(inorder, preorder)
    sort.Ints(inorder)
    return check(preorder, inorder)
}

//非递归方案，O(N*Log(N)),内存使用O(Log(N)) 结果比递归慢了一倍
func verifyPreorder2(preorder []int) bool {
    queue := [][]int{preorder}

    for len(queue) != 0 {
        current := queue[0]
        queue = queue[1:]
        if len(current) == 0 {
            continue
        }
        i := 1
        for ; i < len(current); i++ {
            if current[i] > current[0] {
                break
            }
        }
        queue = append(queue, current[1:i])

        for j := i; j < len(current); j++ {
            if current[j] < current[0] {
                return false
            }
        }
        queue = append(queue, current[i:])
    }
    return true
}

//单调栈方案，单调递减栈，每次入栈表示在某个左子树上，大于栈顶表示这里开始进入某个右子树，而这个右子树的根节点就是最后出栈的那个元素，因为这是preOrder左边比他小的最大元素，此时只要判断剩下的节点都比这个根节点大即可
func verifyPreorder(preorder []int) bool {
    var stack []int
    root := math.MinInt32

    for _, ele := range preorder {
        if ele < root {
            return false
        }
        if len(stack) == 0 || ele < stack[len(stack)-1] {
            stack = append(stack, ele)
        } else {
            for len(stack) > 0 && ele > stack[len(stack)-1] {
                root = stack[len(stack)-1]
                stack = stack[:len(stack)-1]
            }
            stack = append(stack, ele)
        }
    }
    return true
}
