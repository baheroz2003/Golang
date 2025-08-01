package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("Full timestamp:", now)
	fmt.Println("Milliseconds:", now.UnixMilli())
	fmt.Println("Nanoseconds:", now.UnixNano())
	fmt.Println(now)
	fmt.Println(now.Format("01-02-2006 15:04:05 Monday"))
	createdDate := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
