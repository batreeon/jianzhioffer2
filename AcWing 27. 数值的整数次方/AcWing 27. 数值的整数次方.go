package main

import "fmt"

// func Power(base float64, exponent int) float64 {
// 	//fmt.Println(exponent)
// 	if base == 0 {
// 		return 0
// 	}
// 	if exponent == 0 {
// 		return 1
// 	}
// 	if exponent == 1 {
// 		return base
// 	}
// 	if exponent < 0 {
// 		base, exponent = 1/base, -1*exponent
// 	}
// 	//return base * Power(base,exponent-1)  //这个方法太慢了，大指数内存爆炸
// 	var power float64
// 	if exponent%2 == 0 {
// 		power = Power(base, exponent/2)
// 		power = power * power
// 	} else {
// 		power = Power(base, exponent/2)
// 		power = base * power * power
// 	}
// 	return power
// }

// func Power(base float64, exponent int) float64 {
//  //fmt.Println(exponent)
//     if base == 0 {
//         return 0
//     }
//     if exponent == 0 {
//         return 1
//     }
//     if exponent == 1 {
//         return base
//     }
//     if exponent < 0 {
//         base,exponent = 1/base,-1*exponent
//     }
//     //return base * Power(base,exponent-1)  //这个方法太慢了，大指数内存爆炸

//     power := Power(base,exponent>>1)
//     power *= power
//  if exponent%2 == 1 {
//         power *= base
//     }
//     return power
// }
//golang这个，两种写法逻辑一样的，得到的数值有一些差异，我不太了解
//啊啊啊啊，上面死在了1.00000001，999999997上

func Power(base float64, exponent int) float64 {
	//fmt.Println(exponent)
	// if base == 0 {
	//     return 0
	// }
	// if exponent == 0 {
	//     return 1
	// }
	// if exponent == 1 {
	//     return base
	// }
	//return base * Power(base,exponent-1)  //这个方法太慢了，大指数内存爆炸
	is_minus := exponent < 0
	var power float64
	power = 1.0
	if exponent < 0 {
		exponent = -exponent
	}
	for ; exponent > 0; exponent >>= 1 {
		if exponent&1 == 1 {
			power *= base
		}
		base *= base
	}
	if is_minus {
		power = 1 / power
	}
	return power
}

//快速幂，递归法会使精度降低，因此在acwing上无法通过
func main() {
	//fmt.Println(Power(-2,-4))
	fmt.Println(Power(1.00000001, 999999997))
}
