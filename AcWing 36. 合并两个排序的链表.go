/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func merge(l1 *ListNode, l2 *ListNode) *ListNode {
    cur := new(ListNode)
    head := cur
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            cur.Next = l1
            cur,l1 = cur.Next,l1.Next
        }else{
            cur.Next = l2
            cur,l2 = cur.Next,l2.Next
        }
    }
    if l1 != nil {
        cur.Next = l1
    }else{
        cur.Next = l2
    }
    return head.Next
}