func searchArray(array [][]int, target int) bool {
    for i := 0 ; i < len(array) ;{//为什么不把i,j放在一个for里，我是为了先确保i合法，再用j
        for j := len(array[i])-1; ; {
            //fmt.Println(array[i][j])
            if array[i][j] == target {
                return true
            }else if array[i][j] < target {
                i++
                //不需要break,因为若array[i][j+1]>target,则array[i+1][j+1]一定也大于target
            }else{
                j--
            }
            if i >= len(array) || j < 0 {
                return false
            }
            //不判断j可能导致外循环无限循环，无法退出。某一行全部大于target,i++执行不到
            //不判断i,可能导致越界，因为i++在内部循环执行，而原有的判断条件在外循环，因此需要在内循环增加判断
            //代码结构不太好，容易挖坑
        }
    }
    return false
}