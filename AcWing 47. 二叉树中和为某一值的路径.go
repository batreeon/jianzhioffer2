/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func findPath(root *TreeNode, sum int) [][]int {
    path := []int{}
    ans := [][]int{}
    
    var dfs func(root *TreeNode,sum int)
    dfs = func(root *TreeNode,sum int) {
        if root == nil {
            return
        }
        path = append(path,root.Val)
        sum = sum-root.Val
        // fmt.Println(root.Val)
        // fmt.Println(path)
        if root.Right == nil && root.Left == nil && sum == 0 {
            // fmt.Println(len(ans))
            // fmt.Println(sum-root.Val)
            // fmt.Println(path)
            // fmt.Println(ans)
            tmp := make([]int,len(path))
			copy(tmp,path)//注意这两句是不能省的，跟golang语言的特性有关，省了就有错，也不能用tmp := path[:]
			//错误原因应该是共享了path,后面结果会修改前面的结果
            
            ans = append(ans,tmp)
            // fmt.Println(len(ans))
            // fmt.Println(ans)
        }
        dfs(root.Left,sum)
        dfs(root.Right,sum)
        
        path = path[:len(path)-1]
    }
    
    dfs(root,sum)
    //fmt.Println(ans)
    return ans
}