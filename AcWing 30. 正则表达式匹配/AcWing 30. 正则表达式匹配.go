package main

import "fmt"

func isMatch(s string, p string) bool {
    var isMatchRec func(s string,p string) bool
    isMatchRec = func(s string,p string) bool {
        if len(s) == 0 && len(p) == 0 {
            return true
        }
        if len(s) != 0 && len(p) == 0 {
            return false
        }
		
		//模式向后移两位，匹配0个
        //字符串向后移一个，匹配1个，模式向后移两个或者不移
        if len(p) > 1 && string(p[1]) == "*" {
            if len(s) != 0 && ( s[0] == p[0] || string(p[0]) == "." ) {
                return isMatchRec(s[1:],p[2:]) || isMatchRec(s[1:],p) || isMatchRec(s,p[2:])//ab,a*b;/aab,a*b;/ba,a*ba;b,.*b
            }
            return isMatchRec(s,p[2:])
        }
		
		//这个判断放在后面是因为.*要在前面判断
        if len(s) != 0 && ( s[0] == p[0] || string(p[0]) == ".") {
                return isMatchRec(s[1:],p[1:])
        }
        
        return false
        
    }
    return isMatchRec(s,p)
}
//对于.*这种情况不太懂啊，好难的题，.*相当于a*,b*吧，那么aba,与.*a匹配吗?答：是匹配的

func main() {
	s, p := "aba", ".*a"
	fmt.Println(isMatch(s, p))
}
