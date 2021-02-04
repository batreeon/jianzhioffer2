package main

import (
	"fmt"
	"strings"
)

func isNumber(s string) bool {
	// fmt.Println(s)
	ans := strings.Trim(s, " ")
	// fmt.Println(ans)
	// fmt.Println(ans[0])
	// fmt.Println(ans[0]-'0')
	// fmt.Println('0')
	// fmt.Println('9')
	// fmt.Println(ans[0]<'9')
	if len(ans) == 0 {//排除""
		return false
	}
	if ans[0] == '-' || ans[0] == '+' {
		ans = ans[1:]
	}
	if len(ans) == 0 {//排除"+","-"
		return false
	}
	dot,e := 0,0 //记录小数点和e的个数
	for i := 0 ; i < len(ans) ; i++ {
		if ans[i] >= '0' && ans[i] <= '9' {
			
		}else if ans[i] == '.' {
			dot++
			if e > 0 || dot > 1 {   //小数点在e后出现或者有多个小数点
				return false
			}
			if i == 0 && i+1 == len(ans){   //.
				return false
			}
		}else if ans[i] == 'e' || ans[i] == 'E' {
			e++
			if i+1 == len(ans) || i == 0 || (i == 1 && ans[0] == '.') {  //e后面没有数字，e前面没有数字,.e12
				return false
			}
			if ans[i+1] == '+' || ans[i+1] == '-' {
				if i+2 == len(ans) {
					return false
				}
				i++
			} 
		}else{//其他非法字符
			return false
		}
	}
	return true
}

func main() {
	// s := string("0")
	//s := string("123.45e+6")
	s := string("-")
	fmt.Println(isNumber(s))
}
