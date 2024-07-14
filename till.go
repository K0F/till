package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./till <time HH:MM>")
		os.Exit(1)
	}

	timeString := os.Args[1]

	// in some cases this gives wrong timezone
	//loc, err := time.LoadLocation("Local")
	//to set it manually you can define timezone here

	loc, err := time.LoadLocation("Europe/Prague")
	if err != nil {
		fmt.Println("Error loading location:", err)
		os.Exit(1)
	}

	// Get the current time
	currentTime := time.Now().In(loc)

	// Parse the input time string
	targetTime, err := time.Parse("15:04", timeString)
	if err != nil {
		fmt.Printf("Error parsing time: %v\n", err)
		os.Exit(1)
	}

	localTime := currentTime.In(loc)

	// Calculate the duration until the target time
	targetDateTime := time.Date(localTime.Year(), localTime.Month(), localTime.Day(), targetTime.Hour(), targetTime.Minute(), 0, 0, localTime.Location())
	duration := targetDateTime.Sub(currentTime)

	// If the target time is earlier than the current time, add 24 hours to sleep until the next day
	if duration < 0 {
		duration += 24 * time.Hour
	}

	// Create a timer that will fire after 'duration' time has passed
	timer := time.NewTimer(duration)

	fmt.Printf("Sleeping until %s, that is in %v.\n", targetDateTime.Format("3:04 PM"), duration)

	// Block until the timer expires
	<-timer.C

	fmt.Printf("Target time reached %s!\n", targetDateTime.Format("3:04 PM"))

	// Perform any actions after waking up
	fmt.Println("Unblocking...")

	os.Exit(0)
}
