//这里填你的代码^^
//注意代码要放在两组三个点之间，才可以正确显示代码高亮哦~
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func deleteDuplication(head *ListNode) *ListNode {
//  now := head
//  for ; now != nil; {
//      next := now.Next
//      for  ; next != nil && next.Val == now.Val ; {
//          next = next.Next
//      }
//      now.Next = next
//      now = now.Next
//  }
//  return head
// }//第一遍写错了，写成了删去多余的，原来题意是把有重复的全部删掉


//好难啊，老是处理不好，看的讲解
func deleteDuplication(head *ListNode) *ListNode {
    first := new(ListNode)
    first.Next = head
    //表头指针，指向head

    //first是结果集，我们考察其Next,若合法则将其合并进来
    p := first
    for p.Next != nil {
        q := p.Next
        for q != nil && p.Next.Val == q.Val {
            q = q.Next
        }//循环退出时，要么q.Val != p.Next.Val,要么q == nil

        if p.Next.Next == q {//如果两个不同值相邻，则将前面那个加入结果中
            p = p.Next
        }else{//若不是，则忽略p.Next,考察q
            p.Next = q
        }
    }
    return first.Next
}