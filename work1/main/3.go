package main

import "fmt"

func main() {
	var date [2]int
	var leapdate []int
	var sum int = 0
	for i := range date {
		fmt.Scan(&date[i])
	}
	for i := date[0]; i <= date[1]; i++ {
		if isLeap(i) {
			sum++
			leapdate = append(leapdate, i)
		}
	}
	fmt.Println(sum)
	for i := range leapdate {
		fmt.Print(leapdate[i], " ")
	}
}

func isLeap(date int) bool {
	if (date%400 == 0) || (date%100 != 0 && date%4 == 0) {
		return true
	} else {
		return false
	}
}
