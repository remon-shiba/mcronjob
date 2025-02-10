package main

import (
	"cronjob/service"
)



func main() {
	service.UpdateEverySecond()
	service.UpdateEveryMinute()
}
