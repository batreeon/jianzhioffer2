func duplicateInArray(nums []int) int {
    for i := 0 ; i < len(nums) ; i++ {
        for j := i+1 ; j < len(nums) ; j++ {
            if nums[j] == nums[i] {
                return nums[j]
            }
        }
    }
    return -1
}