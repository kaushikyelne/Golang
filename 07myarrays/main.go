package main

import "fmt"

func main() {
	var countries [4]string
	countries[0] = "India"
	countries[1] = "USA"
	countries[3] = "Russia"
	fmt.Println(countries, len(countries))
}
