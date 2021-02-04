/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func hasSubtree(pRoot1 *TreeNode, pRoot2 *TreeNode) bool {

    var checkT2inT1 func(pRoot1 *TreeNode, pRoot2 *TreeNode) bool
    checkT2inT1 = func(pRoot1 *TreeNode, pRoot2 *TreeNode) bool {
        if pRoot2 == nil {
            return true
        }
        if pRoot1 == nil {
            return false
        }
        if pRoot1.Val != pRoot2.Val {
            return false
        }//上面是checkT2inT1的递归退出条件
        
        return checkT2inT1(pRoot1.Left,pRoot2.Left) && checkT2inT1(pRoot1.Right,pRoot2.Right)
    }
    
    if pRoot1 != nil && pRoot2 != nil {
        //这一条判断，一是防止访问空指针，而是不满足这个条件当然也不会是正确答案。
        //其实也可在函数最开头判断，若有一个rroot为nil，则直接返回f
        if pRoot1.Val == pRoot2.Val {
            return checkT2inT1(pRoot1,pRoot2) || hasSubtree(pRoot1.Left,pRoot2) || hasSubtree(pRoot1.Right,pRoot2)
        }else{
            return hasSubtree(pRoot1.Left,pRoot2) || hasSubtree(pRoot1.Right,pRoot2)
        }
    }
    return false
}