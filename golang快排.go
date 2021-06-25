func QuickSort(nums []int) []int {
	quickSort(nums,0,len(nums)-1)
	return nums
}
func quickSort(nums []int,start int,end int) {
	if start < end {
		q := partition(nums,start,end)
		quickSort(nums,start,q-1)
		quickSort(nums,q+1,end)
	}
}
func partition(nums []int,p int,r int) int {
	x := nums[r]
	i := p-1
	for j := p ; j < r ; j++ {
		if nums[j] < x {
			i++
			nums[i],nums[j] = nums[j],nums[i]
		}
	}
	nums[i+1],nums[r] = nums[r],nums[i+1]
	return i+1
}