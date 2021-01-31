/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func printListReversingly(head *ListNode) []int {
    ans := []int{}
    var rec func(head *ListNode)
    rec = func(head *ListNode) {
        if head != nil {
            rec(head.Next)
            ans = append(ans,head.Val)
        }
    }
    rec(head)
    return ans
}