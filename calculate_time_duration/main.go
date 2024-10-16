package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	end := start.Add(time.Minute)

	fmt.Println(end.Sub(start)) // 1m0s

	fmt.Println(end.Unix() - start.Unix())           // 60          seconds
	fmt.Println(end.UnixMilli() - start.UnixMilli()) // 60000       milliseconds
	fmt.Println(end.UnixMicro() - start.UnixMicro()) // 60000000    microseconds
	fmt.Println(end.UnixNano() - start.UnixNano())   // 60000000000 nanoseconds
}
