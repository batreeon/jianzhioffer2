package main

import "sort"
// 有重复元素
func permutation(nums []int) [][]int {
	sort.Ints(nums)
	l := len(nums)
	seen := map[int]bool{}
	ans := [][]int{}

	var backtracking func(res *[]int)
	backtracking = func(res *[]int) {
		if len(*res) == l {
			resCopy := make([]int,l)
			copy(resCopy,*res)
			ans = append(ans,resCopy)
			return
		}
		for i := 0 ; i < l ; i++ {
			if seen[i] {continue}
			if i > 0 && nums[i-1] == nums[i] && !seen[i-1] {
				continue
			}
			*res = append(*res,nums[i])
			seen[i] = true
			backtracking(res)
			*res = (*res)[:len(*res)-1]
			delete(seen,i)
		}
	}

	backtracking(&[]int{})
	return ans
}