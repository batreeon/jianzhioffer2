/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func reverseList(head *ListNode) *ListNode {
//     if head == nil {
//         return nil
//     }

//     next := head.Next
//     head.Next = nil
//     last := head
//     head = next

//     for ; head != nil ; {
//         next = next.Next
//         head.Next = last
//         last = head
//         head = next
//     }

//     return last
// }
//y神的解答好清晰，惭愧

func reverseList(head *ListNode) *ListNode {
    var pre *ListNode
    cur := head
    for cur != nil {
        next := cur.Next
        cur.Next = pre
        pre,cur = cur,next
    }
    return pre
}