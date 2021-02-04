
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func entryNodeOfLoop(head *ListNode)  *ListNode{
	nodeMap := map[*ListNode]bool{}
	ans := head
	for ; ans != nil && nodeMap[ans] != true ; {
	    nodeMap[ans] = true
	    ans = ans.Next
	}
	return ans
}