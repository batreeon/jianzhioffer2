func reOrderArray(array []int){
    oddTail := 0
    for i := 0 ; i < len(array) ; i++ {
        if array[i]%2 != 0 {//奇数
            array[oddTail],array[i] = array[i],array[oddTail]
            oddTail++
        }
    }
    //把奇数移到前面就行
}