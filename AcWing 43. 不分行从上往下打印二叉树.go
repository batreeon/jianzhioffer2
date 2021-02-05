/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func printFromTopToBottom(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    node := []*TreeNode{}
    node = append(node,root)
    ans := []int{}
    for len(node) > 0 {
        l := len(node)
        for i := 0 ; i < l ; i++ {
            ans = append(ans,node[i].Val)
            if node[i].Left != nil {
                node = append(node,node[i].Left)
            }
            if node[i].Right != nil {
                node = append(node,node[i].Right)
            }
        }
        node = node[l:len(node)]
    }
    return ans
}//广度