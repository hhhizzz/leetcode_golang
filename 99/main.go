package _99

import "sort"

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

//方法一：直接中序
func inorderAppend(root *TreeNode, array *[]int) {
    if root == nil {
        return
    }
    inorderAppend(root.Left, array)
    *array = append(*array, root.Val)
    inorderAppend(root.Right, array)
}

func inorderFill(root *TreeNode, array []int, current *int) {
    if root == nil {
        return
    }
    inorderFill(root.Left, array, current)
    root.Val = array[*current]
    *current++
    inorderFill(root.Right, array, current)
}

func recoverTree1(root *TreeNode) {
    var array []int
    inorderAppend(root, &array)
    sort.Ints(array)
    iterator := 0
    inorderFill(root, array, &iterator)
}

//方法二：Morris方法中序遍历+冒泡
func recoverTree(root *TreeNode) {
    current := root
    var prevent *TreeNode
    var firstNode *TreeNode
    var secondNode *TreeNode
    for current != nil {
        //如果左子树为空那么说明这个节点是某个子树最左边的节点，可以遍历这个节点
        if current.Left == nil {
            //fmt.Println(prevent,current.Val)
            if prevent != nil && prevent.Val > current.Val {
                if firstNode == nil {
                    firstNode = prevent
                }
                secondNode = current
            }
            prevent = current
            //同时进入它的右子树
            current = current.Right
        } else {
            //不为空说明还可以向左进行，首先找到current左子树最右边的节点leftRight的右子树接到current上
            leftRight := current.Left
            for leftRight.Right != nil && leftRight.Right != current {
                leftRight = leftRight.Right
            }
            //如果左子树没有被遍历过,就把pre的右节点指向current,并把current向左迈一步
            if leftRight.Right != current {
                leftRight.Right = current
                current = current.Left
            } else {
                //如果遍历过，就把这个pre的指向去掉，恢复原来树结构，并让current进入右子树
                leftRight.Right = nil
                //fmt.Println(prevent,current.Val)
                if prevent != nil && prevent.Val > current.Val {
                    if firstNode == nil {
                        firstNode = prevent
                    }
                    secondNode = current
                }
                prevent = current
                current = current.Right
            }
        }
    }
    firstNode.Val, secondNode.Val = secondNode.Val, firstNode.Val
}
