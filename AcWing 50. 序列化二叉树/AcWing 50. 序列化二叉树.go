/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 /**
 * Encodes a tree to a single string.
 */
 func serialize(root *TreeNode) string {
    if root == nil {
        return string("[null]")
    }
    /*
    var ansBuffer bytes.Buffer
    ansBuffer.WriteString("[")
    node := []*TreeNode{}
    node = append(node,root)
    
    ansBuffer.WriteString(strconv.Itoa(root.Val))
    
    l := len(node)
    for i := 0 ; i < l ; i++ {
        node = append(node,node[i].Left)
        node = append(node,node[i].Right)
    }
    node = node[l:len(node)]
    
    for len(node) < 0 {
        l := len(node)
        for i := 0 ; i < l ; i++ {
            ansBuffer.WriteString(", ")
            if node[i] != nil {
                ansBuffer.WriteString(strconv.Itoa(node[i].Val))
                node = append(node,node[i].Left)
                node = append(node,node[i].Right)
            }else{
                ansBuffer.WriteString(string("nil"))
            }
        }
    }
    ansBuffer.WriteString(string("]"))
    return ansBuffer.String()
    实现结果和下面一样，但有点麻烦
    */
    ansStringArray := []string{}
    node := []*TreeNode{}
    node = append(node,root)

    for len(node) > 0 {
        l := len(node)
        for i := 0 ; i < l ; i++ {
            if node[i] != nil {
                ansStringArray = append(ansStringArray,strconv.Itoa(node[i].Val))
                node = append(node,node[i].Left)
                node = append(node,node[i].Right)
            }else{
                ansStringArray = append(ansStringArray,string("null"))
            }
        }
        node = node[l:len(node)]
    }
    return strings.Join([]string{"[",strings.Join(ansStringArray, ", "),"]"},"")
}
/**
 * Decodes your encoded data to tree.
 */
func deserialize(data string) *TreeNode {
    if data[1:5] == string("null") {
        return nil
    }
    dataPurity := data[1:len(data)-1]   //去掉开头的[和末尾的]
    dataStringArray := strings.Split(dataPurity, ", ")
    
    node := []*TreeNode{}
    rootNode := new(TreeNode)
    rootNode.Val,_ = strconv.Atoi(dataStringArray[0])
    node = append(node,rootNode)
    dataStringArray = dataStringArray[1:]
    
    for len(node) > 0 {
        l := len(node)
        for i := 0 ; i < l ; i++ {
            if node[i] != nil {
                var pLeftNode,pRightNode *TreeNode
                if dataStringArray[0] != string("null") {
                    pLeftNode = new(TreeNode)
                    pLeftNode.Val,_ = strconv.Atoi(dataStringArray[0])
                }
                
                if dataStringArray[1] != string("null") {
                    pRightNode = new(TreeNode)
                    pRightNode.Val,_ = strconv.Atoi(dataStringArray[1])
                }
                node[i].Left,node[i].Right = pLeftNode,pRightNode
                node = append(node,pLeftNode)
                node = append(node,pRightNode)
                dataStringArray = dataStringArray[2:]
            }
        }
        node = node[l:]
    }
    return rootNode
}