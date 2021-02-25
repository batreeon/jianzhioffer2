func maxSubArray(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
	maxV,nowV := nums[0],0
	//为什么nowV设为0，不一样设为nums[0],然后从1开始遍历呢
	//因为那样不能简单地把nowV设为nums[0],而是若nums[0]<0,则应将nowV设为0，要多一步判断，还不如直接这样写呢
    for right := 0 ; right < len(nums) ; right++ {
        nowV += nums[right]
        if nowV > maxV {
            maxV = nowV
        }
        if nowV <= 0 {
            nowV = 0
        }
    }
    return maxV
}