package main

import (
	"cronjob/service"
	"fmt"
	"os"
)

func main() {
	service.UpdateEverySecond()
	service.UpdateEveryMinute()

	fmt.Println("ENV:", os.Getenv("PROJECT"))
}
