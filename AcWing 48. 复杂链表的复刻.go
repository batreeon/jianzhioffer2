/**
 * Definition for singly-linked list with a random pointer.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 *     Random *ListNode
 * }
 */
 func copyRandomList(head *ListNode) *ListNode {
    copyNodes := func(head *ListNode) {
        pNode := head
        for pNode != nil {
            pCopyNode := new(ListNode)
            pCopyNode.Val = pNode.Val
            pCopyNode.Next = pNode.Next
            
            pNode.Next = pCopyNode
            pNode = pCopyNode.Next
        }
    }
    connectRandoms := func(head *ListNode) {
        pNode := head
        for pNode != nil {
            pCopyNode := pNode.Next
            if pNode.Random != nil {
                pCopyNode.Random = pNode.Random.Next
            }
            pNode = pCopyNode.Next
        }
    }
    reConnectNodes := func(head *ListNode) *ListNode{
        pNode := head
        var pCopyHead,pCopyNode *ListNode = nil,nil
        if pNode != nil {
			pCopyHead,pCopyNode = pNode.Next,pNode.Next
			//pCopyHead = pCopyNode = pNode.Next 会报错
            pNode.Next = pCopyNode.Next
            pNode = pNode.Next
        }
        for pNode != nil {
            pCopyNode.Next = pNode.Next
            pCopyNode = pCopyNode.Next
            pNode.Next = pCopyNode.Next
            pNode = pNode.Next
        }
        return pCopyHead
    }
    copyNodes(head)
    connectRandoms(head)
    return reConnectNodes(head)
}