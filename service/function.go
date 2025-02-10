package service

import (
	"fmt"
	"time"
)

func CallMe() {
	fmt.Println("Calling the function")
	fmt.Println("Job executed at", time.Now().Format(time.Kitchen))
}
