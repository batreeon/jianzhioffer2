package main

import (
	"fmt"
	"strings"
	"strconv"
)
func main() {
	n := 13
	numString := []string{}
    for i := 0 ; i <= n ; i++ {
        numString = append(numString,strconv.Itoa(i))
	}
	fmt.Printf("%T ",strings.Join(numString,"")[n] - '0')
    fmt.Println(strings.Join(numString,"")[n] - '0')
}