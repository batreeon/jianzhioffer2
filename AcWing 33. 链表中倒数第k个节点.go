/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func findKthToTail(pListHead *ListNode, k int) *ListNode {
    listLen := func(head *ListNode) int {
        l := 0
        for head != nil {
            l++
            head = head.Next
        }
        return l
    }
    
    l := listLen(pListHead)//这个函数居然没有改变这个指针值
    if k > l {
        return nil
    }
    l = l - k
    for ; l > 0 ; l-- {
        pListHead = pListHead.Next
    } 
    
    return pListHead
}