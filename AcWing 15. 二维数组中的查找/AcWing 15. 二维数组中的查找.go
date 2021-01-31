package main

import (
	"fmt"
)
func searchArray(array [][]int, target int) bool {
    for i := 0 ; i < len(array) ;{
        for j := len(array[i])-1; ; {
			fmt.Println(array[i][j])
            if array[i][j] == target {
                return true
            }else if array[i][j] < target {
				i++
                //不需要break,因为若array[i][j+1]>target,则array[i+1][j+1]一定也大于target
            }else{
				j--
			}
			if i >= len(array) || j < 0 {
				return false
			}
        }
    }
    return false
}
func main() {
	array := [][]int{}
	array0 := []int{1,2,8,9}
    array1 := []int{2,4,9,12}
	array2 := []int{4,7,10,13}
	array3 := []int{6,8,11,15}
	array = append(array,array0)
	array = append(array,array1)
	array = append(array,array2)
	array = append(array,array3)
	fmt.Println(array)
	if searchArray(array,16) {
		fmt.Println(1)
	}
	
}