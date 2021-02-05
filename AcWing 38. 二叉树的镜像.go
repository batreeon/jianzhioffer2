//这里填你的代码^^
//注意代码要放在两组三个点之间，才可以正确显示代码高亮哦~
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func mirror(root *TreeNode)  {
    if root == nil {
        return 
    }
    root.Left,root.Right = root.Right,root.Left
    mirror(root.Left)
    mirror(root.Right)
}