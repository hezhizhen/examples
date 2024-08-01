package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	end := start.Add(time.Minute)

	fmt.Println(end.Sub(start))                      // 1m0s
	fmt.Println(end.Unix() - start.Unix())           // 60    unit: second
	fmt.Println(end.UnixMilli() - start.UnixMilli()) // 60000 unit: millisecond
}
