func getLeastNumbers_Solution(input []int , k int) []int{
    //当快速排序partion函数返回值为k-1,则最小的k个值就排在了前面
    //在用自己实现的快排，对结果进行排序
    if len(input) == 0 || k == 0 {
        return []int{}
    }
    var partition func(nums []int,p int,r int) int
    var quickSort func(nums []int,start int,end int)
    var QuickSort func(nums []int)
	partition = func(nums []int,p int,r int) int {
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
    quickSort = func(nums []int,start int,end int) {
    	if start < end {
    		q := partition(nums,start,end)
    		quickSort(nums,start,q-1)
    		quickSort(nums,q+1,end)
    	}
    }
    QuickSort = func(nums []int) {
    	quickSort(nums,0,len(nums)-1)
    }
    
    start,end := 0,len(input)-1
    index := partition(input,start,end)
    for index != k-1 {
        if index > k-1 {
            end = index-1
        }else{
            start = index+1
        }
        index = partition(input,start,end)
    }
    output := make([]int,k)
    copy(output,input[:k])
    QuickSort(output)
    return output
}