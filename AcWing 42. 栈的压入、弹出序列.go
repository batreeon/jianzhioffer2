func isPopOrder(pushV []int , popV []int) bool{
    if len(pushV) == 0 && len(popV) == 0 {
        return true
    }
    if len(pushV) != len(popV) {
        return false
    }
    s := []int{}
    popIndex := 0
    for pushIndex := 0 ; pushIndex < len(pushV) ; pushIndex++ {
        s = append(s,pushV[pushIndex])
        for len(s) > 0 && s[len(s)-1] == popV[popIndex] {
            s = s[0:len(s)-1]
            popIndex++
        }
    }
    return !(len(s) > 0)
}
//若出队出现312这样的就不合法。213合法。
//模拟入队出队，看的解答