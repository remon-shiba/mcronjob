package service

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

// Minute	Minute field	0-59
// Hour	Hour field	0-23
// Day	Day of month	1-31
// Month	Month field	1-12
// Weekday	Day of week	0-6 (0 is Sunday)

// Example patterns:
// "*/5 * * * *": Every 5 minutes
// "0 9 * * *": Every day at 9:00 AM
// "0 0 * * 0": Every Sunday at midnight
func CallMe(message string) {
	fmt.Println("Calling the function:", message)
	fmt.Println("Job executed at", time.Now().Format(time.Kitchen))
}

func UpdateEverySecond() {
	// Create a new cron with seconds precision enabled
	c := cron.New(cron.WithSeconds())

	// Schedule the task to run every second
	c.AddFunc("*/10 * * * * *", func() {
		CallMe("Every second")
	})

	// Start the cron scheduler
	c.Start()

	// Keep the program running
	select {}
}

func UpdateEveryMinute() {
	c := cron.New()

	// Schedule a job to run every minute
	c.AddFunc("*/1 * * * *", func() {
		CallMe("Every minute")
	})

	// Start the cron scheduler
	c.Start()

	// Keep the program running
	select {}
}
