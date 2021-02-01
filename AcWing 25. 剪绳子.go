func maxProductAfterCutting(length int) int {
	//动规或贪婪
	//动规：f(n) = max(f(i)*f(n-i))  0<i<n  f(2)=1*1=1(必须要砍一刀)，f(3)=2
		// if length < 2 {
		//     return 0//
		// }
		if length == 2 {
			return 1//虽然不砍更大，但必须砍一刀
		}
		if length == 3 {
			return 2//必须砍一刀
		}
		profits := make([]int,length+1)
		profits[0],profits[1],profits[2],profits[3] = 0,1,2,3
		//为什么这个和f(2),f(3)值不同呢？因为前面也说了，2和3的时候，不砍，原值反倒会更大，因此在后面用到f(2),f(3)时，就没必要砍了
	
		for i := 4 ; i <= length ; i++ {
			max_i := 0
			for j := 1 ; j <= i/2 ; j++ {
				profit := profits[j] * profits[i-j]
				if profit > max_i {
					max_i = profit
					profits[i] = max_i
				}
			}
		}
		return profits[length]
	}