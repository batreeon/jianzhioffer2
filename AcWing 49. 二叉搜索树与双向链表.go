/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func convert(root *TreeNode) *TreeNode {
    //好难好难，不会做，看的书上解答
    var convertNode func(root *TreeNode,pLastNodeInList **TreeNode)
    convertNode = func(root *TreeNode,pLastNodeInList **TreeNode) {
        if root == nil {
            return
        }
        
        pCurrent := root
        
        if pCurrent.Left != nil {
            convertNode(pCurrent.Left,pLastNodeInList)
        }
        
        pCurrent.Left =  *pLastNodeInList
        if *pLastNodeInList != nil {
            (*pLastNodeInList).Right = pCurrent
        }
        *pLastNodeInList = pCurrent
        
         if pCurrent.Right != nil {
            convertNode(pCurrent.Right,pLastNodeInList)
        }
    }
    
    var pLastNodeInList *TreeNode
    convertNode(root,&pLastNodeInList)
    
    pHeadOfList := pLastNodeInList
    for pHeadOfList != nil && pHeadOfList.Left != nil {
        pHeadOfList = pHeadOfList.Left
    }
    return pHeadOfList
}