package main

import (
	"fmt"	
	"strings"
	"strconv"
)


func main() {
	s := string("[2, null, null]")
	ss := s[1:len(s)-1]
	sss := strings.Split(ss,", ")//会用指定的双字符来分隔
	fmt.Println(sss[0],sss[1],sss[2])
	fmt.Println(sss[0]," ",sss[1],sss[2])
	fmt.Println(strings.TrimSpace(sss[0])," ",strings.TrimSpace(sss[1]),strings.TrimSpace(sss[2]))
	//在printlin中，用逗号将输出各项隔开，会自动加上空格
	var sss0 int
	sss0,_ = strconv.Atoi(sss[0])
	fmt.Println(sss0)
}