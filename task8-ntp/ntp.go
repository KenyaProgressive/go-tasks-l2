package main

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

// PrintCurrentTime выводит точное время по протоколу SNTP
func PrintCurrentTime(ntpServer string) int {
	currentTime, err := ntp.Time(ntpServer)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occurred: %s\n", err)
		return 1
	}
	fmt.Println("Current time:", currentTime.Format("02.01.2006 15:04:05"))
	return 0
}
