/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func printFromTopToBottom(root *TreeNode) [][]int {
    //和分层打印二叉树差不多，只需要把相应层的读取结果反转一下就可以了
    if root == nil {
        return nil
    }
    node := []*TreeNode{}
    node = append(node,root)
    ans := [][]int{}
    depth := 0//记录高度，用来指示横向读取的方向
    reverse := func(s []int) {
        for i,j := 0,len(s)-1 ; i < j ; i,j = i+1,j-1 {
            s[i],s[j] = s[j],s[i]
        }
    }
    for len(node) > 0 {
        l := len(node)//因为在读取过程中会加入下一层的节点，我们需要先记录上一层节点个数，读一层就够了
        ans = append(ans,[]int{})
        for i := 0 ; i < l ; i++ {
            ans[len(ans)-1] = append(ans[len(ans)-1],node[i].Val)
            if node[i].Left != nil {
                node = append(node,node[i].Left)
            }
            if node[i].Right != nil {
                node = append(node,node[i].Right)
            }
        }
        node = node[l:]//下一层的节点
        if depth%2 == 1 {//从右到左
            reverse(ans[len(ans)-1])
        }
        depth++//下一层
    }
    return ans
}