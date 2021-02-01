func movingCount(threshold, rows, cols int) int {
	//这道题，不能只看数位之和，还要考虑到机器人每次只能移动一格，因此数位之和满足也不一定能到达
		if threshold < 0 || rows <= 0 || cols <= 0 {
			return 0
		}
	
		visited := make([]bool,rows*cols)
	
		var mC func(x,y int) int
		var check func(x,y int) bool
		var getDigitSum func(n int) int
	
		mC = func(x,y int) int {
			if check(x,y) {
				visited[x*cols+y] = true
				return 1+mC(x-1,y)+mC(x,y-1)+mC(x+1,y)+mC(x,y+1)
			}
			return 0
		}
		check = func(x,y int) bool {
			if x < 0 || y < 0 || x >= rows || y >= cols || getDigitSum(x)+getDigitSum(y) > threshold || visited[x*cols+y] {
				return false
			}
			return true
		}
		getDigitSum = func(n int) int {
			sum := 0
			for ; n > 0 ; {
				sum += n%10
				n = n/10
			}
			return sum
		}
	
		count := mC(0,0)
		return count
	}