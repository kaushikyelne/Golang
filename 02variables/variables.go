package main

import "fmt"

var LocalVar = 21000 // Public
// value := 77690 // This is not allowed
func main() {
	var username string = "kaushik"
	fmt.Println(username)
	fmt.Printf("username is of type: %T \n", username)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("username is of type: %T \n", smallVal)

	var largeVal uint32 = 25566
	fmt.Println(largeVal)
	fmt.Printf("username is of type: %T \n", largeVal)

	var floatVal float64 = 255.367788
	fmt.Println(floatVal)
	fmt.Printf("username is of type: %T \n", floatVal)

	value := 77690
	fmt.Printf("value is of type: %T, %d\n", value, value)

}
