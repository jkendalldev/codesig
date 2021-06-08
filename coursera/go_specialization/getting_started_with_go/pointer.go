package main

import (
	"fmt"
)

func main() {
	var x int = 1
	fmt.Println(x)

	var y int
	fmt.Println(y)

	var ip *int // ip is pointer to
	ip = &x     // ip now points to address of x
	fmt.Println(ip)

	y = *ip // y is now 1
	fmt.Println(y)

	ptr := new(int)
	fmt.Println(ptr)
	*ptr = 3
	fmt.Println(*ptr)
}
