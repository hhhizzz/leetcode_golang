package _106

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
    if len(inorder) <= 0 {
        return nil
    }
    rootVal := postorder[len(postorder)-1]
    root := &TreeNode{rootVal, nil, nil}
    for i, num := range inorder {
        if num == rootVal {
            inLeftPart := inorder[:i]
            inRightPart := inorder[i+1:]
            postLeftPart := postorder[:len(inLeftPart)]
            postRightPart := postorder[len(inLeftPart) : len(postorder)-1]
            root.Left = buildTree(inLeftPart, postLeftPart)
            root.Right = buildTree(inRightPart, postRightPart)
        }
    }
    return root
}
