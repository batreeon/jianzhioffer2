//这里填你的代码^^
//注意代码要放在两组三个点之间，才可以正确显示代码高亮哦~
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 *     Father *TreeNode
 * }
 */
 func inorderSuccessor(p *TreeNode) *TreeNode {
    ans := new(TreeNode)
    if p.Right != nil {
        ans := p.Right
        for ; ans.Left != nil ; {
            ans = ans.Left
        }
        return ans
    }else if p.Father != nil {
        if p.Father.Left == p {//p没有右子树，且是其父亲的左子树
            ans = p.Father
        }else{//p没有右子树，且是其父亲的右子树
            p = p.Father
            for p.Father != nil && p.Father.Left != p {
                p = p.Father
            }
            ans = p.Father
        }
        return ans
    }
    return nil
}