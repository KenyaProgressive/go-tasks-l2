package main

import "os"

func main() {
	ntpServer := "0.beevik-ntp.pool.ntp.org"
	if exitCode := PrintCurrentTime(ntpServer); exitCode != 0 { // Если ошибка -- выход из программы с кодом ошибки
		os.Exit(exitCode)
	}

}
