func duplicateInArray(nums []int) int {
    count := make([]int,len(nums))
    ans := -1
    for _,num := range nums {
        if num < 0 || num > len(nums)-1 {
            return -1
        }
        count[num]++
        if count[num] > 1 {
            ans = num
        }//不直接返回是因为有可能重复的元素在前，超界的元素在后，这种情况应返回-1，而不是重复元素
    }
    return ans
}