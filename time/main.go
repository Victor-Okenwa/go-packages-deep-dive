package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Go time Package Deep Dive ===")

	// 1. Current Time
	now := time.Now()
	fmt.Println("1. Current Time:")
	fmt.Printf("Now: %v\n", now)
	fmt.Printf("UTC: %v\n", now.UTC())
	fmt.Printf("Unix Timestamp: %d\n", now.Unix()) // seconds since 1970
	fmt.Printf("Unix Milli: %d\n", now.UnixMilli())

	// 2. Creating Specific Times
	fmt.Println("\n2. Creating Specific Times:")

	birthday := time.Date(1995, time.June, 15, 14, 35, 0, 0, time.UTC)
	fmt.Printf("Birthday: %s\n", birthday)

	// 3. Formatting Time (Very Important)
	fmt.Println("\n3. Formatting Time:")

	fmt.Printf("Default: %v\n", now)
	fmt.Printf("RFC3339 %s\n", now.Format(time.RFC3339))              // Best for APIs
	fmt.Printf("Custom: %s\n", now.Format("2006-01-02 T15:04:25:06")) // Reference time
	fmt.Printf("Date only: %s\n", now.Format("2006-01-02"))
	fmt.Printf("Time only: %s\n", now.Format("15:04:05"))

	// 4. Parsing Time
	fmt.Println("\n4. Parsing Time:")

	dateStr := "2026-04-27 15:30:00"
	parsedTime, err := time.Parse("2006-01-02 15:04:05", dateStr)
	if err == nil {
		fmt.Printf("Parsed time: %v\n", parsedTime)
	}

	// 5. Durations (Very Important)
	fmt.Println("\n5. Durations:")

	duration := 5*time.Minute + 30*time.Second
	fmt.Printf("Duration: %v\n", duration)
	fmt.Printf("In seconds: %.0f\n", duration.Seconds())
	fmt.Printf("In minutes: %.2f\n", duration.Minutes())
	fmt.Printf("In hours: %.2f\n", duration.Hours())
	fmt.Printf("In milliseconds: %d \n", duration.Milliseconds())

	// Using durations
	fmt.Println("\n6. Using Durations:")
	fmt.Printf("Will finish in: %v\n", time.Now().Add(10*time.Minute))

	// 7. Time Differences
	fmt.Println("\n7. Time Difference:")

	start := time.Now()
	time.Sleep(2 * time.Second) // Simulate work
	elapsed := time.Since(start)

	fmt.Printf("Operation took: %v\n", elapsed)
	fmt.Printf("In microseconds: %d\n", elapsed.Microseconds())

	// 8. Tick & Ticker (Useful for periodic tasks)
	fmt.Println("Creating a ticker that fires every 1 second...")

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop() // Always stop ticker to prevent memory leak!

	count := 0

	// We use a for-range loop to receive ticks
	for tick := range ticker.C {
		count++
		fmt.Printf("Tick #%d at %s\n", count, tick.Format("15:04:05"))

		if count >= 5 {
			fmt.Println("Stopping ticker after 5 ticks...")
			break
		}
	}

	fmt.Println("\n=== Real-world Use Cases of Ticker ===")

	// Example 1: Sending metrics every 10 seconds
	fmt.Println("1. Sending metrics every 10 seconds (in real app)")

	// Example 2: Refreshing cache periodically
	fmt.Println("2. Refreshing cache every 5 minutes")

	// Example 3: Health check / Heartbeat
	fmt.Println("3. Sending heartbeat to monitoring system every 30 seconds")

	// 9. Best Practices Summary
	fmt.Println("\n=== Best Practices ===")
	fmt.Println("• Store times in UTC in databases")
	fmt.Println("• Use time.RFC3339 for JSON APIs")
	fmt.Println("• Use time.Duration for timeouts and intervals")
	fmt.Println("• Prefer time.Now().UTC() for server timestamps")

	infiniteTicker()
}

func infiniteTicker() {
	fmt.Println("Ticker running infinitely... (Press Ctrl+C to stop)")

	// Create a ticker that fires every 2 seconds
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop() // Important: This will run when program exits

	count := 0

	// This loop runs forever until you stop the program
	for range ticker.C {
		count++
		fmt.Printf("Tick #%d - Time: %s\n", count, time.Now().Format("15:04:05"))

		// You can put any periodic task here
		// e.g., send metrics, clean cache, refresh data, etc.
	}
}
