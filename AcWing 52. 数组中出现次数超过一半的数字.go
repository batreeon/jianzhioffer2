func moreThanHalfNum_Solution(nums []int) int {
    //要找的数字比其他数字总和还要多
    times := 1
    ans := nums[0]
    for i := 1 ; i < len(nums) ; i++ {
        if times == 0 {
            ans = nums[i]
            times = 1
        }else if nums[i] == ans {
            times++
        }else {
            times--
        }
    }
    return ans
}