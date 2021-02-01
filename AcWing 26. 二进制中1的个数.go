func NumberOf1(n int) int {
    count := 0
    for ; n != 0 ; {
        count++
        n = (n-1)&n
    }
    return count
}