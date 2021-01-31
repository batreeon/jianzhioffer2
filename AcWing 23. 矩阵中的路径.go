func hasPath(matrix [][]byte, str string) bool {
    if len(matrix) < 1 {
        return false
    }//防止空
    r,c := len(matrix),len(matrix[0])
    //direc := [][]int{[-1,0],[0,-1],[0,1],[1,0]}
    visited := make([]bool,r*c)//原本详见一个二维数组，但golang这块好像不太方便
    
    //istr := 0
    
    var pathRec func(pos []int,istr int, visited []bool) bool
    pathRec = func(pos []int,istr int,visited []bool) bool {//istr用来标记下面要比对第几个字符
        if istr == len(str) {
            return true
        }//已经超界了，说明找到了
        
        if pos[0] < 0 || pos[0] >= r || pos[1] < 0 || pos[1] >= c || visited[pos[0]*c+pos[1]] {
            return false
        }//要访问的节点非法或者已经访问过了
        
        if matrix[pos[0]][pos[1]] == str[istr] {
            //istr++
            visited[pos[0]*c+pos[1]] = true//当前节点匹配上了，将其即为已读，然后搜索其邻居
            if pathRec([]int{pos[0],pos[1]-1}, istr+1, visited) || 
            pathRec([]int{pos[0]-1,pos[1]}, istr+1, visited) ||
            pathRec([]int{pos[0],pos[1]+1}, istr+1, visited) ||
            pathRec([]int{pos[0]+1,pos[1]}, istr+1, visited) {
                return true
            }//找到了路径
            
            visited[pos[0]*c+pos[1]] = false//当前节点虽然匹配了，但后面走不通
        }
        return false
    }
    for i,row := range matrix {
        for j,_ := range row {
            if pathRec([]int{i,j},0,visited) {
                return true
            }
        }
    }
    return false
}