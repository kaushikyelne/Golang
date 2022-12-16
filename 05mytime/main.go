package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study")
	presentTime := time.Now()
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(1998, time.October, 11, 0, 0, 0, 0, time.Local)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
