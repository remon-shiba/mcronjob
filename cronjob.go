package main

import (
	"cronjob/service"

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

func main() {
	c := cron.New()

	// Schedule a job to run every minute
	c.AddFunc("*/1 * * * *", service.CallMe)

	// Start the cron scheduler
	c.Start()

	// Keep the program running
	select {}
}
