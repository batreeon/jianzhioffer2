/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func printFromTopToBottom(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
    node := []*TreeNode{}
    node = append(node,root)
    ans := [][]int{}
    for len(node) > 0 {
        l := len(node)//用来记录当前层的结点数
        ans = append(ans,[]int{})//这一步别忘了，新的一层
        for i := 0 ; i < l ; i++ {
            ans[len(ans)-1] = append(ans[len(ans)-1],node[i].Val)
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
}//对43题稍微改一下就行