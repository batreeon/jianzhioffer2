/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    var checkSymmetric func(root1 *TreeNode,root2 *TreeNode) bool
    checkSymmetric = func(root1 *TreeNode,root2 *TreeNode) bool {
        if root1 == nil && root2 == nil {
            return true
        }
        if (root1 == nil && root2 != nil) || (root1 != nil && root2 == nil) || root1.Val != root2.Val {
            return false
        }
        return checkSymmetric(root1.Right,root2.Left) && checkSymmetric(root1.Left,root2.Right)
    }
    
    return checkSymmetric(root.Right,root.Left)
}