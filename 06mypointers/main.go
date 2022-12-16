package main

import "fmt"

func main() {
	fmt.Println("wELCOME TO PONTERS")
	num := 23
	var ptr = &num
	fmt.Println(num, ptr, &ptr, &num, *ptr+10)
}
