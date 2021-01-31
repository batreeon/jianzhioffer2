func Fibonacci(n int) int {
    if n == 0 {
        return 0
    }
    if n == 1 {
        return 1
    }
    f1,f2 := 0,1
    for ; n > 1 ; n-- {
        f1,f2 = f2,f1+f2
    }
    return f2
}