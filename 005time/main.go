package main

import (
	"fmt"
	"time"
)

func main() {

	presentTime := time.Now()

	fmt.Println(presentTime)

	fmt.Print(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.UTC)

	fmt.Println(createdDate.Format("2006-01-02 Monday"))

	// fmt.Printf()
}
