package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter any value: ")
	// comma ok || err ok
	input, _ := reader.ReadString('\n')
	fmt.Println("The value is:  ", input)

}
