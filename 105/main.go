package _105

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) <= 0 {
        return nil
    }
    rootVal := preorder[0]
    root := &TreeNode{rootVal, nil, nil}
    for i, num := range inorder {
        if num == rootVal {
            inOrderLeft := inorder[:i]
            inOrderRight := inorder[i+1:]
            preOrderLeft := preorder[1 : 1+len(inOrderLeft)]
            preOrderRight := preorder[1+len(inOrderLeft):]
            root.Left = buildTree(preOrderLeft, inOrderLeft)
            root.Right = buildTree(preOrderRight, inOrderRight)
        }
    }
    return root
}
