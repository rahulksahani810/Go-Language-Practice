package main

import (
	"fmt"
)

func change(val *int) {
	*val = 55
}
func main() {
	b := 255
	a := &b
	c := 500
	fmt.Println("address of b is", a)
	fmt.Println("value of b is", *a)
	*a++
	fmt.Println("new value of b is", b)
	d := &c
	change(d)
	fmt.Println("value of a after function call is", c)
}
