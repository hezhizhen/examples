package main

import (
	"fmt"
	"time"
)

func main() {
	for year := 2021; year <= 2024; year++ {
		for i := 1; i <= 12; i++ {
			now := time.Date(year, time.Month(i), 31, 0, 0, 0, 0, time.Local)
			for now.Month() != time.Month(i) {
				now = now.AddDate(0, 0, -1)
			}
			fmt.Println(now.Format(time.DateOnly))
			previousMonth(now)
			nextMonth(now)
			fmt.Println("------")
		}
	}
}

func previousMonth(now time.Time) {
	same := true
	result := now.AddDate(0, -1, 0)
	fmt.Printf("Previous month:\t%v -> ", result.Format(time.DateOnly))
	for result.Month() == now.Month() {
		result = result.AddDate(0, 0, -1)
		same = false
	}
	fmt.Printf("%v %v\n", result.Format(time.DateOnly), generateMark(same))
}

func nextMonth(now time.Time) {
	same := true
	result := now.AddDate(0, 1, 0)
	fmt.Printf("Next month:\t%v -> ", result.Format(time.DateOnly))
	for result.Month()-now.Month() > 1 {
		result = result.AddDate(0, 0, -1)
		same = false
	}
	fmt.Printf("%v %v\n", result.Format(time.DateOnly), generateMark(same))
}

func generateMark(same bool) string {
	if same {
		return "✅ "
	}
	return "❌ "
}
