func printMatrix(matrix [][]int) []int {
    if len(matrix) == 0 {
        return []int{}
    }   //nil
    if len(matrix[0]) == 0 {
        return []int{}
    }   //[[],[]]
    r,c := len(matrix),len(matrix[0])
    ans := []int{}
    corner := 0
    cornerPos := make([][]int,4)    //标记四个角的坐标，我是通过草稿观察，然后总结角坐标变化规律的
    cornerPos[0] = []int{0,0}
    cornerPos[1] = []int{0,c-1}
    cornerPos[2] = []int{r-1,c-1}
    cornerPos[3] = []int{r-1,0}
    rr,cc := cornerPos[0][0],cornerPos[0][1]
    
    for i := 0 ; i < r*c - 1 ; {    
        //i用来记录已经打印了几个元素，退出时，cornerPos中有连续两个相同的坐标，这指示着最后一个元素的坐标
        //根据执行append的条件，相邻两个角点，append的范围为[corner1,corner2),因此若corner1=corner2时，（最后一个元素）是不会打印的
        //所以当打印了r*c-1个元素后，就可以退出循环了，在循环外面再将最后一个元素添加进去即可
        if corner%4 == 0 {  //右移
            for ; cc < cornerPos[1][1] ; cc++ {
                i++
                ans = append(ans,matrix[rr][cc])
            }
            cornerPos[0][0],cornerPos[0][1] = cornerPos[0][0]+1,cornerPos[0][0]
            corner = 1
        }else if corner%4 == 1 {    //下移
            for ; rr < cornerPos[2][0] ; rr++ {
                i++
                ans = append(ans,matrix[rr][cc])
            }
            cornerPos[1][0],cornerPos[1][1] = cornerPos[1][0]+1,cornerPos[1][1]-1
            corner = 2
        }else if corner%4 == 2 {    //左移
            for ; cc > cornerPos[3][1] ; cc-- {
                i++
                ans = append(ans,matrix[rr][cc])
            }
            cornerPos[2][0],cornerPos[2][1] = cornerPos[2][0]-1,cornerPos[2][1]-1
            corner = 3
        }else{  //上移
            for ; rr > cornerPos[0][0] ; rr-- {
                i++
                ans = append(ans,matrix[rr][cc])
            }
            cornerPos[3][0],cornerPos[3][1] = cornerPos[3][0]-1,cornerPos[3][1]+1
            corner = 0
        }
    }
    ans = append(ans,matrix[rr][cc])
    
    return ans
}//啊我好笨