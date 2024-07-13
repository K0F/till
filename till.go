package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <time>")
		return
	}

	timeString := os.Args[1]

	// Get the current time
	currentTime := time.Now()

	// Parse the input time string
	targetTime, err := time.Parse("15:04", timeString)
	if err != nil {
		fmt.Printf("Error parsing time: %v\n", err)
		return
	}

	// Calculate the duration until the target time
	targetDateTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), targetTime.Hour(), targetTime.Minute(), 0, 0, currentTime.Location())
	duration := targetDateTime.Sub(currentTime)

	// If the target time is earlier than the current time, add 24 hours to sleep until the next day
	if duration < 0 {
		duration += 24 * time.Hour
	}

	// Create a timer that will fire after 'duration' time has passed
	timer := time.NewTimer(duration)

	fmt.Printf("Sleeping until %s, that is in %v ahead.\n", targetDateTime.Format("3:04 PM"), duration)

	// Block until the timer expires
	<-timer.C

	fmt.Printf("Target time reached %s!\n", targetDateTime.Format("3:04 PM"))

	// Perform any actions after waking up
	fmt.Println("Unblocking...")
}
