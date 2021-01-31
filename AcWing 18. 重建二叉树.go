/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }
    //貌似只有最外层才会用到这个判断，后面都保证了preorder不为空，
    //但是由于root的左右子树默认初始为nil,因此得到的结果是对的

    root := new(TreeNode)
    root.Val = preorder[0]
    lenLeft := 0
    for ; inorder[lenLeft] != root.Val; lenLeft++ {}
    //找到根节点在inorder中的下标，这个也等于他前面的元素个数，它前面的节点都在该根的左子树里
    //inorder里，根节点后面的都在其右子树里

    if lenLeft > 0 {
        root.Left = buildTree(preorder[1:1+lenLeft],inorder[:lenLeft])
    }
    if len(preorder)-lenLeft-1 > 0 {
        root.Right = buildTree(preorder[lenLeft+1:],inorder[lenLeft+1:])
    }
    return root
}