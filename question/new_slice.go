/*
t := baseStr[low:hight:max]
len = hight - low
cap = max - low
*/
package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	t := a[3:4:4]
	fmt.Println(t[0])           // 4
	fmt.Println(len(t), cap(t)) // 1 1
}
